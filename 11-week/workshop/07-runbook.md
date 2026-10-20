# CP8 - Runbook: ClinicalProcedureCompletionFastBurn

Alert: `ClinicalProcedureCompletionFastBurn` (`06-alerts.yml`), severity `page`, team `di-lucca-clinical`.
Meaning: more than 7.2 % of `POST /api/v1/procedures/{id}/completions` failed with 5xx or timeout in the last hour
and in the last 5 minutes, with at least 2 failures. 2 % of the 30-day budget is already spent.

## Symptoms

- A Dentist clicks "complete procedure" in the clinical portal and gets an error or a spinner, with the patient in the
  chair.
- The portal shows the error envelope `{error, message, traceId}` with `INTERNAL_ERROR` or a gateway timeout.
- The Assistant cannot finalize visits later, because no `ProcedureCompleted` reached Appointments.

## 1. Confirm (5 minutes)

Error ratio and volume right now:

```promql
clinical:procedure_completion_error_ratio:rate5m
sum(increase(clinical_procedure_completions_total{outcome="error"}[1h]))
histogram_quantile(0.95, sum by (le) (rate(clinical_procedure_completion_duration_seconds_bucket[5m])))
```

If p95 latency is near or above 2.5 s, the failures are timeouts (MongoDB). If latency is normal, the failures are
fast 5xx (code or config).

Take the `traceId` from the portal error or from one failed log line, then read the whole story:

```bash
docker compose -f deploy/compose.yml logs --no-log-prefix --since 1h clinical-api | grep '"trace_id":"<trace_id>"'
docker compose -f deploy/compose.yml logs --no-log-prefix --since 1h clinical-api | grep '"level":"ERROR"' | grep '"event":"procedure.completion' | tail -20
```

Look at the `event` field (CP2 contract): `mongo.transaction.timeout`, `mongo.transaction.aborted` or a panic logged as
`INTERNAL_ERROR`. Also check if the alert started right after a deploy: compare with the last image tag of
`dlc-clinical-api` on `qa` or `main`.

## 2. Mitigate (ordered from lowest to highest risk)

| # | Action | When | How | Risk |
|---|---|---|---|---|
| 1 | Restart the `clinical-service` container | errors after a MongoDB failover or a stuck connection pool | `docker compose -f deploy/compose.yml restart clinical-api` | low: completions are idempotent by `procedureId` and the idempotency key, so a Dentist retry does not duplicate anything |
| 2 | Roll back to the previous image tag | the alert started within 1 hour of a deploy | set the previous tag (for example `v2.0.0`) in the compose file and `docker compose up -d clinical-api` | medium: any new endpoint of the release is gone; the `dlc-clinical-db` schema must still be compatible with the old tag |
| 3 | Fix MongoDB capacity | p95 latency is high and logs show `mongo.transaction.timeout` | check `dlc-infra-mongo` container health and disk; restart the replica member that is unhealthy | high: touches the shared clinical data store, coordinate with the infra owner first |

What we do NOT do: we have no feature flags and no circuit breaker in `dlc-clinical-api`, and we never disable the
outbox insert, because the procedure and `ProcedureCompleted` must stay in one transaction or Billing loses charges.

While mitigating, tell the clinic: the Dentist writes the procedure on paper and the Assistant enters it later. The
completion can be safely retried, because the same `procedureId` gives the same `CareClosure`.

## 3. Escalate

| After | Who | Role |
|---|---|---|
| 0 min | Luis Bonilla (LuisBonilla2260) | `clinical-service` owner, first responder |
| 15 min without mitigation | team lead of Di-Lucca | decides rollback and talks to the clinic |
| MongoDB or network cause | owner of `dlc-infra-mongo` / infra | data store and compose network `platform` |
| failures move downstream (`CLOSURE_FAILED`) | owners of `appointments-service` and `billing-service` | consumers of `ProcedureCompleted` |

Names for the other roles are assigned in the team `-docs` repository; this runbook names the role so it does not go
stale.

## 4. Verify recovery

- `clinical:procedure_completion_error_ratio:rate5m` is 0 for 15 minutes and the alert resolves on its own (the 5m
  window drops first).
- One real completion by a Dentist returns `202` with `CLOSURE_PENDING`, and its `trace_id` shows
  `outbox.publish.sent` and then `procedure.progress.recorded` in Appointments.
- No `CareClosure` from the incident stays in `CLOSURE_PENDING` for more than 5 minutes (SLO 3). Any that do are
  reconciled by hand with `GET /api/v1/care-closures/{id}`.
- Write down the minutes and failed completions spent from the budget for the postmortem (CP9).

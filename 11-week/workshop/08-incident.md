# CP9 - Incident drill and blameless postmortem

## Scenario

The hop marked in CP1 as the weak point is step 3: `clinical-service` -> Clinical MongoDB (transaction on `procedures`,
`procedure_extras`, `care_closures`, `outbox_events`). In the drill, MongoDB answers in **3 seconds for 25 % of the
requests for 40 minutes**, Tuesday 2026-10-20 from 10:00 to 10:40, during clinic hours.

Facts from `dlc-clinical-api` that decide what happens: `MONGO_TIMEOUT` is 5 s and `WRITE_TIMEOUT` is 15 s
(`internal/config/config.go`). A 3 s answer is slow but **does not fail**.

## One slow request

Span tree of trace `c3e81f0a5b2d4c97a6e1f8b04d2a9c35`:

- `POST /api/v1/procedures/{id}/completions` - API Gateway - **3,071 ms**
  - `POST /api/v1/procedures/{id}/completions` - clinical-service - 3,058 ms
    - `jwt.verify` - 2 ms
    - `mongo.transaction complete_procedure` - **3,041 ms**
      - `mongo.update procedures` - 3,012 ms <- waits on the slow primary
      - `mongo.insert procedure_extras` - 9 ms
      - `mongo.insert care_closures` - 8 ms
      - `mongo.insert outbox_events` - 10 ms
- `outbox.relay ProcedureCompleted` - clinical-service - 14 ms
  - `ProcedureCompleted publish` - 11 ms
    - `ProcedureCompleted process` - appointments-service - 22 ms
    - `ProcedureCompleted process` - billing-service - 31 ms

Logs of that trace (CP2 format):

```json
{"timestamp":"2026-10-20T10:17:42.118Z","level":"WARN","service":"clinical-service","version":"2.0.0","environment":"main","trace_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","span_id":"5e2a9d1c7b3f4086","correlation_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","event":"mongo.transaction.slow","message":"complete_procedure transaction above 1s","duration_ms":3041,"procedure_id":"8b4f1e27-6c90-4d3a-b5e8-2a7f0c9d1e46"}
{"timestamp":"2026-10-20T10:17:42.135Z","level":"INFO","service":"clinical-service","version":"2.0.0","environment":"main","trace_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","span_id":"a71d3f9e2c5b4810","correlation_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","event":"procedure.completion.accepted","message":"procedure completion committed with outbox event","duration_ms":3058,"procedure_id":"8b4f1e27-6c90-4d3a-b5e8-2a7f0c9d1e46","procedure_code":"OPER-01","closure_status":"CLOSURE_PENDING"}
{"timestamp":"2026-10-20T10:17:42.162Z","level":"INFO","service":"clinical-service","version":"2.0.0","environment":"main","trace_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","span_id":"0d6b4a8e1f2c9357","correlation_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","event":"outbox.publish.sent","message":"ProcedureCompleted published","duration_ms":11,"procedure_id":"8b4f1e27-6c90-4d3a-b5e8-2a7f0c9d1e46"}
{"timestamp":"2026-10-20T10:17:42.215Z","level":"INFO","service":"billing-service","version":"0.3.0","environment":"main","trace_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","span_id":"e94c2b7a0d5f1368","correlation_id":"c3e81f0a5b2d4c97a6e1f8b04d2a9c35","event":"charge.snapshot.created","message":"charge snapshot created for procedure","duration_ms":31,"procedure_id":"8b4f1e27-6c90-4d3a-b5e8-2a7f0c9d1e46","procedure_code":"OPER-01"}
```

## What the SLIs show

Traffic is 6 completions per hour (`05-slo.md`), so 40 minutes = **4 completions**, and 25 % of them = **1 slow**.

| SLI | Value in the 40 min | 30-day effect |
|---|---|---|
| Availability (no 5xx/timeout) | 4 / 4 = 100 % | 0 of 5.28 failures used |
| Latency (<= 300 ms) | 3 / 4 = 75 %, target 95 % | 1 of 52 slow completions = **1.9 %** of the budget; 40 x 25 % = 10 degraded minutes of 528 |
| Closure convergence (<= 5 min) | 4 / 4 = 100 % | 0 used: the outbox and the consumers were not affected |

## Which alert fires

**None.** Both rules in `06-alerts.yml` watch availability, and no request failed. The fast burn needs at least 2
errors in 1 hour. In this drill the first signal would come from a Dentist complaining around 10:17, or from the
`mongo.transaction.slow` WARN logs if someone is reading them. This gap is the main finding of the drill.

If MongoDB had been slower than 5 s, the requests would end in `outcome="error"`: the 2nd failure (about minute 20 at
6 per hour and 25 %) would meet the guard and the 5m window, and `ClinicalProcedureCompletionFastBurn` would fire then.

## Which runbook step stops the damage

Step 1 confirm: `histogram_quantile(0.95, ...)` near 3 s with 0 errors points to MongoDB. Step 2 mitigation 3 (check
`dlc-infra-mongo` health and the slow replica member) is the one that fixes it; restarting `clinical-api` (mitigation
1) would not help, because the slowness is in the data store.

## Blameless postmortem

**Summary:** for 40 minutes one in four procedure completions took about 3 s instead of under 300 ms because the
MongoDB primary answered slowly. No data was lost and no charge was duplicated. Nobody was alerted.

**Timeline (2026-10-20):**

| Time | Event |
|---|---|
| 10:00 | MongoDB starts answering in 3 s for 25 % of operations |
| 10:17 | first slow completion (trace `c3e81f0a...`), `mongo.transaction.slow` WARN |
| 10:25 | Dentist reports "the save button is slow" to the Assistant |
| 10:30 | clinical owner checks p95 latency and the trace, sees `mongo.update procedures` at 3 s |
| 10:35 | infra owner restarts the unhealthy replica member (runbook step 2, mitigation 3) |
| 10:40 | latency back under 300 ms; 0 closures stuck in `CLOSURE_PENDING` |

**Impact:** 4 completions, 1 slow; 1.9 % of the latency budget (1 of 52 slow requests, 10 of 528 degraded minutes);
0 % of the availability and convergence budgets.

**Root cause:** the MongoDB primary of `dlc-infra-mongo` served part of its operations slowly; every completion runs
one transaction on it, so its latency goes straight to the Dentist.

**Contributing factors:**
- There is no latency alert, only availability alerts.
- `MONGO_TIMEOUT` (5 s) is longer than the time a Dentist will wait, so slowness never turns into a visible error.
- There is no MongoDB latency metric in `04-metrics.md`, so the cause had to be found from one trace.

**Actions (user stories):**

1. **As** the clinical owner, **I want** a burn-rate alert on the latency SLO **so that** slow completions page us
   before a Dentist complains.
   - Given `clinical_procedure_completion_duration_seconds` with more than 5 % over 300 ms in 1h and 5m, when the rule
     is evaluated, then `ClinicalProcedureCompletionLatencyBurn` fires with `severity: page`.
   - `promtool test rules` passes with a series of 4 completions where 2 take 3 s.
2. **As** a Dentist, **I want** completion to fail fast and tell me to retry **so that** I am not left waiting with
   the patient in the chair.
   - Given MongoDB slower than 1.5 s, when I complete a procedure, then I get an error envelope with `traceId` in
     under 2 s.
   - Retrying with the same idempotency key creates exactly one `ProcedureCompleted` in `outbox_events`.
3. **As** the infra owner, **I want** MongoDB operation latency exported **so that** we can tell a slow data store from
   slow code.
   - `/metrics` of `clinical-service` exposes `clinical_mongo_operation_duration_seconds{collection,operation}` with
     fewer than 20 series.
   - A dashboard panel shows its p95 next to the completion p95.

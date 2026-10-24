<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       12-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 12

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Make the clinical closing journey run on a real Clinical MongoDB replica set and close HU-XCT-001 end to end (portal on the ADR-014 encounter closure), then carry the week 11 workshop into dlc-clinical-api - snake_case logs with W3C traceparent through the outbox and RabbitMQ, a /metrics endpoint, and a tested latency alert linked to the MongoDB recovery runbook.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-012-001 | Provision the Clinical MongoDB 8.0 replica set in `dlc-infra-mongo` with owner-scoped access, so multi-document transactions (procedure + `outbox_events`) run outside tests | todo | [dlc-infra-mongo #1](https://github.com/code-corhuila/dlc-infra-mongo/issues/1), [#2](https://github.com/code-corhuila/dlc-infra-mongo/issues/2) |
| HU-012-002 | Align the clinical portal with the ADR-014 encounter closure: `ProcedureCompletion` model, one `CareClosure` from `care-completions`, `AWAITING_ADMINISTRATIVE_FINALIZATION` as a normal state | todo | [dlc-clinical-portal #111](https://github.com/code-corhuila/dlc-clinical-portal/issues/111) |
| HU-012-003 | Close HU-XCT-001 with acceptance evidence across `dlc-clinical-api` and `dlc-clinical-db` (idempotent procedure facts, Dentist declaration, outcomes from RabbitMQ) | todo | [dlc-docs #67](https://github.com/code-corhuila/dlc-docs/issues/67), [api #6](https://github.com/code-corhuila/dlc-clinical-api/issues/6), [db #8](https://github.com/code-corhuila/dlc-clinical-db/issues/8) |
| HU-012-004 | Observability in `dlc-clinical-api`: CP2 snake_case log contract, W3C `traceparent` stored in `outbox_events` and copied to AMQP headers by the relay, and `/metrics` with the completion counter and latency histogram (carries over HU-011-005) | todo | [workshop CP2-CP4](../../11-week/workshop/README.md) |
| HU-012-005 | Latency burn-rate alert `ClinicalProcedureCompletionLatencyBurn` with `promtool test rules`, and the MongoDB backup/restore runbook registered in `dlc-docs` (carries over HU-011-006) | todo | [dlc-docs #117](https://github.com/code-corhuila/dlc-docs/issues/117), [dlc-infra-mongo #3](https://github.com/code-corhuila/dlc-infra-mongo/issues/3) |

## 2. My individual contribution
- Reviewed the state of `dlc-docs`, `dlc-clinical-api`, `dlc-clinical-db`, `dlc-clinical-portal` and `dlc-infra-mongo` to plan the week:
  - `dlc-clinical-api` `develop` already has the procedure-completion and care-completion endpoints, the outbox relay to RabbitMQ with reconnect and dead-lettering, owner calls to Patients and Appointments, and the HU-CLN-003 analytics consumer.
  - It still has no `traceparent` handling and no `/metrics`, so the week 11 SLOs and alerts cannot be measured yet.
  - `dlc-infra-mongo` only contains a README; its four issues (replica set, access, backup, security) are open.
  - The portal still follows the per-procedure closure replaced by ADR-014 (#111).

## 3. Blockers and risks
- **No provisioned Clinical MongoDB:** transactions need a replica set and `dlc-infra-mongo` is empty. HU-012-003 and HU-012-004 depend on HU-012-001.
- **Portal and API out of sync:** until #111 is fixed, the Dentist sees a closure per procedure while the API creates one per encounter.
- **No Prometheus or Alertmanager in the stack:** HU-012-005 can be validated with `promtool` in Docker but not deployed.
- **Low traffic:** about 6 completions per hour; one failure is a large share of the error budget (week 11 workshop CP6).

## 4. Plan for next week
- Deploy Prometheus and Alertmanager next to the compose stack and load `06-alerts.yml` plus the latency rule.
- Propagate `traceparent` in the Appointments and Billing consumers of `ProcedureCompleted`.
- Run the week 11 incident drill (MongoDB at 3 s) against the real stack.

## 5. Compliance self-check
- [ ] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [ ] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [ ] DDD / hexagonal boundaries respected (domain has no I/O)
- [ ] No secrets; config via environment variables

## 6. Evidence links
- [dlc-docs #67 - HU-XCT-001 safe clinical-to-financial closure](https://github.com/code-corhuila/dlc-docs/issues/67)
- [dlc-docs #117 - register dlc-infra-mongo and the MongoDB recovery runbook](https://github.com/code-corhuila/dlc-docs/issues/117)
- [dlc-clinical-portal #111 - align procedure completion and care closure with ADR-014](https://github.com/code-corhuila/dlc-clinical-portal/issues/111)
- [dlc-infra-mongo #1 - Clinical MongoDB 8.0 replica set](https://github.com/code-corhuila/dlc-infra-mongo/issues/1)
- [Week 11 workshop](../../11-week/workshop/README.md)

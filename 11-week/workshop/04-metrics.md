# CP4 - Golden signals and metric catalog

Journey: `POST /api/v1/procedures/{id}/completions` -> `ProcedureCompleted` -> Appointments and Billing.
Names follow Prometheus conventions: snake_case, service prefix, unit suffix, `_total` for counters.

## clinical-service (Go, Clinical MongoDB)

| Signal | Metric | Type | Labels | PromQL |
|---|---|---|---|---|
| Latency | `clinical_procedure_completion_duration_seconds` | histogram | none | `histogram_quantile(0.95, sum by (le) (rate(clinical_procedure_completion_duration_seconds_bucket[5m])))` |
| Traffic | `clinical_procedure_completions_total` | counter | `outcome` | `sum(rate(clinical_procedure_completions_total[5m]))` |
| Errors | `clinical_procedure_completions_total` | counter | `outcome="error"` | `sum(rate(clinical_procedure_completions_total{outcome="error"}[5m])) / sum(rate(clinical_procedure_completions_total[5m]))` |
| Saturation | `clinical_outbox_pending_events` | gauge | `event_type` | `max(clinical_outbox_pending_events{event_type="ProcedureCompleted"})` |
| Saturation | `clinical_care_closures_failed` | gauge | none | `clinical_care_closures_failed` |

`outcome` values: `accepted` (202), `rejected` (400/403/409, business rule or stale version), `error` (5xx or timeout).
The first two rows are already exposed by `03-instrumentation/clinical-service/metrics.go`.

## appointments-service (Spring Boot, PostgreSQL schema `appointments`)

| Signal | Metric | Type | Labels | PromQL |
|---|---|---|---|---|
| Latency | `appointments_event_processing_duration_seconds` | histogram | `event_type` | `histogram_quantile(0.95, sum by (le) (rate(appointments_event_processing_duration_seconds_bucket{event_type="ProcedureCompleted"}[5m])))` |
| Traffic | `appointments_events_consumed_total` | counter | `event_type`, `outcome` | `sum(rate(appointments_events_consumed_total{event_type="ProcedureCompleted"}[5m]))` |
| Errors | `appointments_events_consumed_total` | counter | `outcome="failed"` | `sum(rate(appointments_events_consumed_total{event_type="ProcedureCompleted",outcome="failed"}[5m]))` |
| Saturation | `rabbitmq_queue_messages_ready` | gauge | `queue` | `rabbitmq_queue_messages_ready{queue="appointments.procedure-completed"}` |

## billing-service (Spring Boot, PostgreSQL schema `billing`)

| Signal | Metric | Type | Labels | PromQL |
|---|---|---|---|---|
| Latency | `billing_event_processing_duration_seconds` | histogram | `event_type` | `histogram_quantile(0.95, sum by (le) (rate(billing_event_processing_duration_seconds_bucket{event_type="ProcedureCompleted"}[5m])))` |
| Traffic | `billing_events_consumed_total` | counter | `event_type`, `outcome` | `sum(rate(billing_events_consumed_total{event_type="ProcedureCompleted"}[5m]))` |
| Errors | `billing_charge_snapshots_total` | counter | `outcome` (`created`, `duplicate`, `rejected_no_price`) | `sum(rate(billing_charge_snapshots_total{outcome="rejected_no_price"}[5m]))` |
| Saturation | `rabbitmq_queue_messages_ready` | gauge | `queue` | `rabbitmq_queue_messages_ready{queue="billing.procedure-completed"}` |

Queue names are proposals; the real names come from the RabbitMQ topology of each service.

## Label cardinality

| Label | Values | Cardinality |
|---|---|---|
| `outcome` | accepted, rejected, error / processed, failed, duplicate | 3 to 4 |
| `event_type` | `ProcedureCompleted`, `ClinicalCareCompleted`, `BillableItemsCreated`, ... (Clinical Care has 5) | under 25 |
| `queue` | one per consumer queue | under 15 |

Forbidden labels and what replaces them:

| Unbounded label | Why | Replacement |
|---|---|---|
| `procedure_id`, `appointment_id`, `patient_id` | one series per procedure or patient | log fields `procedure_id`, `appointment_id` (CP2), found by `trace_id` |
| raw path `/api/v1/procedures/5d0e.../completions` | one series per id | route template or a fixed metric per operation |
| `dentist_id`, email, JWT `sub` | personal data and unbounded | no label; `actor_hash` only in logs |
| `procedure_code` | catalog can grow to hundreds | log field; only aggregated reports in Billing |

## Latency buckets

`clinical_procedure_completion_duration_seconds`: 0.05, 0.1, 0.2, 0.3, 0.5, 1, 2.5 seconds.
The Dentist completes the procedure with the patient still in the chair, so the answer should feel instant: 300 ms is
the target and the dense buckets sit around it. 1 s is already noticeable in the portal and 2.5 s is where the
Dentist clicks again, which is the retry case that the idempotency key must absorb. The completion returns `202` after
one MongoDB transaction, so normal values should stay under 200 ms.

Consumer buckets (Appointments, Billing): 0.05, 0.1, 0.5, 1, 5, 30 seconds, because nobody waits in front of them;
what matters is that progress and charges arrive before the Assistant finalizes the visit.

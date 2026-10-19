# CP2 - Structured log contract

Applies to `clinical-service`, `appointments-service` and `billing-service` for the journey
`POST /api/v1/procedures/{id}/completions` -> `ProcedureCompleted` -> `BillableItemsCreated`.
One JSON object per line on stdout. Field names are snake_case in every service.

## Required fields

| Field | Type | Req. | Example | Why it exists |
|---|---|---|---|---|
| `timestamp` | string, ISO-8601 UTC, ms | yes | `2026-10-19T14:31:07.215Z` | order events across services |
| `level` | string: ERROR, WARN, INFO, DEBUG | yes | `INFO` | filter by severity |
| `service` | string | yes | `clinical-service` | who wrote the line |
| `version` | string, semver or git sha | yes | `0.1.0` | tie a failure to a release |
| `environment` | string: dev, qa, main | yes | `qa` | same values as `APP_ENV` in `dlc-clinical-api` |
| `trace_id` | string, 32 hex | yes | `7a1c2e9f04b84d3e9c1f6a0b2d4e8f13` | join all lines of one request, from `traceparent` |
| `span_id` | string, 16 hex | yes | `b3f1a07c9d2e4a51` | the hop inside the trace |
| `correlation_id` | string, max 128 printable chars | yes | `7a1c2e9f04b84d3e9c1f6a0b2d4e8f13` | value of `X-Correlation-Id`, returned as `traceId` in the error envelope |
| `event` | string, `domain.action.result` | yes | `procedure.completion.accepted` | stable key for queries and alerts |
| `message` | string | yes | `procedure completion committed` | short human text, no personal data |
| `duration_ms` | integer | yes (0 when not timed) | `48` | latency per hop |

Today `dlc-clinical-api` (PR #11) writes `correlationId` and `durationMs`. They are renamed to `correlation_id` and
`duration_ms` in CP3, and `correlation_id` takes the `trace_id` value when no `X-Correlation-Id` arrives.

## Domain fields

| Field | Type | Req. | Example | Why it exists |
|---|---|---|---|---|
| `procedure_id` | uuid | when known | `5d0e8a7b-...` | the `procedures` document being completed; idempotency key of the closure |
| `appointment_id` | uuid | when known | `c41b2f60-...` | lets Appointments find the visit; it must be in `EN_ATENCION` |
| `care_closure_id` | uuid | clinical only | `9e7a3c15-...` | the `CareClosure` returned with `202`, followed with `GET /care-closures/{id}` |
| `procedure_code` | string, Billing catalog code | when known | `D2391` | Billing price lookup; a code is not clinical detail |
| `event_id` | uuid | messaging only | `0f6c1d2e-...` | `eventId` of `ProcedureCompleted` / `BillableItemsCreated`; `inbox_events` drops duplicates by it |
| `closure_status` | string: CLOSURE_PENDING, COMPLETED, CLOSURE_FAILED | clinical only | `CLOSURE_PENDING` | state of the closure after this step |

## Never log

| Data | Where it appears | Rule |
|---|---|---|
| Patient name, document number, email, phone, birth date | Patients data shown in the clinical portal | never log; `patient_id` only at DEBUG |
| Diagnosis text, clinical notes, `clinicalReason`, allergies, antecedents | `clinical_entries`, `diagnoses`, `procedure_extras` | never log; log the id and `reason_length` |
| JWT, refresh token, `Authorization` header | every request | never log headers; log the role and `actor_hash` (SHA-256 of `sub`) |
| Prices, totals, payment data | Billing charge snapshots, payments | never log amounts; log `charge_count` |
| `MONGODB_URI`, RabbitMQ credentials | config | printed as `***` if config is logged at startup |
| Raw request body | completion payload | never log; log only the invalid `field` names |

## Log levels

| Level | When | Example in this journey |
|---|---|---|
| ERROR | the request or message failed and needs a person or a retry policy | Billing cannot price a procedure and the closure moves to `CLOSURE_FAILED` |
| WARN | something failed but the system recovers on its own | RabbitMQ publish confirm timed out and the outbox relay retries |
| INFO | a business step finished | completion committed, `ProcedureCompleted` published, progress recorded |
| DEBUG | detail for local diagnosis, off in `main` | outbox polling batch size, Mongo transaction retries |

## Example: one trace

`trace_id` `7a1c2e9f04b84d3e9c1f6a0b2d4e8f13`. The completion is accepted, the first publish times out, the retry
works, Appointments records progress and Billing rejects the charge because the code has no catalog price.

```json
{"timestamp":"2026-10-19T14:31:07.215Z","level":"INFO","service":"clinical-service","version":"0.1.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"b3f1a07c9d2e4a51","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"procedure.completion.accepted","message":"procedure completion committed with outbox event","duration_ms":48,"procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05","appointment_id":"c41b2f60-7d8e-4a93-b5c1-0e6f9a2d3b74","care_closure_id":"9e7a3c15-6b2d-4f80-a4e9-1c5d7b0f2a68","procedure_code":"D2391","closure_status":"CLOSURE_PENDING"}
{"timestamp":"2026-10-19T14:31:07.731Z","level":"WARN","service":"clinical-service","version":"0.1.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"4c8e2b9a1f7d3e60","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"outbox.publish.timeout","message":"rabbitmq publish confirm not received, will retry","duration_ms":500,"event_id":"0f6c1d2e-8a3b-4c5d-9e7f-a1b2c3d4e5f6","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05"}
{"timestamp":"2026-10-19T14:31:09.744Z","level":"INFO","service":"clinical-service","version":"0.1.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"e17a5d3c9b2f4086","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"outbox.publish.sent","message":"ProcedureCompleted published","duration_ms":12,"event_id":"0f6c1d2e-8a3b-4c5d-9e7f-a1b2c3d4e5f6","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05"}
{"timestamp":"2026-10-19T14:31:09.802Z","level":"INFO","service":"appointments-service","version":"0.3.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"2a9f6c4e8d1b7035","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"procedure.progress.recorded","message":"procedure progress recorded for appointment","duration_ms":21,"event_id":"0f6c1d2e-8a3b-4c5d-9e7f-a1b2c3d4e5f6","appointment_id":"c41b2f60-7d8e-4a93-b5c1-0e6f9a2d3b74","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05"}
{"timestamp":"2026-10-19T14:31:09.815Z","level":"INFO","service":"billing-service","version":"0.3.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"8d3b1f7a5e9c2046","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"procedure.completed.received","message":"ProcedureCompleted received","duration_ms":0,"event_id":"0f6c1d2e-8a3b-4c5d-9e7f-a1b2c3d4e5f6","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05","procedure_code":"D2391"}
{"timestamp":"2026-10-19T14:31:09.839Z","level":"ERROR","service":"billing-service","version":"0.3.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"8d3b1f7a5e9c2046","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"charge.snapshot.rejected","message":"no catalog price and no justified manual price for procedure code","duration_ms":24,"event_id":"0f6c1d2e-8a3b-4c5d-9e7f-a1b2c3d4e5f6","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05","procedure_code":"D2391"}
{"timestamp":"2026-10-19T14:31:10.104Z","level":"ERROR","service":"clinical-service","version":"0.1.0","environment":"qa","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"f05c7e2a9b4d1368","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","event":"care.closure.failed","message":"billing outcome failed, closure needs reconciliation","duration_ms":9,"care_closure_id":"9e7a3c15-6b2d-4f80-a4e9-1c5d7b0f2a68","procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05","closure_status":"CLOSURE_FAILED"}
```

Failures shown: line 2 is the timeout to a downstream dependency (RabbitMQ publish confirm) and line 6 is the business
rule rejection (Billing never fabricates a zero charge, `dlc-docs` 07-api/design-decisions.md). Service versions are
placeholders until each service publishes a release tag.

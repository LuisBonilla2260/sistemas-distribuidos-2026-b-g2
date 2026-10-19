# How to run - clinical-service instrumentation

Minimal Go module for `clinical-service` (`dlc-clinical-api` is Go). Standard library only, no external dependencies.

| File | What it does |
|---|---|
| `trace.go` | `WithTrace` middleware: reads or creates W3C `traceparent`, returns it and `X-Correlation-Id` in the response |
| `logger.go` | `slog` JSON logger with the CP2 fields; injects `trace_id`, `span_id`, `correlation_id` |
| `propagate.go` | copies the trace to the `ProcedureCompleted` outbox headers and to outgoing HTTP calls |
| `metrics.go` | `clinical_procedure_completions_total{outcome}` and `clinical_procedure_completion_duration_seconds` on `/metrics` |
| `trace_test.go` | proves the incoming `trace_id` is the one sent to Billing, to the outbox and in the response |

## Install

Go 1.22 or newer.

```bash
cd 11-week/workshop/03-instrumentation/clinical-service
go version
```

## Run the tests

```bash
go vet ./...
go test -v ./...
```

## Real output (Windows 11, go1.26.7)

```text
=== RUN   TestTraceIDReceivedEqualsTraceIDSentDownstream
{"timestamp":"2026-10-19T15:06:03.115Z","level":"INFO","message":"procedure completion committed with outbox event","service":"clinical-service","version":"2.0.0","environment":"dev","event":"procedure.completion.accepted","duration_ms":0,"procedure_id":"5d0e8a7b-3c21-4f6e-9a10-2b7c4d8e1f05","procedure_code":"OPER-01","closure_status":"CLOSURE_PENDING","trace_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13","span_id":"f19be6cc44a09015","correlation_id":"7a1c2e9f04b84d3e9c1f6a0b2d4e8f13"}
--- PASS: TestTraceIDReceivedEqualsTraceIDSentDownstream (0.00s)
=== RUN   TestInvalidTraceparentStartsNewTrace
--- PASS: TestInvalidTraceparentStartsNewTrace (0.00s)
PASS
ok  	dlc/clinical-service-obs	4.482s
```

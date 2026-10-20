<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       11-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 11

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Start Corte 3 by turning dlc-clinical-api into a running, observable Go service - health endpoint, structured JSON logs with a propagated correlation/trace ID, a single error envelope, in-service JWT verification and the clinical domain model - and plan the SLIs/SLOs, alerts and runbook for the clinical service.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-011-001 | Bootstrap the clinical-api Go service: validated config, public `/health`, distroless image, compose on the `platform` network, CI workflow | doing | [dlc-clinical-api PR #10](https://github.com/code-corhuila/dlc-clinical-api/pull/10) |
| HU-011-002 | Structured JSON logs with a propagated `X-Correlation-Id` used as `traceId`, plus a single error envelope | doing | [dlc-clinical-api PR #11](https://github.com/code-corhuila/dlc-clinical-api/pull/11) |
| HU-011-003 | Verify RS256 JWTs inside the clinical service (issuer/audience allowlists, `401` with envelope and `traceId`) | doing | [dlc-clinical-api PR #12](https://github.com/code-corhuila/dlc-clinical-api/pull/12) |
| HU-011-004 | Clinical domain model (record, append-only entries and amendments, treatments) with an enforced dependency rule | doing | [dlc-clinical-api PR #13](https://github.com/code-corhuila/dlc-clinical-api/pull/13) |
| HU-011-005 | Golden-signal metrics (rate, errors, latency) and one distributed trace across a multi-service request | todo | - |
| HU-011-006 | Clinical service SLOs, error-budget policy, a symptom-based alert and one runbook | todo | - |

## 2. My individual contribution

### Session 1 - Observability (logs, metrics, traces)
- **Correlation ID propagated at the edge:** the middleware reuses an incoming `X-Correlation-Id` when it is safe (at most 128 printable characters) or generates a UUID. It echoes the ID in the response, returns it as `traceId` in every error body and writes it on every log line (PR #11).
- **Structured logs:** JSON access log per request (no headers or tokens logged), panic recovery, and unknown errors logged with detail while the client only gets a neutral `500 INTERNAL_ERROR` (PR #11).
- **Consistent error envelope:** `{error, message, details?, traceId}` built in one place; router 404/405 and malformed or unknown-field JSON (`400 VALIDATION_ERROR`) also use it, so every failure can be traced back to its logs (PR #11).
- **Health signal:** public `GET /health` returning only `{"status":"UP"}`, checked in the container smoke test (PR #10).
- **Explicit server limits:** every timeout is declared and validated; a zero or negative duration makes the service refuse to start, because zero means "no limit" in Go (PR #10).
- **Security, traceable on failure:** in-service RS256 JWT verification; `401` responses carry the envelope and the echoed `traceId` (PR #12).
- **Domain first:** pure clinical domain model, with `architecture_test.go` blocking `net/http`, database drivers and `log/slog` from the domain (PR #13).
- All four PRs were written test-first (RED -> GREEN); `gofmt`, `go vet` and `go test` pass locally.

### Session 2 - Planning (SLIs, SLOs, error budget, alerting)
- Proposed SLIs for the clinical service, built on the signals from PR #11:
  - **Availability:** share of non-5xx responses.
  - **Latency:** p95 of clinical write requests.
  - **Freshness:** for the future care-completion -> billing event flow, how far behind its consumer may run.
- Not yet written down as SLOs with windows, error-budget policy, alert or runbook (HU-011-006).

### Workshop delivery 1 (2026-10-19) - journey "complete a procedure"
Critical journey `POST /api/v1/procedures/{id}/completions` -> MongoDB outbox -> `ProcedureCompleted` -> Appointments and Billing, in [11-week/workshop](../workshop/README.md).
- **CP0 - setup:** workshop README with CONFIG ([754f938](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/754f938)).
- **CP1 - journey and trace map:** hop table, Mermaid sequence with `traceparent` per arrow, span tree across the async boundary; aligned with the `dlc-clinical-db` collections ([63ebf81](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/63ebf81), [810d84a](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/810d84a)).
- **CP2 - log contract:** required and domain fields, never-log list, levels and seven JSON lines of one trace with a publish timeout and a billing rejection ([d267c6a](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/d267c6a), [04b89cc](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/04b89cc)).
- **CP3 - instrumentation in Go:** `traceparent` middleware, `slog` JSON logger, propagation to the outbox headers and to outgoing HTTP, `/metrics` with a counter and a latency histogram, and a passing test that the trace id is kept downstream ([abed306](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/abed306), [6322a0a](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/6322a0a), [6ed2cb1](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/6ed2cb1)).
- **CP4 - metric catalog:** golden signals for clinical, appointments and billing with PromQL, label cardinality and latency buckets ([b210632](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/b210632)).
- HU-011-005 advanced in the workshop module only; `dlc-clinical-api` still has no `/metrics` or `traceparent`.

## 3. Blockers and risks
- **CI not executed:** the organization's GitHub Actions account is locked by billing, so `ci.yml` never starts on PRs #10-#13. Quality is only proven by local runs.
- **Stacked PRs waiting for review:** #11-#13 build on #10. Nothing reaches `develop` until #10 is approved.
- **No metrics or traces yet:** only logs plus correlation exist. There is no metrics endpoint and no OpenTelemetry tracing, so the golden signals cannot be shown on a dashboard and no cross-service trace exists.
- **No business endpoints:** no `clinical-service.yaml` operation is implemented yet, so there is no real multi-service request to trace. Persistence still depends on `dlc-infra-mongo` plus the `dlc-clinical-db` v2.0.0 schema.

## 4. Plan for next week
- Get #10-#13 merged into `develop` and promote them to `qa`.
- Implement the first clinical use case (open record / append entry) against MongoDB.
- Add RED metrics (rate, errors, duration) and OpenTelemetry tracing that propagates the trace ID to the gateway and the clinical portal.
- Write the clinical SLOs (availability and latency over 30 days), the error-budget policy, one burn-rate alert with owner and severity, and a runbook for "clinical writes failing".

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [dlc-clinical-api PR #10 - bootstrap go service skeleton](https://github.com/code-corhuila/dlc-clinical-api/pull/10)
- [dlc-clinical-api PR #11 - error envelope and correlation](https://github.com/code-corhuila/dlc-clinical-api/pull/11)
- [dlc-clinical-api PR #12 - verify RS256 tokens in the service](https://github.com/code-corhuila/dlc-clinical-api/pull/12)
- [dlc-clinical-api PR #13 - clinical domain model](https://github.com/code-corhuila/dlc-clinical-api/pull/13)
- [dlc-docs issue #57 - HU-CLN-001](https://github.com/code-corhuila/dlc-docs/issues/57)
- [dlc-clinical-db v2.0.0 release PR #13 (schema owner)](https://github.com/code-corhuila/dlc-clinical-db/pull/13)

<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       10-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 10

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Ship MVP 2 - promote the clinical repositories and the dlc-front compositor develop -> qa -> main and tag v2.0.0, with persistence ownership, outbox/idempotency contracts and idempotent clinical writes documented and evidenced.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-010-001 | Document outbox publication, idempotent consumers and payment concurrency for billing and patients | done | [dlc-docs PR #92](https://github.com/code-corhuila/dlc-docs/pull/92) |
| HU-010-002 | Version billing and patient event contracts and the billing OpenAPI contract | done | [dlc-docs PR #89](https://github.com/code-corhuila/dlc-docs/pull/89), [PR #90](https://github.com/code-corhuila/dlc-docs/pull/90), [PR #91](https://github.com/code-corhuila/dlc-docs/pull/91) |
| HU-010-003 | Clinical record workspace with idempotent writes and visible processing/retry states for care closure | done | [dlc-clinical-portal PR #46](https://github.com/code-corhuila/dlc-clinical-portal/pull/46), [PR #51](https://github.com/code-corhuila/dlc-clinical-portal/pull/51) |
| HU-010-004 | Expose the clinical portal through composition contract v1 and integrate it into the dlc-front compositor | done | [dlc-clinical-portal PR #57](https://github.com/code-corhuila/dlc-clinical-portal/pull/57), [dlc-front PR #69](https://github.com/code-corhuila/dlc-front/pull/69) |
| HU-010-005 | MVP 2 release: promote develop -> qa -> main, CHANGELOG and tag v2.0.0 for the clinical repositories and dlc-front | done | [dlc-clinical-portal PR #107](https://github.com/code-corhuila/dlc-clinical-portal/pull/107), [dlc-front PR #96](https://github.com/code-corhuila/dlc-front/pull/96) |

## 2. My individual contribution

### Session 1 - Persistence in distributed systems (saga, outbox, CQRS)
- **Database per service:** the Clinical context owns its MongoDB logical database `clinical` through Liquibase migrations (dlc-clinical-db, released as v2.0.0); no other service reads it directly.
- **Outbox and idempotent consumers (design):** documented that Billing and Patients publish committed facts through their own outbox after the local state/outbox transaction commits, and that RabbitMQ at-least-once delivery is handled by consumers with inbox/business-source deduplication, bounded retries and dead-letter queues (dlc-docs PR #92).
- **Concurrency and idempotency:** added the payment concurrency sequence diagram (idempotency key reservation, version check, payment + outbox in one transaction, `409 IDEMPOTENCY_CONFLICT` on replay) (dlc-docs PR #92).
- **Versioned event contracts:** `InvoicePaid`, `ManualChargeRecorded`, `PaymentRecorded`, `PatientRegistered`, `PatientProfileUpdated` and `PatientDeactivated` v1 JSON schemas, plus the aligned billing OpenAPI contract (dlc-docs PRs #89-#91).
- **Idempotent writes from the client:** every clinical write (entries, plan, start, completion, care declaration and closure retry) sends an idempotency/intent key, so a retried request cannot create a duplicate clinical effect (dlc-clinical-portal PRs #46, #51, #91).
- **Eventual consistency in the UI:** the care-closure flow shows pending/processing states and a retry action instead of assuming instant consistency (dlc-clinical-portal care closure follow-up).

### Session 2 - MVP 2 release
- **Per-environment promotion:** delivered every change as a feature branch -> `develop`, promoted in batches to `qa` (`promotion/*` PRs) and released to `main` through `release/2.0.0` PRs approved by the instructor.
- **Tagged and versioned:** `v2.0.0` tags on `main` with a CHANGELOG for dlc-clinical-api, dlc-clinical-db, dlc-clinical-portal and dlc-front.
- **Integrated system:** the clinical portal mounts inside the dlc-front compositor (registry, lifecycle, shell, session and shared HTTP client) next to the patients, billing and appointments portals; `docker compose` brings up the compositor, the portal remotes and the Prism gateway mock behind health checks.
- **Quality:** lint, format, coverage, build and PR gates run on the portal and the compositor; security advisory remediated.
- **Traceability:** all commits and PRs with hashes are recorded in [TRACEABILITY.md](../TRACEABILITY.md).

## 3. Blockers and risks
- **No saga or outbox implemented yet:** the clinical portal runs on synthetic data and the Prism mock; dlc-clinical-api v2.0.0 is a repository baseline with no operations implemented. The outbox/idempotency work this week is contracts and design, not running code.
- **No failure path demonstrated:** a cross-service flow with an injected failure and a saga compensation could not be shown because the clinical API and broker integration do not exist yet.
- **No CQRS read model:** the clinical dashboard is a client-side view over demo data, not a projection built from events.
- dlc-clinical-portal has a follow-up release PR (#110, `2.0.1`) still open.

## 4. Plan for next week
- Implement the first clinical endpoints in dlc-clinical-api with the outbox for one critical event (care completion) and an idempotent consumer.
- Define the care-completion -> billing saga with its compensation and prove the failure path in a demo.
- Run the Corte 2 retrospective and turn integration pain (mocked API, contract mismatches) into Corte 3 backlog items: observability, deployment, security and CI/CD.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Contribution traceability: all commits and PRs with hashes](../TRACEABILITY.md)
- [Corte 2 release tags](../c2-release/RELEASES.md)
- [dlc-docs PR #92 - outbox, idempotent consumers and payment concurrency](https://github.com/code-corhuila/dlc-docs/pull/92)
- [dlc-docs PR #91 - billing OpenAPI contract](https://github.com/code-corhuila/dlc-docs/pull/91)
- [dlc-docs PR #89 - payment, charge and patient event contracts](https://github.com/code-corhuila/dlc-docs/pull/89)
- [dlc-clinical-portal PR #46 - idempotency keys for clinical writes](https://github.com/code-corhuila/dlc-clinical-portal/pull/46)
- [dlc-clinical-portal PR #51 - intent keys for care declaration and closure retry](https://github.com/code-corhuila/dlc-clinical-portal/pull/51)
- [dlc-clinical-portal PR #107 - release 2.0.0 to main](https://github.com/code-corhuila/dlc-clinical-portal/pull/107)
- [dlc-clinical-db PR #13 - release 2.0.0 to main](https://github.com/code-corhuila/dlc-clinical-db/pull/13)
- [dlc-clinical-api PR #9 - release 2.0.0 to main](https://github.com/code-corhuila/dlc-clinical-api/pull/9)
- [dlc-front PR #96 - release 2.0.0 to main](https://github.com/code-corhuila/dlc-front/pull/96)
- [dlc-clinical-portal v2.0.0 release](https://github.com/code-corhuila/dlc-clinical-portal/releases/tag/v2.0.0)
- [dlc-front v2.0.0 release](https://github.com/code-corhuila/dlc-front/releases/tag/v2.0.0)

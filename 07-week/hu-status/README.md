<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       07-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 07

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Provide technical support for inter-service communication, versioned contracts, and safe integration planning for MVP 2.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-007-001 | Document inter-service communication and versioned contract support for MVP 2 | done | [Communication](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/572f866), [contracts](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/bab9e3e), and [delivery backlog](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/56fffb5) |

## 2. My individual contribution
- Produced an English support package covering synchronous REST/gRPC calls, asynchronous messaging, delivery semantics, and idempotent consumer design.
- Defined a decision matrix to select REST, gRPC, event topics, or queues according to interaction requirements and failure behavior.
- Documented machine-readable contract standards, compatibility rules, and the shared API error envelope.
- Proposed testable MVP 2 stories for contract publication, compatible evolution, and consumer-provider verification without claiming implementation in the Di-Lucca application.

## 3. Blockers and risks
- The current MVP communication boundaries and existing contracts must be confirmed in the application repository before any implementation is claimed.
- Contract testing requires an agreed consumer-provider pair and CI integration; this support package does not add a Pact test.
- At-least-once delivery requires idempotent consumers and operational visibility for retries or dead-letter processing.

## 4. Plan for next week
- Review the next weekly material and prioritize only support work that the team validates as necessary.
- Keep documentation aligned with confirmed project evidence and avoid unverified implementation commitments.

## 5. Compliance self-check
- [ ] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [ ] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Communication decisions](../interservice-contracts-support/communication-decisions.md)
- [Versioned contract standards](../interservice-contracts-support/contract-standards.md)
- [Delivery semantics and idempotency](../interservice-contracts-support/delivery-idempotency.md)
- [MVP 2 contracts and testing backlog](../interservice-contracts-support/mvp2-contract-testing-backlog.md)
- [Communication decisions commit](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/572f866)
- [Contract standards commit](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/bab9e3e)
- [Delivery support backlog commit](https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/56fffb5)

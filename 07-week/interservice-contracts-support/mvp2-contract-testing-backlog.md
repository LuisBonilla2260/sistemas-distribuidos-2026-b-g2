# MVP 2 contracts and testing backlog

This is a proposed implementation backlog. It does not claim that these contracts or tests already exist in the MVP.

## HU-007-001 — Publish a versioned service contract

**As an integration consumer, I want a machine-readable contract so that I can implement against an explicit, reviewable API promise.**

Acceptance criteria:

- The selected REST, gRPC, or event contract is stored and versioned with the relevant service or approved documentation repository.
- The contract declares requests or events, successful responses or payloads, and standard errors where applicable.
- Contract examples use no real secrets or production data.
- The contract passes its chosen linter or validator in CI.

## HU-007-002 — Preserve consumers through compatible evolution

**As a service team, I want compatibility rules applied to contract changes so that independently deployed consumers are not silently broken.**

Acceptance criteria:

- Every change is classified as backward-compatible or breaking during review.
- Breaking changes introduce a new version and a documented migration/deprecation path.
- Existing consumers remain supported until the agreed migration window ends.

## HU-007-003 — Verify one consumer-provider promise

**As a consumer team, I want a contract test verified by the producer in CI so that incompatible changes fail before deployment.**

Acceptance criteria:

- A consumer publishes one expectation using Pact or an equivalent approved contract-testing tool.
- The producer verifies that expectation in CI.
- A deliberate incompatible change causes the verification to fail.
- The verification result is retained as build evidence.

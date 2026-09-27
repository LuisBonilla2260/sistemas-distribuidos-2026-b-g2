<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       08-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 08

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Align MVP 2 patient ownership, authorization, and traceability documentation with validated project evidence.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-008-001 | Align MVP 2 patient ownership, authorization, and traceability documentation | done | [Patient ownership and deactivation rules](https://github.com/code-corhuila/dlc-docs/commit/ff21735) |

## 2. My individual contribution
- Consolidated evidence for Patients as a business service and Auth/IAM as a transversal service in the MVP 2 documentation baseline.
- Updated the documented scope of `HU-PAT-001`, including patient administration, role boundaries, protected deactivation, and history preservation.
- Added traceability for patient ownership rules and their verification expectations.
- Aligned the Patients OpenAPI contract and endpoint index with the documented ownership and authorization boundaries.
- Preserved the distinction between documented design decisions and implementation evidence; no executable service or automated test is claimed by this contribution.

## 3. Blockers and risks
- The documented Patients boundary, authorization rules, and deactivation safeguards still require implementation and executable verification.
- Protected deactivation depends on authoritative appointment-state validation and a reviewed concurrency protocol.
- Documentation changes must remain synchronized with contracts, implementation, and tests as MVP 2 evolves.

## 4. Plan for next week
- Review the upcoming weekly material and support only validated team progress.
- Keep the fork evidence aligned with confirmed commits and avoid unverified implementation claims.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [ ] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Product update introducing Patients](https://github.com/code-corhuila/dlc-docs/commit/0ff234a)
- [Patient ownership and deactivation rules](https://github.com/code-corhuila/dlc-docs/commit/ff21735)
- [PR #25 merge evidence](https://github.com/code-corhuila/dlc-docs/commit/7a6e093)
- [Patients and transversal IAM ADR](https://github.com/code-corhuila/dlc-docs/blob/main/05-architecture/decisions/records/ADR-003-patients-and-transversal-iam.md)

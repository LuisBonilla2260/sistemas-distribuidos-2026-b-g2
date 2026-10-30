<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       13-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 13

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Threat-model the api-gateway -> clinical-service boundary and turn its highest risks into security stories, CI gates and a negative authorization test for clinical-service.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| dlc-docs#137 | [security] Only the assigned dentist can read or complete a clinical record (E, risk 9, P0) | doing | https://github.com/code-corhuila/dlc-docs/issues/137 |
| dlc-docs#138 | [security] clinical-service accepts only a verified JWT, never identity headers (S, risk 6, P1) | todo | https://github.com/code-corhuila/dlc-docs/issues/138 |
| dlc-docs#139 | [security] Clinical entries keep server-set author, time and original text (T, risk 6, P1) | todo | https://github.com/code-corhuila/dlc-docs/issues/139 |

## 2. My individual contribution
- Claimed the trust boundary api-gateway -> clinical-service (inbound): https://github.com/code-corhuila/dlc-docs/issues/130
- Mapped the DLC attack surface (TB1-TB6, assets, entry points): `13-week/workshop/01-attack-surface.md`
- STRIDE analysis with risk ranking and library references (Pressman, 6th ed., 005.3 P727): https://github.com/code-corhuila/dlc-docs/pull/133
- Three security stories in the team backlog: #137, #138, #139
- Security workflow (gitleaks, govulncheck, authz tests) and a negative authorization test for `GET /api/v1/clinical-records/{id}`: `13-week/workshop/04-security-ci/`
- Review of the Billing STRIDE of @juan0898 (row E, missing `POST /api/v1/manual-charges`): https://github.com/code-corhuila/dlc-docs/pull/132#pullrequestreview-5700321494
- CP6: answered the review of @juan0898 on PR #133; row E now lists three negative tests (no active assignment, ended assignment, Appointments down) and a Review log was added: https://github.com/code-corhuila/dlc-docs/pull/133#issuecomment-6501351650
- CP7: refined #137 (six testable criteria, two from the review) and linked #138 with #126 (Auth) and #135 (Billing); the three stories are in the team board "Di Lucca — Backlog"
- CP8: three negative authorization tests run against `dlc-clinical-api` (develop, 12438c4) and pass: unassigned dentist 404, assistant 403, Appointments down 503; `security.yml` validated with actionlint v1.7.7 (0 errors)
- CP9: implementation plan written in the reflection document

## 3. Blockers and risks
- The assignment check of clinical-service depends on the Appointments internal contract, not validated at runtime yet.
- `security.yml` is a proposal in this fork; it is not in `dlc-clinical-api` yet.

## 4. Plan for next week
- Implement #137 in `dlc-clinical-api`: current assignment check on every request and fail closed when Appointments does not answer.
- Propose `security.yml` and the negative tests as a pull request to `dlc-clinical-api`.
- Then #138 (coordinated with #126 and #135) and #139.
- Defend the workshop in class.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- Workshop folder: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/tree/main/13-week/workshop
- CP0: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/b5dfcba
- CP1: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/c548db6
- CP3: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/d4c5d58
- CP4: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/436a620
- CP5: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/7a2d567
- CP6: https://github.com/code-corhuila/dlc-docs/commit/48486e2
- CP7: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/c989ecb
- CP8: https://github.com/LuisBonilla2260/sistemas-distribuidos-2026-b-g2/commit/faf16cf
- STRIDE PR: https://github.com/code-corhuila/dlc-docs/pull/133
- Review comment: https://github.com/code-corhuila/dlc-docs/pull/132#pullrequestreview-5700321494
- Reflection: https://docs.google.com/document/d/1n3qOHYYvCyfAkfuE8M9zqQtdTRCtJVat5t9cUdIdxyU/edit?usp=sharing

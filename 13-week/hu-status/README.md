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

## 3. Blockers and risks
- The assignment check of clinical-service depends on the Appointments internal contract, not validated at runtime yet.
- The three security issues still need to be added to the team GitHub Project.
- `security.yml` is a proposal in this fork; it is not in `dlc-clinical-api` yet.

## 4. Plan for next week
- Answer the review on PR #133 and finish the STRIDE (CP6).
- Refine the three security stories (CP7).
- Run the negative test against clinical-service and validate `security.yml` with actionlint (CP8).
- Write the implementation plan (CP9).

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
- STRIDE PR: https://github.com/code-corhuila/dlc-docs/pull/133
- Review comment: https://github.com/code-corhuila/dlc-docs/pull/132#pullrequestreview-5700321494
- Reflection: https://docs.google.com/document/d/1n3qOHYYvCyfAkfuE8M9zqQtdTRCtJVat5t9cUdIdxyU/edit?usp=sharing

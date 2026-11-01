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

### Session 1 - Security in distributed systems (authn, authz, zero trust)
Already merged in `dlc-clinical-api`:
- **Zero trust: the service verifies every token itself.** It does not rely on the gateway. Only RS256 is accepted, checked against the identity public key, with issuer/audience allowlists and mandatory `exp`/`sub`. A failure returns `401` with the error envelope and `traceId` ([PR #12](https://github.com/code-corhuila/dlc-clinical-api/pull/12)).
- **Authorization on the owner's side:** the principal's roles and permissions travel in the request context. Ownership checks call the Patients and Appointments owner contracts instead of trusting the caller ([PR #40](https://github.com/code-corhuila/dlc-clinical-api/pull/40)).
- **Secrets out of code:** the MongoDB runtime password is read from a secret file, not from code or a plain variable ([PR #118](https://github.com/code-corhuila/dlc-clinical-api/pull/118) -> qa [#119](https://github.com/code-corhuila/dlc-clinical-api/pull/119)). Runtime database roles have no DDL rights ([dlc-clinical-db PR #22](https://github.com/code-corhuila/dlc-clinical-db/pull/22)), and each stage migrates over TLS ([dlc-clinical-db PR #27](https://github.com/code-corhuila/dlc-clinical-db/pull/27)).
- **Least privilege in `dlc-appointments-db`:**
  - Separate `appointments_migrator` and `appointments_runtime` roles: no superuser, createrole or bypassrls rights, and no passwords in the repo ([PR #26](https://github.com/code-corhuila/dlc-appointments-db/pull/26)).
  - Database `CREATE` is revoked once the schema exists, because it would also allow installing trusted extensions; a test found this with `pgcrypto`.
  - The runtime role gets DML-only grants ([PR #16](https://github.com/code-corhuila/dlc-appointments-db/pull/16)), and the clinical assignment table has bounded grants and no cross-owner foreign keys ([PR #19](https://github.com/code-corhuila/dlc-appointments-db/pull/19)).
- **Encryption in transit:** the Clinical migration runner refuses `qa` and `prod` without the stage CA, and connects as `clinical_migrator` with the password from a secret ([dlc-clinical-db PR #27](https://github.com/code-corhuila/dlc-clinical-db/pull/27)).
- **Documented in `dlc-docs`:** stage deployment and runtime connection records for Clinical ([#143](https://github.com/code-corhuila/dlc-docs/pull/143)) and Appointments ([#141](https://github.com/code-corhuila/dlc-docs/pull/141)), the Appointments runtime-grant correction ([#142](https://github.com/code-corhuila/dlc-docs/pull/142)) and the DB-04 to DB-07 persistence contracts ([#140](https://github.com/code-corhuila/dlc-docs/pull/140)).
- **Input validation and safe errors:** malformed or unknown-field JSON returns `400`. Unknown errors become a neutral `500` with the detail only in the log ([PR #11](https://github.com/code-corhuila/dlc-clinical-api/pull/11)).
- **Contract-compliant identifiers:** idempotency keys and correlation IDs are UUIDs, as the contracts declare ([PR #116](https://github.com/code-corhuila/dlc-clinical-api/pull/116) -> qa [#117](https://github.com/code-corhuila/dlc-clinical-api/pull/117)).

### Session 2 - Threat modeling and security backlog (workshop)
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
- STRIDE PR [dlc-docs #133](https://github.com/code-corhuila/dlc-docs/pull/133) is approved but not merged yet; #137-#139 are still open.

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
- Zero-trust token verification in clinical-api: https://github.com/code-corhuila/dlc-clinical-api/pull/12
- Secret-file database password in clinical-api: https://github.com/code-corhuila/dlc-clinical-api/pull/118
- Owner roles and least-privilege bootstrap in appointments-db: https://github.com/code-corhuila/dlc-appointments-db/pull/26
- TLS stage migration in clinical-db: https://github.com/code-corhuila/dlc-clinical-db/pull/27
- Clinical stage deployment record: https://github.com/code-corhuila/dlc-docs/pull/143
- Review comment: https://github.com/code-corhuila/dlc-docs/pull/132#pullrequestreview-5700321494
- Reflection: https://docs.google.com/document/d/1n3qOHYYvCyfAkfuE8M9zqQtdTRCtJVat5t9cUdIdxyU/edit?usp=sharing

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
| HU-013-004 | Real api-gateway as the edge choke point: correlation IDs (UUID v4), CORS limited to the front origin, credential presence filter, login and per-IP rate limits, client address only from the trusted front proxy | done | [dlc-api-gateway PRs #3-#16](https://github.com/code-corhuila/dlc-api-gateway/pulls?q=is%3Apr+author%3ALuisBonilla2260), [dlc-docs #157](https://github.com/code-corhuila/dlc-docs/pull/157), [#158](https://github.com/code-corhuila/dlc-docs/pull/158), [#159](https://github.com/code-corhuila/dlc-docs/pull/159), [#163](https://github.com/code-corhuila/dlc-docs/pull/163) |
| HU-013-005 | Real auth session in dlc-front against the gateway (cookie renewal, cross-tab sync, logout and protected-401 handling), Prism mock and auth double removed | done | [dlc-front PRs #101-#112](https://github.com/code-corhuila/dlc-front/pulls?q=is%3Apr+author%3ALuisBonilla2260), [#124](https://github.com/code-corhuila/dlc-front/pull/124), [#126](https://github.com/code-corhuila/dlc-front/pull/126), [#128](https://github.com/code-corhuila/dlc-front/pull/128) |
| HU-013-006 | Billing service with in-service JWT and private IAM session validation, administrator-only price creation, idempotent invoices, outbox, payments and protected PDF | done | [dlc-billing-api PRs #11-#52](https://github.com/code-corhuila/dlc-billing-api/pulls?q=is%3Apr+author%3ALuisBonilla2260), [dlc-billing-db PRs #9-#18](https://github.com/code-corhuila/dlc-billing-db/pulls?q=is%3Apr+author%3ALuisBonilla2260) |
| HU-013-007 | dlc-worker reminders and invoice/payment emails, calling owner contracts with service tokens; own queues, DLQ and inbox | done | [dlc-worker PRs #4-#29](https://github.com/code-corhuila/dlc-worker/pulls?q=is%3Apr+author%3ALuisBonilla2260), [dlc-docs #155](https://github.com/code-corhuila/dlc-docs/pull/155), [#168](https://github.com/code-corhuila/dlc-docs/pull/168) |
| HU-013-008 | Owner authorization contracts and real portals: IAM staff validation and dentist directory, Patients private billing contact, Clinical admin read-only; appointments, billing and clinical portals on the real APIs | done | [dlc-iam-api #88](https://github.com/code-corhuila/dlc-iam-api/pull/88), [#89](https://github.com/code-corhuila/dlc-iam-api/pull/89), [dlc-patient-api #50](https://github.com/code-corhuila/dlc-patient-api/pull/50), [dlc-clinical-api #124](https://github.com/code-corhuila/dlc-clinical-api/pull/124), [dlc-appointments-portal #94-#108](https://github.com/code-corhuila/dlc-appointments-portal/pulls?q=is%3Apr+author%3ALuisBonilla2260), [dlc-billing-portal #114-#116](https://github.com/code-corhuila/dlc-billing-portal/pulls?q=is%3Apr+author%3ALuisBonilla2260), [dlc-clinical-portal #132-#136](https://github.com/code-corhuila/dlc-clinical-portal/pulls?q=is%3Apr+author%3ALuisBonilla2260) |
| HU-013-009 | Register Billing portal 2.1.0 in the compositor | doing | [dlc-front PR #116](https://github.com/code-corhuila/dlc-front/pull/116) |

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

### Integration week (Nov 1-8) - zero trust applied to the real system
- **Gateway as choke point, not the only guard:**
  - the edge checks that a credential is present and rate-limits login and general traffic per client IP;
  - CORS only accepts the front origin;
  - the client address is trusted only from the front proxy ([dlc-api-gateway #5](https://github.com/code-corhuila/dlc-api-gateway/pull/5), [#6](https://github.com/code-corhuila/dlc-api-gateway/pull/6), [#7](https://github.com/code-corhuila/dlc-api-gateway/pull/7), [#14](https://github.com/code-corhuila/dlc-api-gateway/pull/14)).
- **Every service still verifies the token:**
  - Billing enforces JWT authorization and private IAM session eligibility ([#12](https://github.com/code-corhuila/dlc-billing-api/pull/12), [#24](https://github.com/code-corhuila/dlc-billing-api/pull/24)), tested with RSA-signed tokens ([#13](https://github.com/code-corhuila/dlc-billing-api/pull/13));
  - price creation requires administrator access ([#16](https://github.com/code-corhuila/dlc-billing-api/pull/16));
  - IAM validates staff state and roles for owner services ([dlc-iam-api #89](https://github.com/code-corhuila/dlc-iam-api/pull/89)).
- **Broken access control prevented:** Clinical lets administrators read but never write clinical data ([dlc-clinical-api #124](https://github.com/code-corhuila/dlc-clinical-api/pull/124)).
- **Service-to-service trust:** dlc-worker calls the Appointments and Patients private contracts with service tokens ([dlc-worker #9](https://github.com/code-corhuila/dlc-worker/pull/9)). Patients exposes only a private billing contact for Billing and the Worker ([dlc-patient-api #50](https://github.com/code-corhuila/dlc-patient-api/pull/50)).
- **Least privilege in data:** the Billing DB has least-privilege roles and scoped grants ([dlc-billing-db #9](https://github.com/code-corhuila/dlc-billing-db/pull/9), [#10](https://github.com/code-corhuila/dlc-billing-db/pull/10)). Billing also runs non-root, as a test requires ([#23](https://github.com/code-corhuila/dlc-billing-api/pull/23)).
- **Session security in the browser:** cookie renewal is serialized across tabs. Restoration is blocked after logout or a protected 401, and the development auth double is removed ([dlc-front #103](https://github.com/code-corhuila/dlc-front/pull/103), [#104](https://github.com/code-corhuila/dlc-front/pull/104), [#110](https://github.com/code-corhuila/dlc-front/pull/110)).

### Coverage of the week's checklist
| Session asks for | Where it is |
|---|---|
| Authenticate at the edge | api-gateway credential presence filter, login rate limit, real auth session in dlc-front (HU-013-004, HU-013-005) |
| Validate tokens in each service (stateless signature + expiry) | clinical-api RS256 verification (PR #12), Billing JWT + IAM session eligibility (#12, #24) |
| Authorize every protected action | assigned-dentist check (#137), admin-only prices (Billing #16), admin read-only in Clinical (#124); negative authz tests 404/403/503 |
| Secure service-to-service calls | short-lived service tokens from dlc-worker to Appointments/Patients (dlc-worker #9); no mTLS yet |
| Least privilege and secrets | owner/runtime DB roles (appointments-db #26, billing-db #9-#10), secret-file passwords, TLS migrations (clinical-db #27) |
| Validate inputs, no verbose errors | strict JSON decoding and neutral `500` in clinical-api (#11) |
| Hash secrets | user passwords are owned by the IAM/auth service; no hashing work of mine this week |
| Threat model (data-flow, trust boundaries, STRIDE) | attack surface TB1-TB6 and STRIDE for api-gateway -> clinical-service (dlc-docs #133) |
| Prioritized security backlog with testable criteria | dlc-docs #137 (P0), #138 (P1), #139 (P1) |
| Security gates in CI (deps, secrets, authz) | `security.yml` with gitleaks, govulncheck and authz tests (proposal in this fork, not yet in dlc-clinical-api) |

## 3. Blockers and risks
- The assignment check of clinical-service depends on the Appointments internal contract, not validated at runtime yet.
- `security.yml` is a proposal in this fork; it is not in `dlc-clinical-api` yet.
- STRIDE PR [dlc-docs #133](https://github.com/code-corhuila/dlc-docs/pull/133) is approved but not merged yet; #137-#139 are still open.
- dlc-front PR #116 (Billing portal 2.1.0) is still open. dlc-billing-api #15 and dlc-billing-db #8 were closed without merge; later PRs redid that scope.

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

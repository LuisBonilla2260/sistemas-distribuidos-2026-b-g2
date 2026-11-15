<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       14-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 14

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Ship release 3.0.0 of the whole Di Lucca system through the develop -> qa -> main pipeline, with a `ci.yml` quality gate in every repository, IAM-issued service tokens between services (ADR-020) and the last scheduling stories (waiting list, appointment types) closed.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-014-001 | Release 3.0.0: changelog -> qa -> `main` and tag `v3.0.0` in every component repository | done | [dlc-clinical-api #133](https://github.com/code-corhuila/dlc-clinical-api/pull/133), [dlc-front #135](https://github.com/code-corhuila/dlc-front/pull/135), [dlc-api-gateway #26](https://github.com/code-corhuila/dlc-api-gateway/pull/26), [dlc-worker #41](https://github.com/code-corhuila/dlc-worker/pull/41), [dlc-docs #175](https://github.com/code-corhuila/dlc-docs/pull/175) |
| HU-014-002 | CI pull request workflow for dlc-infra-mongo (Compose per environment, replica set and transactions, network allow-list, TLS required, backup/restore) - TD-003 | done | [dlc-infra-mongo #17](https://github.com/code-corhuila/dlc-infra-mongo/pull/17), [#19](https://github.com/code-corhuila/dlc-infra-mongo/pull/19) |
| HU-014-003 | IAM-issued service tokens by client credentials, adopted by every service that calls another (ADR-020) | done | [dlc-docs #171](https://github.com/code-corhuila/dlc-docs/pull/171), [dlc-iam-api #114](https://github.com/code-corhuila/dlc-iam-api/pull/114), [dlc-iam-db #31](https://github.com/code-corhuila/dlc-iam-db/pull/31), [clinical-api #128](https://github.com/code-corhuila/dlc-clinical-api/pull/128), [billing-api #55](https://github.com/code-corhuila/dlc-billing-api/pull/55), [patient-api #52](https://github.com/code-corhuila/dlc-patient-api/pull/52), [appointments-api #87](https://github.com/code-corhuila/dlc-appointments-api/pull/87), [worker #33](https://github.com/code-corhuila/dlc-worker/pull/33) |
| HU-APT-005 | Waiting list | done | [dlc-docs #177](https://github.com/code-corhuila/dlc-docs/issues/177), [appointments-api #90](https://github.com/code-corhuila/dlc-appointments-api/pull/90), [appointments-db #34](https://github.com/code-corhuila/dlc-appointments-db/pull/34), [#35](https://github.com/code-corhuila/dlc-appointments-db/pull/35), [appointments-portal #112](https://github.com/code-corhuila/dlc-appointments-portal/pull/112), [#113](https://github.com/code-corhuila/dlc-appointments-portal/pull/113) |
| HU-APT-006 | Appointment types and calendar legend | done | [dlc-docs #178](https://github.com/code-corhuila/dlc-docs/issues/178), [appointments-api #89](https://github.com/code-corhuila/dlc-appointments-api/pull/89), [appointments-db #34](https://github.com/code-corhuila/dlc-appointments-db/pull/34), [appointments-portal #112](https://github.com/code-corhuila/dlc-appointments-portal/pull/112) |
| HU-014-004 | Billing encounters waiting for a price (API, portal, gateway route, compositor catalogue) | done | [billing-api #54](https://github.com/code-corhuila/dlc-billing-api/pull/54), [billing-portal #117](https://github.com/code-corhuila/dlc-billing-portal/pull/117), [api-gateway #20](https://github.com/code-corhuila/dlc-api-gateway/pull/20), [dlc-front #130](https://github.com/code-corhuila/dlc-front/pull/130) |
| HU-014-005 | Release 3.0.0 documentation: status, technical debt, user manual and deployment guide | done | [dlc-docs #175](https://github.com/code-corhuila/dlc-docs/pull/175), [#176](https://github.com/code-corhuila/dlc-docs/pull/176), [#179](https://github.com/code-corhuila/dlc-docs/pull/179), [#180](https://github.com/code-corhuila/dlc-docs/pull/180) |
| HU-014-006 | Gates actually block merges: CI runs on every PR and branch protection on develop/qa/main | todo | blocked - see section 3 |

## 2. My individual contribution

### Session 1 - CI/CD pipelines and quality gates
- **One pipeline per service:** every component repository has its own `.github/workflows/ci.yml`, triggered on pull requests to `develop`, `qa` and `main`. Gates by repository:
  - **dlc-clinical-api:** `gofmt`, `go vet`, `go test -race`, Docker image build, and an integration job against MongoDB 8 migrated by `dlc-clinical-db` and RabbitMQ 4.
  - **dlc-worker:** promotion trace validation, `mvn -B verify` (tests and coverage baselines) and image build.
  - **dlc-front:** PR gates for branch name, size, commits, traceability and architecture.
  - **dlc-infra-mongo (new this week):** environment Compose validation, replica set and transactions, network allow-list, TLS required, backup/restore ([#17](https://github.com/code-corhuila/dlc-infra-mongo/pull/17)).
- **Continuous delivery flow:** short-lived feature branch -> PR to `develop` -> promotion PR to `qa` -> `release: 3.0.0` PR to `main` -> tag `v3.0.0`. Release 3.0.0 followed this path in clinical (api, db, portal), patients, billing, worker, front, gateway and the infra repos.
- **Security stays a gate:** service-to-service calls now use short-lived IAM-issued service tokens by client credentials, replacing per-service secrets (ADR-020, [dlc-docs #171](https://github.com/code-corhuila/dlc-docs/pull/171)).
- **DORA (estimated from PRs this week):**
  - deployment frequency: one coordinated release to `main` (3.0.0, 2026-11-14) plus daily promotions to `qa`;
  - lead time: feature PR on 2026-11-09 -> `main` on 2026-11-14, about 5 days;
  - change failure rate and time to restore are not measured yet.

### Session 2 - DoD as gates and branch protection
| DoD item | Gate (verifier) | State |
|---|---|---|
| Code style consistent | `gofmt` / lint in `ci.yml` | defined, not running (see section 3) |
| Types / static checks | `go vet`, Maven compile, front architecture gate | defined, not running |
| Behavior tested | unit + integration (MongoDB, RabbitMQ) + contract tests of owner calls | defined; passes locally |
| Enough coverage | per-module coverage baselines in dlc-worker | dlc-worker only; missing elsewhere |
| No secrets / known CVEs | gitleaks + govulncheck (`security.yml`, week 13 proposal) | not yet in the component repos |
| Runs for real | Docker image build in CI; infra-mongo replica/TLS/backup checks | defined, not running |
| Reviewed and traceable | PR template, conventional commits, promotion trace, instructor approval to `main` | enforced by process, not by the platform |

- **Branch protection:** not verifiable from my account; the protection API returns `404` and the repositories have no rulesets. It must be enabled by an organization admin on `develop`, `qa` and `main`, with required checks and one review.
- **Remaining MVP 3 CI/CD stories:**
  - CI unblocked so gates run;
  - required checks on all three branches;
  - `security.yml` in each Go repository;
  - coverage floor outside dlc-worker;
  - merge to `main` pushes a versioned image;
  - smoke test (compose up + `/health`) after deployment.

## 3. Blockers and risks
- **CI does not run.** Every workflow run on 2026-11-14, including all `release: 3.0.0` runs, ends with: "The job was not started because your account is locked due to a billing issue." The gates are defined but do not block anything. Today they are self-certified by local runs (`go test`, `mvn verify`) and review.
- **Branch protection not confirmed** on `develop`/`qa`/`main` (see section 2).
- Some Billing release PRs (dlc-billing-api #103, dlc-billing-db #34, dlc-billing-portal #175) show as closed rather than merged, although their repositories carry the `v3.0.0` tag. The release path for Billing must be confirmed.

## 4. Plan for next week
- Ask the organization admin to restore GitHub Actions and enable branch protection with required checks.
- Add `security.yml` (gitleaks, govulncheck) and a coverage floor to the clinical repositories.
- Add a post-deploy smoke test (`docker compose up` + `/health` on gateway, front and services).
- Testing at scale and the MVP 3 demo on release 3.0.0.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [x] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Corte 3 full activity - week 14: every PR, commit, issue and review of LuisBonilla2260, verified](../../C3-ACTIVITY.md#week-14)
- [Full activity of the week: 91 PRs (85 merged), 97 commits, 43 issues, 2 reviews across the 22 repositories](../ACTIVITY.md)
- [dlc-infra-mongo #17 - CI pull request workflow](https://github.com/code-corhuila/dlc-infra-mongo/pull/17)
- [dlc-docs #171 - ADR-020 IAM service-token issuer](https://github.com/code-corhuila/dlc-docs/pull/171)
- [dlc-docs #175 - release 3.0.0 status](https://github.com/code-corhuila/dlc-docs/pull/175)
- [dlc-docs #176 - technical debt from release 3.0.0](https://github.com/code-corhuila/dlc-docs/pull/176)
- [dlc-docs #179 - user manual and deployment guide](https://github.com/code-corhuila/dlc-docs/pull/179)
- [dlc-clinical-api v3.0.0](https://github.com/code-corhuila/dlc-clinical-api/releases/tag/v3.0.0)
- [dlc-front v3.0.0](https://github.com/code-corhuila/dlc-front/releases/tag/v3.0.0)
- [dlc-clinical-api Actions runs (billing lock)](https://github.com/code-corhuila/dlc-clinical-api/actions)

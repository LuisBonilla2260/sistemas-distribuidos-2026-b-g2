# Week 12 - Session 2: Deployment plan, IaC and rollback for MVP 3 (clinical service)

Deployment plan for the Clinical context of Di-Lucca. Builds on session 1 (`../01-session/README.md`).

## 1. Artifact promotion

```text
merge to develop -> CI: test + build once -> ghcr.io/code-corhuila/dlc-clinical-api:sha-<commit>
                                           -> deploy to develop (compose, .env of develop)
PR develop -> qa  -> same image, retag :qa-<date>  -> deploy to qa   (.env of qa)
release tag vX.Y.Z on main -> same image, retag :vX.Y.Z -> deploy to main (.env of main)
```

| Repo | Artifact | Registry name | Promoted by |
|---|---|---|---|
| `dlc-clinical-api` | one image, five entrypoints | `ghcr.io/code-corhuila/dlc-clinical-api` | image tag in `CLINICAL_API_TAG` |
| `dlc-clinical-portal` | nginx image with the built bundle | `ghcr.io/code-corhuila/dlc-clinical-portal` | `CLINICAL_PORTAL_TAG` |
| `dlc-clinical-db` | Liquibase image with the changelogs | `ghcr.io/code-corhuila/dlc-clinical-db-migrate` | `CLINICAL_DB_TAG` |

Rules: no `build:` in a deploy compose file; config per environment only in `.env` and `CLINICAL_SECRETS_DIR`; the
release notes record the image digest of each repo.

## 2. Infrastructure as Code

| Piece | IaC today | Gap |
|---|---|---|
| Clinical processes | `dlc-clinical-api/deploy/compose.yml` (versioned, reviewed) | uses `build:`; needs `image:` and a blue/green pair for `clinical-api` |
| Clinical schema | `dlc-clinical-db` changelogs and `deploy/compose.yml` (`tooling` profile) | image not published |
| Clinical MongoDB | **none**: `dlc-infra-mongo` contains only a README | snowflake risk: the replica set that transactions need exists only on someone's machine |
| Networks `platform`, `clinical-data` | declared `external` in compose | created by hand; add a `make networks` or a compose file in the infra repo |
| Gateway route to Clinical | outside the clinical repos | the blue/green switch needs it as a reviewed file |

Done means: `qa` is recreated from git with `docker compose up -d` in each repo, and `develop`, `qa` and `main` differ
only in `.env` and secrets.

## 3. Strategy and rollback

Strategy (from session 1): blue-green for `clinical-api` and `dlc-clinical-portal`; rolling for consumers; stop-start
for the relay and the reconciler; expand/contract for migrations.

Rollback of `clinical-api` (target: under 5 minutes):

| Step | Command / check | Expected time |
|---|---|---|
| 1 | detect: availability or latency SLI of procedure completions out of target (week 11 alerts) | 0-2 min |
| 2 | switch the gateway route back to `clinical-api-blue` | under 1 min |
| 3 | verify: `/clinical-api healthcheck` in blue (same check as compose, `GET /health`) and one test completion returns `202` | 1 min |
| 4 | stop green: `docker compose stop clinical-api-green`; record version, time and reason | 1 min |

Rollback of a migration: only migrations written expand/contract can be reverted by redeploying the previous API
image; a contracting change (drop a field, tighter validator) waits one release after the code stops using it.
Liquibase `rollback` in `dlc-clinical-db` is the last resort, executed from the same migrate image tag.

Rehearsal: **not done yet.** It counts only when it is timed in `qa` and written here:

| Date | Env | From -> to | Detect | Switch | Verify | Total | Notes |
|---|---|---|---|---|---|---|---|
| pending | qa | v2.1.0 -> v2.0.0 | | | | | |

## 4. MVP 3 deployment stories

| ID | Story | Acceptance criteria |
|---|---|---|
| DEP-1 | As the team, CI builds and pushes a versioned image of `dlc-clinical-api` on merge, so every environment runs the tested artifact. | a merge to `develop` publishes `ghcr.io/code-corhuila/dlc-clinical-api:sha-<commit>`; a `v*` tag publishes `:vX.Y.Z`; the digest appears in the job summary |
| DEP-2 | As an operator, I deploy Clinical from images, not builds. | `deploy/compose.yml` has no `build:`; `CLINICAL_API_TAG` selects the image; `docker compose pull && up -d` in `qa` runs the same digest as `develop` |
| DEP-3 | As an operator, I recreate the Clinical MongoDB replica set from `dlc-infra-mongo`. | `docker compose up -d` creates a MongoDB 8.0 replica set; a two-collection transaction commits and aborts against it; readiness fails without a primary |
| DEP-4 | As a Dentist, a new clinical release does not interrupt procedure completions. | blue-green pair for `clinical-api`; during a switch in `qa` a loop of completions every second gets 0 failed responses; the switch waits for 3 healthy checks |
| DEP-5 | As the team, we can revert a bad clinical release in under 5 minutes. | rollback rehearsed in `qa` with the table in section 3 filled and total under 5 minutes; verification uses the week 11 SLIs |

## 5. Self-check

| # | Answer |
|---|---|
| 1 | Deploying the same built image to each environment (config differs) |
| 2 | Reproducible, version-controlled, reviewable environments |
| 3 | It has been rehearsed or tested |
| 4 | Blue-green deployment |
| 5 | A manually configured, non-reproducible machine |
| 6 | qa matches prod (built from the same IaC) |

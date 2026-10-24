# Week 12 - Session 1: Cloud deployment, zero downtime and scaling for the clinical service

Applies the session to the Clinical context of Di-Lucca: `dlc-clinical-api` (Go, five processes from one image),
`dlc-clinical-db` (Liquibase migrations for MongoDB), `dlc-clinical-portal` (React behind nginx) and
`dlc-infra-mongo` (the Clinical MongoDB replica set). State reviewed on 2026-10-24 on `develop`.

## 1. Deployment model

| Model | Fit for Clinical | Decision |
|---|---|---|
| IaaS (VM) | full control, but we would install Docker, networks and MongoDB by hand | host only: one Linux VM (or production-like host) runs Docker |
| Containers on the host / orchestrator | `dlc-clinical-api` already ships a distroless image with five entrypoints and a compose file | **chosen** for MVP 3: Docker Compose on one host, same image in every environment |
| PaaS | pushes code and rebuilds it; cannot run the outbox relay, inbox consumer, intent reconciler and analytics consumer as one release | no |
| Serverless | the relay and consumers are long-running RabbitMQ processes; MongoDB transactions need stable connections | no |

Compose on one host is the most control we can operate today with a five-person team. Kubernetes stays a later
option: the same image and health check move to it unchanged.

## 2. Build once, deploy the artifact

Current state:

| Repo | What happens today | Problem |
|---|---|---|
| `dlc-clinical-api` | `deploy/compose.yml` uses `build:` for every process | each environment rebuilds the image, so `qa` and `main` may not run what was tested |
| `dlc-clinical-api` | `ci.yml` runs tests on `develop`, `qa`, `main`; nothing is pushed to a registry | there is no artifact to promote |
| `dlc-clinical-portal` | `deploy/compose.yml` uses `build:` (plus a `build:demo` variant) | same rebuild problem; tags `v2.0.0` and `v2.0.1` exist only as git tags |
| `dlc-clinical-db` | the migration image is built locally under the `tooling` profile | migrations run from whatever checkout the operator has |

Target (12-factor build -> release -> run):

1. **Build:** on a merge to `develop`, CI builds `ghcr.io/code-corhuila/dlc-clinical-api` once, tags it with the git
   sha and, on a release tag, with `vX.Y.Z`.
2. **Release:** `compose.yml` uses `image: ghcr.io/code-corhuila/dlc-clinical-api:${CLINICAL_API_TAG}` instead of
   `build:`; each environment adds only its `.env` and its secrets in `CLINICAL_SECRETS_DIR`.
3. **Run:** `qa` and `main` pull the tag that passed `develop`, never a new build. The record of what runs is the
   image digest, written in the release notes next to `10-week/c2-release/RELEASES.md`.

## 3. Zero-downtime strategy per process

| Process | State | Strategy | Why |
|---|---|---|---|
| `clinical-api` (HTTP, port 8080) | stateless; state in MongoDB | **blue-green** behind the gateway | instant switch and instant rollback; the health check `/clinical-api healthcheck` already exists |
| `clinical-outbox-relay` | reads `outbox_events` | stop old, start new (seconds) | messages wait in `outbox_events`; nothing is lost, only delayed |
| `clinical-inbox-consumer`, `clinical-analytics-consumer` | RabbitMQ queues + `inbox_events` | rolling, one at a time | unacked messages are redelivered; `inbox_events` drops duplicates by `eventId` |
| `clinical-intent-reconciler` | settles PENDING write intents | stop old, start new | a short gap only delays reconciliation (ADR-017) |
| `dlc-clinical-portal` | static files in nginx | blue-green | switch the gateway route; old bundle stays for rollback |
| `clinical-db-migrate` | changes MongoDB validators and indexes | expand / contract | a migration must accept both the old and the new API image, or rollback breaks |

Blue-green for `clinical-api`:

```text
gateway -> clinical-api-blue  (v2.0.0, live)
           clinical-api-green (v2.1.0, starting) -- healthcheck OK x3 --> switch route to green
           keep blue running 30 min; if error ratio or p95 rises (week 11 SLOs), switch back to blue
```

The switch criteria reuse the week 11 workshop: the availability and latency SLIs of
`POST /api/v1/procedures/{id}/completions` (`11-week/workshop/05-slo.md`) must stay inside target for 15 minutes.

## 4. Scale out and expect failure

- `clinical-api` can scale out: it keeps no session state, JWTs are verified per request against the auth JWKS, and
  every write is idempotent by `procedureId` and the idempotency key.
- Consumers can run more than one replica on the same queue; RabbitMQ splits the messages and `inbox_events` keeps the
  effect once.
- `clinical-outbox-relay` and `clinical-intent-reconciler` stay at **one replica** until we verify that two copies do not
  publish the same outbox event twice or race on the same intent. To check in the code before scaling.
- Signal to scale: p95 latency and RabbitMQ queue depth, not CPU. At about 6 completions per hour (week 11 CP6) one
  replica is enough; a second `clinical-api` replica is for availability, not load.
- Failure: `restart: unless-stopped` and the health check already restart a dead process. The single point of failure
  left is MongoDB: `dlc-infra-mongo` has no replica set provisioned yet (issue #1).

## 5. The Friday manual deploy, in our project

Today a deploy is: `git pull` on the host, `docker compose up -d --build`. That rebuilds a new image, recreates every
process at once (seconds of downtime for completions) and leaves no record of which commit runs. It is exactly the
scenario of the session. The fix is the pipeline in section 2 plus the blue-green switch in section 3.

## 6. Self-check

| # | Answer |
|---|---|
| 1 | More convenience, less control |
| 2 | Build the image in CI and run that same image in every environment |
| 3 | Switch traffic to the new version and roll back instantly to the old |
| 4 | Scale out (more instances) behind a load balancer with health checks |
| 5 | It is not repeatable or reproducible and causes downtime and mistakes |
| 6 | Rolling, blue-green, canary |

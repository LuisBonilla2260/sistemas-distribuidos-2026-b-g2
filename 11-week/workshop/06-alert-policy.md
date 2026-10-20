# CP7 - Alert policy

Rules: [06-alerts.yml](06-alerts.yml). SLO: 99.5 % of `POST /api/v1/procedures/{id}/completions` without
`outcome="error"` in 30 days (`05-slo.md`). Budget: 5.28 failed completions, or 52.8 minutes of open hours.

## How much budget each alert represents

Burn rate 1 spends the whole budget in exactly 30 days (720 h).

| Alert | Burn rate | Long / short window | Error ratio threshold | Budget spent before firing | In our numbers |
|---|---|---|---|---|---|
| `ClinicalProcedureCompletionFastBurn` | 14.4 | 1h / 5m | 14.4 x 0.5 % = 7.2 % | 14.4 x 1 / 720 = **2 %** | 0.1 completions, about 1.1 min of outage |
| `ClinicalProcedureCompletionSlowBurn` | 6 | 6h / 30m | 6 x 0.5 % = 3 % | 6 x 6 / 720 = **5 %** | 0.26 completions, about 2.6 min of outage |

At full speed the fast burn empties the budget in 720 / 14.4 = 50 hours, a bit more than two clinic weeks of open
hours. The slow burn empties it in 120 hours.

The short window (5m, 30m) stops the alert as soon as the problem is fixed, so nobody is paged for an outage that
already ended.

## The low-traffic problem

The clinic does about 6 completions per hour. One error in an hour is 1 / 6 = 16 % error ratio, which is above the
7.2 % fast-burn threshold. Without a guard, every single failure would page, and with 5 failures allowed per month
that is noise.

Both rules add `sum(increase(...{outcome="error"}[1h or 6h])) >= 2`: two failures in the window are already 38 % of
the monthly budget, which is worth waking someone. One isolated error is left to the log review in the 50 % budget
policy step.

## Routing

| Severity | Where it goes | Who |
|---|---|---|
| `page` | team chat channel with mention, during clinic hours (Mon to Fri, 8 h) | clinical service owner, then the team lead |
| `ticket` | GitHub issue in `dlc-clinical-api` with the `reliability` label | reviewed the same day |

There is no paging outside clinic hours: no Dentist completes procedures then, so there is no user impact to stop.

## Cause-based alerts we will NOT create

| Alert | Why not |
|---|---|
| CPU of the `clinical-service` container above 80 % | The service is a light Go binary with 6 requests per hour. High CPU without failed completions is not user pain, and if completions fail the burn-rate alert already fires. |
| RabbitMQ queue `billing.procedure-completed` has more than 0 messages | The `202` is answered after the MongoDB transaction, so a short queue never hurts the Dentist. Real delays already show up in SLO 3 (closure convergence) and `CLOSURE_FAILED`. |
| MongoDB connection pool usage above 70 % | A symptom of load we do not have. If the pool is exhausted, completions time out and SLO 1 and SLO 2 catch it. |

## Validation

Check the rules with `promtool` from the official Prometheus image (Docker Desktop must be running):

```bash
docker run --rm -v "$PWD":/w --entrypoint promtool prom/prometheus:v2.53.0 check rules /w/06-alerts.yml
```

Result on 2026-10-20 (Windows 11, Docker Desktop, `prom/prometheus:v2.53.0`):

```text
Checking /w/06-alerts.yml
  SUCCESS: 6 rules found
```

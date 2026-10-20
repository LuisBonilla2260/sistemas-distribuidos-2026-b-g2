# CP6 - SLIs, SLOs and error budget

Journey: `POST /api/v1/procedures/{id}/completions` in `clinical-service`, then `ProcedureCompleted` to Appointments and
Billing. Metrics come from `04-metrics.md`. Window for every SLO: 30 rolling days.

## Traffic estimate

| Input | Value | Source |
|---|---|---|
| Clinics | 1 | `dlc-docs` 01-context/overview.md: "the current delivery serves one clinic" |
| Dentists | 2 | team estimate for the clinic |
| Appointments per Dentist per day | 24 | team estimate (about 20 minutes per patient in an 8 hour day) |
| Working days per week | 5 | team estimate |
| Completed procedures per appointment | 1 | assumption: most visits close one procedure; extras travel inside the same completion |

- Appointments per day: 2 x 24 = **48**
- Working days in 30 days: 5 / 7 x 30 = about **22**
- Completions in 30 days: 48 x 22 = **1,056**
- Clinic open minutes in 30 days: 22 days x 8 h x 60 = **10,560 min**
- Peak rate: 48 / 8 h = 6 per hour, about one completion every 10 minutes

Minutes are counted over open hours only, because nobody completes procedures when the clinic is closed.

## SLO 1 - Availability

| Item | Value |
|---|---|
| SLI | completions that did not fail on our side / all valid completions |
| Good | `outcome="accepted"` or `outcome="rejected"` (a business rule or stale version answer is correct behavior) |
| Bad | `outcome="error"` (5xx or timeout) |
| Target | **99.5 %** in 30 days |

```promql
sum(increase(clinical_procedure_completions_total{outcome!="error"}[30d]))
/
sum(increase(clinical_procedure_completions_total[30d]))
```

Error budget: 0.5 % of 1,056 = **5 failed completions** in 30 days, or 0.5 % of 10,560 = **52.8 minutes** of full outage
during open hours.

## SLO 2 - Latency

| Item | Value |
|---|---|
| SLI | completions answered in 300 ms or less / all completions |
| Target | **95 %** in 30 days |

```promql
sum(increase(clinical_procedure_completion_duration_seconds_bucket{le="0.3"}[30d]))
/
sum(increase(clinical_procedure_completion_duration_seconds_count[30d]))
```

300 ms is already a bucket boundary in `metrics.go`, so the SLI is exact. Error budget: 5 % of 1,056 = **52 slow
completions**, or 5 % of 10,560 = **528 minutes** in which every completion is slow.

## SLO 3 - Closure convergence

The `202` only means Clinical saved the procedure. The Dentist and the Assistant need the `CareClosure` to reach
`COMPLETED` (Appointments progress and Billing snapshots done) before the visit is finalized.

| Item | Value |
|---|---|
| SLI | care closures that reach `COMPLETED` within 5 minutes of the `202` / all care closures started |
| Target | **99 %** in 30 days |
| New metric | `clinical_care_closure_convergence_seconds` histogram, observed when `BillableItemsCreated` closes the `CareClosure` |

```promql
sum(increase(clinical_care_closure_convergence_seconds_bucket{le="300"}[30d]))
/
sum(increase(clinical_procedure_completions_total{outcome="accepted"}[30d]))
```

A closure that ends in `CLOSURE_FAILED` (for example no catalog price for `OPER-01`) never enters the bucket, so it is
counted as bad. Error budget: 1 % of 1,056 = **10 closures**, or 1 % of 10,560 = **105.6 minutes**.

## Why these targets

- 99.5 % and not 99.9 %: with 1,056 completions a month, 99.9 % allows one failure. One restart of the service during
  a deploy would break the SLO. 99.5 % still means fewer than one failed completion per week.
- Low volume makes the budget fragile: one bad deploy that fails three completions burns 60 % of SLO 1. Alerts in CP7
  must not page for a single error.
- 95 % under 300 ms is reachable: the completion is one MongoDB transaction plus the outbox insert, with no
  synchronous call to another service.

## Error budget policy

| Budget consumed (any SLO) | Action |
|---|---|
| 50 % | Review the failed traces by `trace_id` in the next team meeting. New features continue. |
| 75 % | Only fixes and low-risk changes go to `qa`. The clinical owner reviews every merge to `develop` that touches the completion path. |
| 100 % | Feature freeze for `dlc-clinical-api` until the budget recovers. The next sprint starts with the reliability stories from the postmortem (CP9). |

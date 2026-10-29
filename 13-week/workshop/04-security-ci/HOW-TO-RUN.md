# How to run the security gates and the negative test

## What it covers

| Item | Value |
|---|---|
| Service | `clinical-service` (repository `code-corhuila/dlc-clinical-api`, Go) |
| Boundary | api-gateway -> clinical-service (inbound) |
| Story | [dlc-docs#137](https://github.com/code-corhuila/dlc-docs/issues/137) - Only the assigned dentist can read or complete a clinical record (STRIDE E, risk 9, P0) |
| Acceptance criteria covered | 1 (unassigned DENTIST cannot read `GET /api/v1/clinical-records/{id}`) and 3 (SECRETARY_ASSISTANT gets `403`; here on `POST /api/v1/clinical-records`) |
| Test file | [authz_negative_test.go](authz_negative_test.go) |
| Workflow | [security.yml](security.yml) |

## Run the negative test locally

The test uses the helpers of the package `httpapi` (`recordsRouter`, `tokenFor`, `post`, `do`), so it runs inside the clinical-service code with an in-memory store and RS256 test keys. No MongoDB or Gateway is needed.

1. Clone `code-corhuila/dlc-clinical-api` (branch `develop`) and install Go 1.26.
2. Copy `authz_negative_test.go` to `internal/adapter/in/httpapi/`.
3. Run:

```
go test ./internal/adapter/in/httpapi/ -run TestClinicalRecordAuthz -v
```

Expected result: both tests `PASS`. If clinical-service ever lets a non-assigned dentist read the record (`200`), the first test fails and the pull request is blocked.

## Workflow gates

| Job | Tool | Blocks the PR when |
|---|---|---|
| `secret-scan` | gitleaks | a secret or key is committed |
| `dependency-audit` | govulncheck | a dependency of `go.mod` has a known vulnerability that the code reaches |
| `authz-tests` | `go test` | a negative authorization test fails |

The workflow is a proposal for the clinical-service pipeline; it lives in this fork folder and is not yet added to `dlc-clinical-api`.

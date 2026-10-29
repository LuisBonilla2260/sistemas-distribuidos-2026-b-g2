# How to run the security gates and the negative test

## What it covers

| Item | Value |
|---|---|
| Service | `clinical-service` (repository `code-corhuila/dlc-clinical-api`, Go) |
| Boundary | api-gateway -> clinical-service (inbound) |
| Story | [dlc-docs#137](https://github.com/code-corhuila/dlc-docs/issues/137) - Only the assigned dentist can read or complete a clinical record (STRIDE E, risk 9, P0) |
| Acceptance criteria covered | 1 (unassigned DENTIST cannot read `GET /api/v1/clinical-records/{id}`), 4 (SECRETARY_ASSISTANT gets `403` on `POST /api/v1/clinical-records`) and 5 (Appointments down: access denied, fail closed) |
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

## Recorded run (CP8)

Run on 2026-10-29 against `code-corhuila/dlc-clinical-api`, branch `develop`, commit `12438c4`, Go 1.26.7 on Windows. The test file was copied into `internal/adapter/in/httpapi/` of a local copy; the team repository was not changed.

```
> go test ./internal/adapter/in/httpapi/ -run TestClinicalRecordAuthz -v -count=1
=== RUN   TestClinicalRecordAuthzUnassignedDentistCannotRead
--- PASS: TestClinicalRecordAuthzUnassignedDentistCannotRead (0.04s)
=== RUN   TestClinicalRecordAuthzAssistantCannotOpenRecord
--- PASS: TestClinicalRecordAuthzAssistantCannotOpenRecord (0.00s)
=== RUN   TestClinicalRecordAuthzFailsClosedWhenAppointmentsIsDown
    authz_negative_test.go:90: Appointments down -> 503
--- PASS: TestClinicalRecordAuthzFailsClosedWhenAppointmentsIsDown (0.00s)
PASS
ok  	github.com/code-corhuila/dlc-clinical-api/internal/adapter/in/httpapi	8.444s
```

What it proves: the unassigned dentist gets `404` (the record is not revealed), the assistant gets `403`, and when the Appointments assignment check fails the assigned dentist gets `503` instead of the record.

## Workflow validation (CP8)

```
> go run github.com/rhysd/actionlint/cmd/actionlint@v1.7.7 -verbose security.yml
verbose: Found 0 parse errors in 0 ms for security.yml
verbose: Rule "shellcheck" was disabled: exec: "shellcheck": executable file not found in %PATH%
verbose: Rule "pyflakes" was disabled: exec: "pyflakes": executable file not found in %PATH%
verbose: Found total 0 errors in 50 ms for security.yml
```

Result: actionlint v1.7.7 reports 0 errors (exit code 0). The shellcheck rule for `run:` scripts was not checked because shellcheck is not installed on this machine.

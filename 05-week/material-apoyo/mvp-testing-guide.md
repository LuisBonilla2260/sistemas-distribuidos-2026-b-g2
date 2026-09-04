# Di-Lucca MVP testing guide

> Week 05 support material. This guide records the testing work performed on the Di-Lucca MVP, a modular monolith composed of Angular 20, Spring Boot, and PostgreSQL.

## Purpose

Testing gives the team rapid feedback while the MVP evolves. The goal is not only to obtain a green build: each test should verify an observable business rule, HTTP contract, security behavior, or application configuration.

The backend suite and its QA evidence are available in the [`code-corhuila/di-lucca` QA branch](https://github.com/code-corhuila/di-lucca/tree/qa). The source tests were added in commit [`5fc9fff`](https://github.com/code-corhuila/di-lucca/commit/5fc9fff322e52f96c75a686c9bb2f862186ff415); the detailed [QA report](https://github.com/code-corhuila/di-lucca/blob/qa/HU-03/qa.md) describes the execution results and remaining risks.

## What was tested

### Backend: application services

The application-layer tests use JUnit and Mockito to isolate business services from repositories, email delivery, token providers, and other external adapters. They cover both successful and error paths for:

- Appointments: creation, cancellation, rescheduling, availability, and listing.
- Authentication: registration, login, refresh tokens, logout, and password recovery.
- Patients, procedures, schedules, clinical records, invoices, and payments.

These tests check decisions and collaborations, such as rejecting duplicates, reserving or releasing a slot, raising domain exceptions, or avoiding an email when it is not applicable. A mocked dependency is intentional at this level: it proves the service behavior, but not a real database write or email delivery.

### Backend: HTTP and security

Controller tests use `@WebMvcTest`, MockMvc, and Spring Security Test to verify routes, request/response mapping, status codes, and delegation to use cases. The suite also tests JWT generation and validation, refresh-token hashing, and the authentication filter behavior for valid, invalid, missing, and disabled-user cases.

The application-context test uses `@SpringBootTest` with the `test` profile and H2. It checks that Spring can build the context, configure JPA repositories, and initialize reference data without startup errors.

### Frontend baseline

The Angular MVP includes an `app.spec.ts` baseline using TestBed. It verifies that the application can be created and provides a starting point for component and service tests. The frontend itself covers authentication, dashboard, patients, procedures, scheduling, appointments, clinical records, invoices, calendar, and user management; its route guards and JWT interceptor are key candidates for the next test cases.

## Verified backend result

| Measure | Result |
|---|---:|
| Test classes | 28 |
| Executed / passed | 112 / 112 |
| Failures / errors / skipped | 0 / 0 / 0 |
| Pass rate | 100% |
| Branch coverage | 39.44% |

The result shows that the covered application, controller, security, and context behavior passed. It is not a guarantee of full production behavior: H2 differs from PostgreSQL, dependencies such as email are mocked, and browser-level flows are not covered yet.

## How to run the tests

### Backend

From the MVP repository root:

```powershell
Set-Location backend/odontosys-api
.\mvnw.cmd test
```

The test profile uses an in-memory H2 database. After execution, Maven writes Surefire results under `target/surefire-reports`; the JaCoCo HTML report is generated under `target/site/jacoco/index.html` when coverage generation is configured.

### Frontend

From the MVP repository root:

```powershell
Set-Location frontend/dilucca
npm ci
npm test
```

Run the frontend suite in the same environment used by CI so browser and dependency versions are reproducible.

## Reading a test result

A passed unit test means that, for its prepared input and test doubles, the observed behavior matched its assertions. It does not automatically prove that every input, real PostgreSQL persistence, external email delivery, concurrent requests, or the complete browser-to-database flow works in production.

## Next testing priorities

1. Automate the existing tests and monitor their results over time.
2. Extend coverage across the system integrations and user-facing behavior.
3. Keep test evidence and quality documentation up to date with each MVP change.

## Test evidence checklist

- Keep test source files versioned with the feature they verify.
- Use conventional commits such as `test(backend): add unit tests for core MVP services`.
- Attach the CI result and coverage report to the pull request or QA evidence.
- Record known limitations and failed scenarios; a transparent gap is more useful than an unsupported quality claim.

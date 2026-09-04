<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       05-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 05

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Strengthen the MVP testing baseline and support backend development.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-005-001 | Support backend and frontend development in the Di-Lucca MVP monolith | done | [MVP frontend and core workflows](https://github.com/code-corhuila/di-lucca/commit/80ba844992d8d94732f788e6e3de6f617677ef8c) |
| HU-005-002 | Verify the MVP testing baseline and define its improvement plan | done | [Backend unit-test suite](https://github.com/code-corhuila/di-lucca/commit/5fc9fff322e52f96c75a686c9bb2f862186ff415) |

## 2. My individual contribution
- Reviewed the Di-Lucca MVP as a monolith composed of an Angular 20 frontend, a Spring Boot backend, and PostgreSQL.
- Supported backend development by preserving the separation between application, domain, and infrastructure packages, and by focusing on validation, error handling, API consistency, and testability.
- Reviewed the frontend development: Angular routes cover authentication, dashboard, patients, procedures, scheduling, appointments, medical records, invoices, calendar, and user management; route guards and the JWT interceptor protect communication with the backend API.
- Added and verified 28 backend test classes covering application services, JWT security, and HTTP controllers. The QA report records 112 of 112 tests passed using JUnit, Mockito, MockMvc, Spring Security Test, H2, and JaCoCo.
- Verified the backend application-context test and the Angular `app.spec.ts` baseline, which checks application creation and basic rendering.
- Reviewed the containerized MVP topology: Angular/Nginx, Spring Boot, and PostgreSQL run together through Compose, with a backend health check, internal network, and persistent PostgreSQL volume.

## 3. Blockers and risks
- The MVP is still a monolith; future service extraction requires validated ownership boundaries and integration contracts.
- Backend unit, controller, security, and application-context coverage is available, but PostgreSQL integration, API-contract, and end-to-end flows still require automated evidence.
- The current Angular title assertion is a scaffold-level check and should be replaced by tests for the real user-facing views and critical workflows.

## 4. Plan for next week
- Automate the backend test suite in CI and monitor coverage.
- Extend testing with database integration, API-contract, frontend, and end-to-end checks.
- Keep the MVP documentation and architecture aligned with the tested behavior.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [x] Tests added/updated (unit / integration)
- [ ] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Di-Lucca MVP main branch](https://github.com/code-corhuila/di-lucca/tree/main)
- [MVP frontend and core workflows commit](https://github.com/code-corhuila/di-lucca/commit/80ba844992d8d94732f788e6e3de6f617677ef8c)
- [MVP backend and core workflows commit](https://github.com/code-corhuila/di-lucca/commit/541b313)
- [Backend unit-test suite commit](https://github.com/code-corhuila/di-lucca/commit/5fc9fff322e52f96c75a686c9bb2f862186ff415)
- [Backend quality test report and QA evidence](https://github.com/code-corhuila/di-lucca/blob/qa/HU-03/qa.md)
- [MVP backend test suite](https://github.com/code-corhuila/di-lucca/tree/qa/backend/odontosys-api/src/test)
- [MVP frontend test baseline](https://github.com/code-corhuila/di-lucca/blob/main/frontend/dilucca/src/app/app.spec.ts)

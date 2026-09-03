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
| HU-005-001 | Support backend development in the Di-Lucca MVP monolith | doing | [di-lucca-mvp main branch](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main) |
| HU-005-002 | Define the MVP testing baseline and improvement plan | doing | [di-lucca-mvp main branch](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main) |

## 2. My individual contribution
- Reviewed the Di-Lucca MVP as a monolithic application composed of an Angular 20 frontend, a Spring Boot backend, and PostgreSQL.
- Reviewed the backend separation into application, domain, and infrastructure packages to support ongoing backend development without mixing responsibilities.
- Assessed the current testing baseline: backend application-context verification and frontend application specifications are present, while unit, integration, and API-contract coverage require expansion.
- Supported backend development by emphasizing validation, error handling, API consistency, and testability as the MVP evolves.

## 3. Blockers and risks
- The MVP is still a monolith; service extraction requires validated ownership boundaries and integration contracts.
- Current automated test coverage is limited and must be expanded before changes to backend behavior are considered reliable.

## 4. Plan for next week
- Add backend unit tests for application and domain behavior.
- Add integration tests for persistence and API endpoints.
- Define API-contract and error-response checks between Angular and Spring Boot.
- Continue supporting backend improvements while preserving the monolith's current behavior.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [ ] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [ ] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Di-Lucca MVP main branch](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main)
- [MVP backend test baseline](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main/backend/odontosys-api/src/test)
- [MVP frontend test baseline](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main/frontend/dilucca/src/app/app.spec.ts)

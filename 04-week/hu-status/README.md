<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       04-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 04

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Design validation and microservices boundary definition in di-lucca MVP monolith
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-004-001 | Validate Figma design alignment with current monolith implementation | doing | [OdontoSys Figma design](https://www.figma.com/design/HB2cfqrV1HFzARClICuitz/OdontoSys?node-id=0-1&t=B9PUkD44i8l2u72o-1) |
| HU-004-002 | Document current domain boundaries in monolithic architecture | doing | [di-lucca-mvp main branch](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main) |
| HU-004-003 | Define microservices boundaries for appointment management system | todo |  |
| HU-004-004 | Map service dependencies and cross-domain interactions | todo |  |

## 2. My individual contribution
- Analyzed the di-lucca MVP monolithic system (Java 58.3% / TypeScript 16.3% / HTML 14% / CSS 11.4%)
- Conducted design-to-implementation comparison between initial Figma specifications and actual developed codebase
- Identified domain boundaries and logical service separation points within the monolith
- Documented architectural misalignments and service coupling issues for microservices decomposition planning
- Reviewed frontend (Angular) and backend architecture to establish clear service boundaries

## 3. Blockers and risks
- **HIGH - Design validation pending**: Figma screens and the current implementation still require a documented, screen-by-screen comparison.
- **HIGH - Boundary validation pending**: The monolith contains appointment, patient, scheduling, billing, and clinical responsibilities that require a validated decomposition proposal.
- **MEDIUM - Migration risk**: The current Angular, Spring Boot, and PostgreSQL application must be decomposed incrementally to preserve existing behavior.
- **MEDIUM - Cross-domain dependencies**: Service interactions and data ownership must be mapped before defining asynchronous integration contracts.

## 4. Plan for next week
- [ ] Complete architecture audit report documenting all design-to-implementation discrepancies
- [ ] Create detailed domain event matrix showing cross-service communication points
- [ ] Propose refactored domain model aligned with original Figma specifications
- [ ] Design event-driven architecture pattern to decouple current monolithic domains
- [ ] Define clear microservices boundaries (Appointment Service, Patient Service, Billing Service, Notification Service)
- [ ] Establish API contracts and message schemas for inter-service communication
- [ ] Create migration roadmap from monolith to distributed system with minimal disruption

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-001-dev -> develop, ...)
- [ ] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration) - deferred to HU-001-003
- [ ] DDD / hexagonal boundaries respected (domain has no I/O) - assessment in progress
- [x] No secrets; config via environment variables - verified in repository

## 6. Evidence links
- [OdontoSys Figma design](https://www.figma.com/design/HB2cfqrV1HFzARClICuitz/OdontoSys?node-id=0-1&t=B9PUkD44i8l2u72o-1)
- [di-lucca-mvp main branch](https://github.com/DanielPerez1822/di-lucca-mvp/tree/main)

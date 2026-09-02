<!-- HU-STATUS TEMPLATE - do NOT remove the <!-- ... --> markers or the table headers.
     Your weekly grade is read AUTOMATICALLY from this file:
       03-week/hu-status/README.md  (inside YOUR fork). English. -->

# Weekly Status - Week 03

<!-- CONFIG-START - must match your profile repo (username/username) CONFIG -->
- FULL_NAME: Luis Ignacio Bonilla Delgado
- GITHUB_USER: LuisBonilla2260
- TEAM: Di-Lucca
- SPRINT_GOAL: Define the Analytics bounded context, its event-driven integration, and its independent read model.
<!-- CONFIG-END -->

## 1. User stories worked this week
| HU ID | Title | Status (todo/doing/done) | Evidence (PR or commit URL) |
|---|---|---|---|
| HU-003-001 | Define Clinical and Analytics architecture and data models | done | [Analytics context and domain commit](https://github.com/code-corhuila/dlc-docs/commit/fd17915082db52668910f742bf69436b3052d970) |
| HU-003-002 | Define Analytics domain events and projection model | done | [Analytics context and domain commit](https://github.com/code-corhuila/dlc-docs/commit/fd17915082db52668910f742bf69436b3052d970) |

## 2. My individual contribution
- Defined and documented the responsibilities and boundaries of the Clinical and Analytics microservices.
- Defined Analytics as a downstream service that consumes domain events instead of accessing transactional databases directly.
- Defined the use of independent MongoDB databases for Clinical and Analytics, respecting microservice data ownership.
- Defined RabbitMQ as the asynchronous communication mechanism for domain events.
- Defined the `AnalyticsProjection` model and the main events consumed by Analytics.
- Established rules for event processing, projection reconstruction, data independence, and prevention of duplicate event processing.
- Prepared the technical explanation of the architecture and data flow for the Weekly presentation.

## 3. Blockers and risks
- Final validation of the domain event contracts between services is still pending.
- Documentation for domains, data models, events, and governance rules must remain synchronized.
- Analytics must not access the databases owned by transactional services directly.

## 4. Plan for next week
- Finalize and validate the Clinical and Analytics data models.
- Finalize the domain event contracts consumed by Analytics.
- Review and align the architecture, domain, and governance documentation.
- Validate the consistency of the Analytics projection model with the defined events.
- Continue preparing the technical documentation required for the next development stage.

## 5. Compliance self-check
- [x] Conventional Commits - `type(scope): summary`
- [ ] Per-environment HU branch + PR to that environment (hu-xxx-dev -> develop, ...)
- [x] Testable acceptance criteria
- [ ] Tests added/updated (unit / integration)
- [x] DDD / hexagonal boundaries respected (domain has no I/O)
- [x] No secrets; config via environment variables

## 6. Evidence links
- [Analytics context and domain commit](https://github.com/code-corhuila/dlc-docs/commit/fd17915082db52668910f742bf69436b3052d970)
- [Analytics context and domain changes](https://github.com/code-corhuila/dlc-docs/tree/fd17915082db52668910f742bf69436b3052d970)

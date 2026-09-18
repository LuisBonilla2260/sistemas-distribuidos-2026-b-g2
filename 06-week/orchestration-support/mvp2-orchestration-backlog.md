# MVP 2 orchestration support backlog

These stories are proposed support work. They are not claims that the Di-Lucca application already implements the behavior.

## HU-006-001 — Deterministic local startup

**As a developer, I want all local MVP services to start through one Compose command so that integration work begins from a reproducible topology.**

Acceptance criteria:

- A documented `docker compose up --build` command starts the required MVP services.
- Services communicate through Compose service names on a shared network.
- Database state is mounted on a named volume.
- The database exposes a health check, and dependent startup waits for the healthy state.
- The dependent service reports a clear error or retries with bounded backoff if the dependency is unavailable.

## HU-006-002 — Environment-owned configuration

**As a team member, I want a shared configuration contract so that the same artifact can run in develop, QA, and production without code changes.**

Acceptance criteria:

- Required variable names are documented in `.env.example` without real credentials.
- Application startup validates mandatory configuration.
- Local secrets are ignored by Git.
- Develop and QA run the same image version with environment-specific injected values.

## Definition of evidence

Completion evidence should include the relevant HU branch or pull request, the Compose configuration, startup output with secrets redacted, and a short verification record against the acceptance criteria. The documentation in this folder supplies the baseline; implementation evidence belongs to the application repository when the team performs the work.

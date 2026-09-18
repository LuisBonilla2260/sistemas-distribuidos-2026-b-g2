# Configuration and environment strategy

## Principle

Build one immutable application image and promote that same image through `develop`, `qa`, and `prod`. Environment values change; application code and image contents do not. This follows the 12-factor configuration principle and prevents configuration drift.

## Environment matrix

| Variable | develop | qa | prod |
|---|---|---|---|
| `DB_URL` | Compose service URL for the development database | QA database endpoint | Production database endpoint |
| `LOG_LEVEL` | `debug` | `info` | `warn` or approved operational level |
| `JWT_SECRET` | Locally injected development secret | QA secret injection | Production secret store injection |
| `APP_ENV` | `develop` | `qa` | `prod` |

Values shown here are categories, not deployable credentials. No secret value belongs in this repository.

## Shared configuration contract

All services should publish one documented set of variable names. A service validates required values during startup and fails fast with an actionable message when a non-secret value is missing or malformed. It must never print secret values in logs.

Commit a `.env.example` near the relevant Compose definition with placeholders only:

```dotenv
APP_ENV=develop
DB_URL=jdbc:postgresql://db:5432/app
DB_USERNAME=change-me
DB_PASSWORD=change-me
JWT_SECRET=change-me
LOG_LEVEL=debug
```

The real `.env` file must be ignored by Git. QA and production receive their values from the CI/CD environment or an approved secret store.

## Promotion and branches

The expected progression is `hu-xxx-dev` to `develop`, then `hu-xxx-qa` to `qa`, then `hu-xxx-main` to `main` for production. A change should pass QA before production promotion. A branch is not a configuration mechanism: the matching environment injects its own values while running the tested artifact.

## Drift-prevention review

- [ ] Variable names match across Compose, service startup validation, CI/CD, and `.env.example`.
- [ ] URLs do not contain hard-coded environment hosts such as `localhost` when running inside Compose.
- [ ] Secrets are ignored locally and injected outside Git.
- [ ] The image digest promoted to QA is the same artifact validated in develop.
- [ ] Configuration changes receive the same review discipline as code changes.

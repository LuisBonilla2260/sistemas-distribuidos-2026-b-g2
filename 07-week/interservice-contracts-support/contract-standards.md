# Versioned contract standards

## Source of truth

Every externally consumed API or event must have a versioned, machine-readable contract stored with the service or in the approved documentation repository:

| Communication style | Contract artifact |
|---|---|
| REST | `openapi.yaml` |
| gRPC | `.proto` |
| Event | JSON Schema or Avro schema |

Implementation and tests must follow the contract. A behavior not described in the contract is not a supported integration promise.

## Minimum contract content

A REST operation declares method, path, request shape, response shape, error responses, and version. An event declares its name, version, payload schema, producer, expected delivery semantics, and idempotency identifier.

Use these shared conventions unless an approved project convention replaces them:

- UUIDs for public identifiers.
- ISO-8601 timestamps in UTC.
- `page` and `limit` for pagination when applicable.
- One error envelope:

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "A human-readable explanation",
    "details": [],
    "trace_id": "uuid"
  }
}
```

## Compatibility rules

Safe within the same major version:

- Add an optional response or event field.
- Add a new optional request field with a documented default behavior.

Breaking changes that require a new version:

- Remove or rename a field.
- Change a field type or its meaning.
- Make an optional field mandatory.
- Change endpoint behavior in a way that invalidates existing consumers.

A breaking REST change uses a new API version, such as `/api/v2`. A breaking event change uses a new event version or name. Keep the previous version available during the agreed migration period, announce deprecation, and use a `Sunset` header for HTTP endpoints when applicable.

## Review checklist

- [ ] Contract change is reviewed before or together with implementation.
- [ ] Compatibility impact is classified as safe or breaking.
- [ ] Consumers and migration window are identified for a breaking change.
- [ ] Examples contain no secrets or production data.
- [ ] The contract can be validated by automated tooling.

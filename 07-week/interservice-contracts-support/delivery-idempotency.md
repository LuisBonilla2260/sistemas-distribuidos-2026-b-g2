# Delivery semantics and idempotent consumers

## Delivery is not processing

Networks and brokers can lose or repeat a delivery. End-to-end exactly-once delivery should not be assumed. The practical baseline is at-least-once delivery plus idempotent processing: a consumer can receive the same event more than once without repeating its business effect.

| Semantic | Consequence | Design response |
|---|---|---|
| At-most-once | A message can be lost | Use only when loss is acceptable or independently recoverable |
| At-least-once | A message can be duplicated | Require idempotent consumers and deduplication |
| Exactly-once processing | Desired business outcome | Engineer from idempotency, atomic state changes, and observability |

## Consumer pattern

Each event carries a stable identifier, for example `event_id`. Before applying the business effect, the consumer checks whether that identifier was already completed. The record of the effect and the processed-event marker should be committed atomically where the data store supports it.

```text
handle(event):
  if already_processed(event.event_id):
      acknowledge without repeating the effect

  validate event against its versioned schema
  apply the business effect
  record event.event_id as processed in the same transaction
  acknowledge
```

If the effect fails, return or record a failure according to the retry policy; do not mark the event complete. If a retry is exhausted, route it to an observable recovery path such as a dead-letter queue. Never silently discard malformed or permanently failing messages.

## Acceptance checks for a future implementation

- [ ] The event schema includes a stable `event_id` and version.
- [ ] Replaying the same event does not create a second business effect.
- [ ] A failed event can be retried without corrupting state.
- [ ] Failed retries are visible to operators.
- [ ] Logs and traces correlate processing with `event_id` without exposing secrets.

# Inter-service communication decisions

## Purpose

This support guide helps the team choose a communication style per interaction. It is not a statement that the current MVP already has independently deployed services. Confirm the real service boundaries and requirements before implementation.

## Decision matrix

| Interaction need | Recommended option | Why | Required safeguards |
|---|---|---|---|
| A browser or external client needs an immediate result | REST/JSON | Widely supported, inspectable, and suitable for public APIs | OpenAPI contract, input validation, timeout, standard errors |
| An internal caller needs a low-latency typed response | gRPC | Compact payloads and contract-first generated clients | `.proto` versioning, deadlines, retries only when safe |
| Several contexts react independently to one completed fact | Event topic (pub/sub) | Decouples the producer and enables fan-out | Schema versioning, idempotent consumers, observability |
| One background task must be processed by one worker | Queue | Distributes work and buffers spikes | Retry policy, dead-letter handling, idempotency |

## Synchronous calls

REST and gRPC couple a caller to the availability and latency of the downstream service. Keep synchronous chains short. Every remaining call needs an explicit timeout and a failure behavior that the caller can handle. Do not retry non-idempotent operations unless the contract explicitly defines a safe idempotency key.

## Asynchronous events

Events are appropriate when the producer can continue without an immediate answer and multiple consumers may react. An event represents a fact that already happened, for example `OrderPlaced`; it is not a disguised remote procedure call. Consumers must tolerate delayed, duplicated, and out-of-order delivery where the broker or business process permits it.

## Decision record template

For every proposed interaction, record:

1. Caller and provider/consumers.
2. Business need for an immediate answer, if any.
3. Chosen protocol or event channel and its reason.
4. Timeout, retry, and failure behavior for synchronous calls.
5. Delivery semantics and idempotency key for asynchronous consumption.
6. Link to the machine-readable contract.

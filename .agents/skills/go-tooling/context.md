# Context

Use this when code touches `context.Context`, HTTP request lifecycle, database calls, goroutines, cancellation, timeouts, or background work.

## Core rules

- Accept `context.Context` as the first parameter when an operation may block, perform I/O, wait on another service, or be canceled.
- Propagate the caller's context downward. Do not replace it with `context.Background()` in the middle of a request path.
- Use `context.Background()` at process roots: `main`, tests, startup jobs, and explicit background workers.
- Use `context.TODO()` only as a temporary marker. Do not leave it in finished code without a comment explaining why.
- Do not pass `nil` context.
- Do not store context in structs except for rare long-lived framework types that document that pattern.

## Deadlines and cancellation

Prefer a timeout at the boundary that owns the latency budget:

- incoming HTTP server: request context is already canceled when the client disconnects
- outgoing HTTP/database call: add a local timeout if the caller did not already provide one
- background worker: derive from a process/shutdown context

Always call the cancel function returned by `context.WithCancel`, `context.WithTimeout`, or `context.WithDeadline`:

```go
ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
defer cancel()
```

## Values

Use context values only for request-scoped metadata that crosses API boundaries, such as request ID, trace ID, authenticated principal, or logger enrichment.

Do not use context values for ordinary optional parameters, configuration, dependencies, or hidden control flow.

Use unexported key types to avoid collisions:

```go
type requestIDKey struct{}
ctx = context.WithValue(ctx, requestIDKey{}, requestID)
```

## Background work after a request

If work must continue after the request returns, make that decision explicit:

- copy the small data needed by the worker
- use a process/shutdown context, not the request context
- if using `context.WithoutCancel`, add a new timeout and do not rely on request-scoped values accidentally

## Verification

For context-sensitive changes, verify the relevant cancellation path:

```sh
go test ./... -run 'Context|Cancel|Timeout' -count=1
go test -race ./...
```

Use the race gate when contexts coordinate goroutines or shared state.

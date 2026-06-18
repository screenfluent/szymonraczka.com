# Safety

Use this when reviewing Go correctness risks: nil behavior, slices, maps, numeric conversions, resource lifetime, goroutines, zero values, or defensive copying.

## Nil traps

Know the different nil behaviors:

- read from nil map: OK, returns zero value
- write to nil map: panic
- append to nil slice: OK
- send/receive on nil channel: blocks forever
- call method through nil interface: depends on the dynamic value and can panic

Watch for the nil interface trap:

```go
var p *MyError = nil
var err error = p
err != nil // true
```

Return a plain `nil` error, not a typed nil behind an interface.

## Slices and maps

Appending to a slice can mutate the same backing array. If the callee must not affect the caller's data, copy first:

```go
items = append([]Item(nil), items...)
```

Maps are reference types. Copy maps when storing caller-provided data or returning internal state.

Do not read and write a map from multiple goroutines without synchronization. Use a mutex, channel ownership, or `sync.Map` only when its access pattern fits.

## Resources

Do not `defer` inside long loops when the resource must close each iteration. Use a helper function or close explicitly after the iteration's work.

Always check close errors when closing a writer or file where buffered data may fail to flush.

## Goroutines

Every goroutine needs an exit path. Name what stops it:

- context cancellation
- channel close
- bounded worker lifetime
- process shutdown

Tests for goroutine code should cover cancellation or completion. Use `go test -race ./...` for shared state.

## Numbers

Be explicit with narrowing conversions. Check ranges before converting external input or database values.

Do not compare floats for exact equality unless the values are discrete by construction. Use a tolerance for measured values.

Guard division by zero when the divisor can come from input or data.

## Initialization

Prefer useful zero values when practical.

Avoid `init()` for behavior that can be expressed as explicit construction. Hidden initialization makes tests and ordering harder to reason about.

Use `sync.Once` only for real lazy initialization. Do not use it to hide unclear lifecycle ownership.

## Verification

Use focused checks based on the risk:

```sh
go test ./... -run '<relevant-test>' -count=1
go test -race ./...
go test ./... -shuffle=on -count=10
```

Race detector proves only executed paths showed no data race. It does not prove protocol correctness.

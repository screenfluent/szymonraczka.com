# Performance

Use this for cache, rendering, JSON/XML, SQL hot paths, allocations, auth/session lookup, static assets, middleware, or background work.

## Rule

Do not optimize without measurement.

Name the hot path first.
Use the simple version first.
Add cache only when there is a real reason and invalidation is clear.

## Benchmarks

Write benchmarks only for a named hot path or regression risk. For new benchmarks on modern Go, prefer `b.Loop()` when it fits the benchmark shape.

```sh
go test ./... -run '^$' -bench . -benchmem
```

Compare before/after with repeated runs and `benchstat`:

```sh
go test ./... -run '^$' -bench BenchmarkName -benchmem -count=10 > /tmp/before.txt
go test ./... -run '^$' -bench BenchmarkName -benchmem -count=10 > /tmp/after.txt
benchstat /tmp/before.txt /tmp/after.txt
```

Report benchmark numbers in the final note. Do not claim a performance win from a single noisy run.

## Profiling

```sh
go test ./... -run '^$' -bench BenchmarkName -cpuprofile /tmp/cpu.pprof -memprofile /tmp/mem.pprof
go tool pprof /tmp/cpu.pprof
go tool pprof /tmp/mem.pprof
```

For blocking/mutex issues:

```sh
go test ./... -run TestName -blockprofile /tmp/block.pprof -mutexprofile /tmp/mutex.pprof
go tool pprof /tmp/block.pprof
go tool pprof /tmp/mutex.pprof
```

## Race and concurrency

When changing goroutines, channels, mutexes, cache, workers, timeouts, session lifecycle, or shared state:

```sh
go test -race ./...
```

Race detector proves only executed paths showed no data race. It does not prove concurrency correctness.

For deadlocks or scheduler issues:

```sh
go test ./... -run TestName -trace /tmp/trace.out
go tool trace /tmp/trace.out
```

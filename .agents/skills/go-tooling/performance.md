# Performance

Use this for cache, rendering, JSON/XML, SQL hot paths, allocations, auth/session lookup, static assets, middleware, or background work.

## Rule

Do not optimize without measurement.

Name the hot path first.
Use the simple version first.
Add cache only when there is a real reason and invalidation is clear.

## Benchmarks

```sh
go test ./... -run '^$' -bench . -benchmem
```

Compare with `benchstat`:

```sh
go test ./... -run '^$' -bench BenchmarkName -benchmem -count=10 > /tmp/before.txt
go test ./... -run '^$' -bench BenchmarkName -benchmem -count=10 > /tmp/after.txt
benchstat /tmp/before.txt /tmp/after.txt
```

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

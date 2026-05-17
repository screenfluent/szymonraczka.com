GO-1.26+ LLM CONTRACT

AUTHORITY ORDER
1. Project files: go.mod, go.work, CI matrix, build tags.
2. Official Go docs: go.dev release notes, pkg.go.dev stdlib docs, go command docs.
3. This contract.
4. Installed Go skills / local rules.
5. Model memory.

VERSION GATE
- First inspect go.mod/go.work. Never emit APIs newer than the module's `go` directive unless explicitly upgrading the module.
- Default target is Go 1.26.3+.
- If target < required API version, emit the fallback and say why.

LANGUAGE / SYNTAX
- Go >=1.22: loop variables are per-iteration for modules with `go 1.22+`; do not add `v := v` solely for closure capture.
- Go >=1.22: integer range is valid: `for i := range n { ... }`.
- Go >=1.23: use `iter.Seq` / `iter.Seq2` and range-over-func when it improves lazy iteration or API clarity.
- Go >=1.24: generic type aliases are stable.
- Go >=1.26: `new(expr)` is valid for pointer construction; use it for simple pointer literals/options when clearer than helper functions.
- Go >=1.26: self-referential generic constraints are allowed; do not reject them as invalid solely due to older memory.

TOOLS / MODULES
- Go >=1.24: pin executable tools with go.mod `tool` directives. Do not create new `tools.go` blank-import files.
- Prefer: `go get -tool <tool-package>@<version>`, `go tool <tool-name> ...`, `go install tool`.
- Use `go mod tidy` / `go mod tidy -diff` when available. For older toolchains, fallback to `go mod tidy && git diff --exit-code -- go.mod go.sum`.
- Do not use `go get` to install global binaries; use `go install pkg@version` for global tools or `go get -tool` for module-pinned tools.

TESTING / BENCHMARKING
- Go >=1.24: new benchmarks must use `for b.Loop() { ... }`, not `for i := 0; i < b.N; i++` and not `for range b.N`, except when preserving old compatibility.
- Go >=1.24: use `t.Context()`, `b.Context()`, `t.Chdir()`, `b.Chdir()` when they simplify test lifecycle.
- Go >=1.25: use `testing/synctest.Test(t, func(t *testing.T) { ... })` and `synctest.Wait()` for deterministic concurrent/time tests.
- Do not use Go 1.24 experimental `synctest.Run` in Go 1.25+ / 1.26+ code.
- Go >=1.25: use `t.Attr`, `b.Attr`, `f.Attr`, and `Output()` when useful for test metadata/output.
- Go >=1.26: use `t.ArtifactDir()`, `b.ArtifactDir()`, `f.ArtifactDir()` for files created by tests, benchmarks, fuzzers.
- Use `go test ./...`, `go test -race ./...` when concurrency is involved, and benchmark comparisons with `benchstat`.

CONCURRENCY
- Go >=1.25: prefer `sync.WaitGroup.Go` for simple fire-and-wait goroutines, but only when `f` must not panic and no error propagation/cancellation is needed.
- For errors, cancellation, limits, or first-error semantics, prefer `golang.org/x/sync/errgroup` with context.
- Do not claim `WaitGroup.Go` exists in Go 1.24.
- Go >=1.22 loop capture fix means redundant `item := item` is unnecessary unless target is Go <1.22.
- Go >=1.26: for leak diagnosis, know about experimental goroutine leak profiling (`GOEXPERIMENT=goroutineleakprofile` / pprof goroutineleak); do not treat it as default stable behavior.

FILESYSTEM / SECURITY
- Go >=1.24: use `os.Root` for user-controlled filesystem access and traversal resistance.
- Pre-Go 1.24 fallback must be lexical validation with `filepath.IsLocal` / `filepath.Rel` and separator-aware checks; never use `filepath.Clean(path)` + `strings.HasPrefix(full, base)` alone.
- For host/port construction, use `net.JoinHostPort(host, port)`, not `fmt.Sprintf("%s:%s", host, port)` or `fmt.Sprintf("%s:%d", host, port)`.
- Prefer confinement over sanitization for paths.

HTTP / NET
- Go >=1.22: use the enhanced `net/http.ServeMux` patterns and path variables where appropriate; avoid hand-parsing route segments when stdlib routing suffices.
- Go >=1.25: use `net/http.CrossOriginProtection` when implementing CSRF/cross-origin protections in stdlib HTTP services.
- Go >=1.26: `httputil.ReverseProxy.Director` is deprecated for new code; use `ReverseProxy.Rewrite`.
- Go >=1.26: know that `ServeMux` redirects changed from 301 to 307 for some non-GET requests; do not hardcode old assumptions.
- Go >=1.26: `net/url.Parse` is stricter for inputs beginning with `?`.

ERRORS
- Go >=1.26: prefer `errors.AsType[T](err)` for typed error extraction when `T` implements `error`.
- Keep `errors.As(err, &target)` for Go <1.26 and for non-error interface targets.
- Continue to use `errors.Is` for sentinel matching and `fmt.Errorf("context: %w", err)` for wrapping.

CRYPTO
- Go >=1.24: prefer stdlib `crypto/hkdf`, `crypto/pbkdf2`, `crypto/sha3` over `golang.org/x/crypto/...` equivalents when available.
- Go >=1.26: use `crypto/hpke` for HPKE instead of third-party HPKE by default.
- Go >=1.26: `crypto/rsa.EncryptPKCS1v15` is deprecated for new encryption use; prefer RSA-OAEP (`EncryptOAEP` / `EncryptOAEPWithOptions`) or modern KEM/HPKE design.
- Go >=1.26: several crypto APIs ignore randomness parameters for deterministic behavior; do not write tests that depend on those random parameters. Use `testing/cryptotest.SetGlobalRandom` where deterministic crypto tests require it.

JSON / ENCODING
- Go >=1.24: use `encoding/json` `omitzero` when zero-value omission is needed, especially for `time.Time`.
- `encoding/json/v2` is experimental behind `GOEXPERIMENT=jsonv2`; do not use it by default in production unless project explicitly opts in.

REFLECTION / SLOG / BYTES
- Go >=1.25: use `reflect.TypeAssert[T](v)` instead of `v.Interface().(T)` when working with reflection and type assertions.
- Go >=1.26: prefer new reflect iterators (`Type.Fields`, `Type.Methods`, `Type.Ins`, `Type.Outs`, `Value.Fields`, `Value.Methods`) where they simplify reflective code.
- Go >=1.26: use `bytes.Buffer.Peek` for non-consuming buffer inspection.
- Go >=1.26: use `log/slog.NewMultiHandler` for simple fan-out slog handling.

RUNTIME / PERFORMANCE
- Go >=1.25: GOMAXPROCS is container-aware by default; remove third-party automaxprocs workarounds unless the project has a measured reason.
- Go >=1.26: Green Tea GC is enabled by default. Re-evaluate old GC/allocator tuning with profiles; do not remove `GOMEMLIMIT` blindly because it remains useful for container memory ceilings.
- `runtime.GOROOT()` is deprecated; use `go env GOROOT` for tooling-level discovery.

MODERNIZATION / VALIDATION
- For migrations on a clean branch: run `go fix ./...`, `gofmt -w`, `go test ./...`, `go vet ./...`, `go mod tidy`.
- Prefer modernizers/lints backed by the target Go version; never modernize past the module's declared Go version.
- Red flags in generated code: new `tools.go`, `rand.Seed(time.Now...)`, new `b.N` benchmark loop, redundant loop var shadowing in Go 1.22+, `Clean+HasPrefix` path guard, manual host:port formatting, `ReverseProxy.Director` in new Go 1.26 code, `synctest.Run`, claiming WaitGroup.Go is Go 1.24.
```

## Minimal validation checklist for generated code

Run in the repository root:

```bash
go version
go env GOMOD GOWORK
go test ./...
go vet ./...
go mod tidy
git diff --exit-code -- go.mod go.sum
```

For migrations:

```bash
git switch -c modernize-go126
go fix ./...
gofmt -w .
go test ./...
go vet ./...
go mod tidy
```

For benchmarks:

```bash
go test -run '^$' -bench . -benchmem -count=10 ./... | tee bench-new.txt
benchstat bench-old.txt bench-new.txt
```

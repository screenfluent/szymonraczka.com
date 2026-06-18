---
name: go-tooling
description: Use Go toolchain and verification gates before claiming code works. Use when reading, changing, testing, verifying, building, securing, profiling, or reviewing Go code.
---

# Go Tooling

## Core rule

Do not guess when a tool can answer.

Read code, ask the Go toolchain, run the smallest relevant check, then report evidence.

Every completion report must say:
- what changed
- which files changed
- which commands passed
- what was not verified

## First tools

Start local and close to the code:

```sh
rg "symbol|endpoint|error|TODO" .
rg --files
go list ./...
go doc package.Symbol
go list -json ./...
gopls check ./...
```

Use `rg` before slow searching.
Use `go doc` before guessing API signatures.
Use `gopls` for symbols, imports, and type diagnostics.

See [code-reading.md](./code-reading.md).

## Ordinary verification gate

```sh
gofmt -w <changed-go-files>
gopls check ./...
go test ./... -run '<relevant-test>' -count=1
go test ./...
go vet ./...
staticcheck ./...
go build ./...
git diff --check
```

If `gopls` is unavailable, stop in `manual/` and `llm-generated/`.

See [quality-gates.md](./quality-gates.md) and [testing.md](./testing.md).

## Risk-based gates

Security/dependencies/auth:

```sh
govulncheck ./...
gosec ./...
```

Concurrency/shared state/cache/session lifecycle:

```sh
go test -race ./...
```

Database/migrations:

```sh
sqlite3 app.db 'PRAGMA integrity_check;'
sqlite3 app.db 'PRAGMA foreign_key_check;'
```

Performance:

```sh
go test ./... -run '^$' -bench . -benchmem
```

Deployment/shell/CI:

```sh
sh -n script.sh
shellcheck script.sh
actionlint
```

Use deeper files only when the task touches that risk:

- [context.md](./context.md) for `context.Context`, cancellation, timeouts, and request lifecycle
- [safety.md](./safety.md) for nil, slices, maps, numeric conversions, resources, goroutines, and defensive copying
- [dependencies.md](./dependencies.md) for `go.mod`, `go.sum`, tools, external packages, and vulnerability review
- [documentation.md](./documentation.md) for README, godoc, examples, ADRs, and AI-facing docs
- [security.md](./security.md)
- [database.md](./database.md)
- [operations.md](./operations.md)
- [performance.md](./performance.md)

## Reporting

Never say “works” without commands.

If a required check was not run, say why.
If a required check cannot run, the change is not complete.

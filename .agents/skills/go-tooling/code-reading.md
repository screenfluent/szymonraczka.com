# Code Reading

Use fast local tools before guessing.

## Search

```sh
rg "symbol|endpoint|error|TODO" .
rg --files
```

Read nearby code and nearby tests first. Do not read the whole project unless needed.

## Go package information

```sh
go list ./...
go list -json ./...
go env
go doc package.Symbol
go doc -all package
```

Use `go doc` instead of guessing signatures or standard library behavior.

## gopls

`gopls` is required in this workflow.
`gopls check` expects Go file paths, not `./...` package patterns.

```sh
go_files=($(git ls-files '*.go'))
if ((${#go_files[@]})); then gopls check "${go_files[@]}"; fi
gopls references path/to/file.go:line:col
gopls definition path/to/file.go:line:col
```

Treat diagnostics practically:

- errors: stop and fix
- warnings: inspect
- hints: optional, never a reason for chaotic refactor

## Debugging

Use a tight troubleshooting loop:

1. reproduce the failure
2. name one hypothesis
3. inspect the smallest relevant code path
4. use a debugger, trace, profile, or focused log to test that hypothesis
5. fix the root cause
6. verify the failing path and nearest regression test

If runtime behavior contradicts code reading, use debugger instead of random logs:

```sh
dlv test ./pkg/name -- -test.run TestName
dlv debug ./cmd/app
```

For races, deadlocks, leaks, or scheduler questions, prefer Go tools over guesswork:

```sh
go test -race ./...
go test ./... -run TestName -trace /tmp/trace.out
go tool trace /tmp/trace.out
```

Remove temporary logs and assumptions after debugging.

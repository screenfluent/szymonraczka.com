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

```sh
gopls check ./...
gopls references path/to/file.go:line:col
gopls definition path/to/file.go:line:col
```

Treat diagnostics practically:

- errors: stop and fix
- warnings: inspect
- hints: optional, never a reason for chaotic refactor

## Debugging

If runtime behavior contradicts code reading, use debugger instead of random logs:

```sh
dlv test ./pkg/name -- -test.run TestName
dlv debug ./cmd/app
```

Remove temporary logs and assumptions after debugging.

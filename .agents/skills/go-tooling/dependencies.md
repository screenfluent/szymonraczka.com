# Dependencies

Use this when `go.mod`, `go.sum`, tools, external packages, or vulnerability reports change.

## Rule

Ask before adding a dependency.

Cover:

- why the standard library or a small local function is not enough
- alternatives considered
- maintenance and API stability
- security and supply-chain risk
- whether the dependency is needed now or only speculative

Prefer standard library first. This repo should not grow a framework because a small behavior slice needed one helper.

## Common commands

Add a dependency deliberately:

```sh
go get example.com/module@latest
go mod tidy
go mod verify
go test ./...
```

Remove unused dependencies:

```sh
go mod tidy
go mod verify
go test ./...
```

Inspect dependencies:

```sh
go list -m all
go list -m -u all
go mod why -m example.com/module
go mod graph
```

Check vulnerabilities:

```sh
govulncheck ./...
```

## Tool dependencies

For project tools, prefer reproducible module-pinned tools over global assumptions.

With modern Go tool management, keep tool additions explicit and review `go.mod`/`go.sum` after changes. Do not add tool dependencies just because a reference skill mentions them.

If a tool is not installed locally, report that the gate was not run. Do not silently skip it and claim completion.

## Review checklist

When dependencies change, inspect the diff for:

- unexpected indirect dependencies
- module path changes or major-version suffixes
- `replace` directives
- new transitive native/system requirements
- license or security concerns when relevant

Do not vendor dependencies unless the task explicitly chooses vendoring.

## Reporting

Dependency reports must say:

- dependency added/removed/updated
- reason it is needed
- commands that passed
- vulnerability check result or why it was not run
- any remaining risk

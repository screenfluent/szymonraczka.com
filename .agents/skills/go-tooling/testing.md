# Testing

Tests should verify behavior through public interfaces.

Good tests:
- describe what the system does
- use public API / HTTP / exported behavior
- survive internal refactors
- avoid mocks of internal code

Bad tests:
- mock internal collaborators
- test private helpers directly
- assert call order or implementation details
- break when implementation changes but behavior stays the same

## Loop

Use vertical slices:

```txt
RED → GREEN → optional refactor → verify
```

Do not write all tests first and all implementation later.

## Commands

Target nearest test first:

```sh
go test ./... -run 'TestSpecificCase' -count=1
```

Then full suite:

```sh
go test ./...
```

For flaky/order-sensitive logic:

```sh
go test ./... -shuffle=on -count=10
```

For parsers, slugs, tokens, URLs, and user-controlled input, consider fuzzing:

```sh
go test ./... -run '^$' -fuzz=FuzzName -fuzztime=30s
```

Coverage is a map, not a goal:

```sh
go test ./... -cover
go test ./... -coverprofile=/tmp/coverage.out
go tool cover -func=/tmp/coverage.out
```

Do not use coverage percentage as proof of safety.

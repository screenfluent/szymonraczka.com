# Quality Gates

## Formatting

After Go changes:

```sh
gofmt -w <changed-go-files>
```

For larger changes:

```sh
gofmt -w .
```

Use `goimports` only if the project already uses it.

## Required ordinary gate

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

`go vet` results are errors, not suggestions.

`staticcheck` usually has high signal. Do not silence it without a concrete reason.

## Lint discipline

Do not add a new lint tool or config unless the task is explicitly about linting, CI, or a defect the current gates miss.

When suppressing a lint warning:

- prefer fixing the code first
- use the narrowest suppression available
- name the linter when the tool supports it
- write a concrete reason, not `//nolint` with no explanation

If a large legacy cleanup is needed, split by package or linter. Do not mix lint churn with behavior changes.

## Build

```sh
go build ./...
```

If the app has several entrypoints:

```sh
go build ./cmd/...
```

For production artifact checks when relevant:

```sh
go build -trimpath -ldflags="-s -w" -o /tmp/app .
```

Build matters for embedded templates/static files.

## Diff hygiene

```sh
git diff --check
```

Catches whitespace errors, conflict markers, and accidental diff garbage.

# Operations

Use this for shell, CI, Docker, deploy, runtime smoke tests, and generated artifacts.

## Shell

```sh
sh -n script.sh
shellcheck script.sh
shfmt -w script.sh
```

## GitHub Actions

CI should mirror the local verification gate before adding extra policy. Add workflow complexity only when the task needs it.

```sh
actionlint
```

For Go CI, start with the ordinary local checks: `go test`, `go vet`, `staticcheck` if installed in CI, and `go build`. Add `govulncheck`, race tests, coverage upload, or release automation only for a concrete risk or workflow need.

## Docker

```sh
docker build .
trivy image image-name
```

## Runtime smoke

For HTTP, auth, cookies, reverse proxy, healthchecks, or config:

```sh
curl -i http://127.0.0.1:PORT/healthz
curl -i http://127.0.0.1:PORT/
```

Smoke test should prove:

- process starts
- route works
- basic headers/config are sane
- obvious deploy/config errors appear before release

## Observability

Start with the smallest useful signal:

- structured logs with `log/slog` for important events and errors
- request IDs when debugging multi-step request flow
- metrics/tracing only when there is a consumer and an operational question

Do not add Prometheus, OpenTelemetry, log aggregation, or dashboards speculatively. Do not log secrets, tokens, passwords, reset links, or unnecessary personal data.

## Generated code

If project uses generators:

```sh
go generate ./...
sqlc generate
buf generate
protoc ...
swag init
```

After generation, inspect diff. Generated churn can mean local tool version differs from CI.

## Do not mutate real infrastructure

Do not run real deploy/apply commands without explicit user approval.

Prefer:

- plan
- validate
- dry-run
- local build
- smoke tests

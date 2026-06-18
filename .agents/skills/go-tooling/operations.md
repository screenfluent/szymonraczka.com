# Operations

Use this for shell, CI, Docker, deploy, runtime smoke tests, and generated artifacts.

## Shell

```sh
sh -n script.sh
shellcheck script.sh
shfmt -w script.sh
```

## GitHub Actions

```sh
actionlint
```

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

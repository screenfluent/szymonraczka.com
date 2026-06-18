# Security

Use security tools when changing dependencies, auth, cookies, tokens, SQL, file handling, shell/deploy, or user input parsing.

## Vulnerabilities

```sh
govulncheck ./...
```

Focus on vulnerabilities in called symbols. If a module is vulnerable but the affected symbol is not called, report it as risk, not active exploitability.

## Static security scan

```sh
gosec ./...
go tool gosec ./...
```

Treat `gosec` as a hypothesis to inspect, not automatic truth.

False positives must be explained concretely.

## Dependency review

When `go.mod` or `go.sum` changes:

```sh
go mod verify
go mod tidy
go list -m -u all
govulncheck ./...
```

After `go mod tidy`, inspect the diff.

## Secrets

When changing config, logging, deploy, fixtures, or docs:

```sh
gitleaks detect --source .
trivy fs --scanners secret .
```

Do not put real secrets in tests, docs, examples, or skills.

## Web security checklist

For web code, check:

- SQL injection: parameterized queries
- XSS: HTML/template/attribute/URL escaping
- CSRF: token plus Origin/Referer for mutations
- auth: random sessions, hashed storage, secure cookies
- reset/verify tokens: random, hashed, expiring, single-use
- redirects: no open redirect
- SSRF/local URLs: validate external URLs
- rate limit: IP/user/account scoped
- cache: no public cache for private responses
- logs: redact tokens/passwords/reset links
- headers: CSP, nosniff, frame deny, referrer policy
- proxy trust: forwarded headers only from trusted peers
- pprof/metrics: not public unless intentionally protected

# cc-skills-golang

Vendored curated subset from `https://github.com/samber/cc-skills-golang`.

Source commit: `a5e0e5997aac169b659e70cb826a20b489bc4c6c`
Copied: 2026-06-18
License: MIT, see `LICENSE`.

Included skills:

- `golang-code-style`
- `golang-naming`
- `golang-error-handling`
- `golang-testing`
- `golang-structs-interfaces`
- `golang-concurrency`
- `golang-security`
- `golang-database`
- `golang-context`
- `golang-safety`
- `golang-dependency-management`
- `golang-documentation`
- `golang-troubleshooting`

Use these as Go reference material only. Project workflow is still governed by `transfer-slice` and `go-tooling`.

Each vendored skill has `disable-model-invocation: true` so it does not compete with the local workflow. Load/read these files explicitly when deeper Go reference is useful.

Local `go-tooling` also contains compact workflow-focused notes distilled from the broader Samber set, especially for context, safety, dependencies, documentation, troubleshooting, linting, performance, observability, and CI.

Some upstream skill files reference companion cc-skills that are not included in this curated subset. Do not auto-vendor those unless `../SOURCE.md` is updated intentionally.

Intentionally not included now:

- framework/library-specific skills: `golang-graphql`, `golang-grpc`, `golang-swagger`, `golang-spf13-cobra`, `golang-spf13-viper`, `golang-google-wire`, `golang-uber-dig`, `golang-uber-fx`, `golang-samber-*`
- architecture-expanding skills: `golang-project-layout`, `golang-design-patterns`, `golang-dependency-injection`, `golang-popular-libraries`, `golang-how-to`
- workflow/tooling skills already covered locally in `go-tooling`: `golang-lint`, `golang-benchmark`, `golang-performance`, `golang-observability`, `golang-continuous-integration`
- broad learning/reference skills not needed for this starter workflow: `golang-data-structures`, `golang-modernize`, `golang-stay-updated`

Do not auto-update. Update `../SOURCE.md` when vendored files are added, removed, or changed.

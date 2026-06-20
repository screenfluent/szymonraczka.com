# Skill Sources

## Local skills

* transfer-slice
    * Purpose: Manual transfer of small behavior slices from ../llm-generated/ into manual/.
    * Last edited: 2026-06-19
* go-tooling
    * Purpose: Go toolchain usage, verification gates, and evidence-based reports.
    * Last edited: 2026-06-18
* sm20-flashcards
    * Purpose: Generate and approve SM-20-style atomic flashcards from LEARNING.md learning events.
    * Last edited: 2026-06-18

## Vendored skill packs

### mattpocock/skills

Source: https://github.com/mattpocock/skills
Reference release: mattpocock-skills@1.0.0
Last checked: 2026-06-18
Local changes: flattened skill directories to `.agents/skills/mattpocock-*`, prefixed `name:` fields so MastraCode can discover them, and updated local cross-skill references

Included in manual:

* mattpocock-ask-matt
* mattpocock-codebase-design
* mattpocock-diagnosing-bugs
* mattpocock-domain-modeling
* mattpocock-grill-with-docs
* mattpocock-grilling
* mattpocock-handoff
* mattpocock-tdd
* mattpocock-writing-great-skills

Notes:

* These skills are vendored under `.agents/skills/mattpocock-*`.
* Keep upstream skill content unchanged unless a local change is intentional.
* If any vendored skill content is edited locally, update Local changes with a short explanation.
* Do not auto-update vendored skills.
* Update this file when skills are added, removed, renamed, or changed.

### cc-skills-golang

Source: https://github.com/samber/cc-skills-golang
Reference commit: a5e0e5997aac169b659e70cb826a20b489bc4c6c
Copied: 2026-06-18
Last checked: 2026-06-18
Local changes: added `disable-model-invocation: true`; flattened skill directories to `.agents/skills/cc-skills-golang-*`; prefixed `name:` fields; moved pack README/LICENSE to `.agents/skills/cc-skills-golang-README.md` and `.agents/skills/cc-skills-golang-LICENSE`

Included in manual:

* cc-skills-golang-golang-code-style
* cc-skills-golang-golang-naming
* cc-skills-golang-golang-error-handling
* cc-skills-golang-golang-testing
* cc-skills-golang-golang-structs-interfaces
* cc-skills-golang-golang-concurrency
* cc-skills-golang-golang-security
* cc-skills-golang-golang-database
* cc-skills-golang-golang-context
* cc-skills-golang-golang-safety
* cc-skills-golang-golang-dependency-management
* cc-skills-golang-golang-documentation
* cc-skills-golang-golang-troubleshooting

Notes:

* These selected skills are vendored under `.agents/skills/cc-skills-golang-*`.
* Use them as Go reference material only; project workflow remains governed by transfer-slice and go-tooling.
* They are hidden from automatic model invocation with `disable-model-invocation: true`; load/read them explicitly when deeper reference is useful.
* Local go-tooling also contains compact workflow-focused notes distilled from broader Samber topics: context, safety, dependencies, documentation, troubleshooting, linting, performance, observability, and CI.
* Skills intentionally not included: framework/library-specific skills, architecture-expanding skills, broad learning/reference skills, and workflow/tooling skills already covered locally in go-tooling.
* Keep upstream skill content unchanged unless a local change is intentional.
* If any vendored skill content is edited locally, update Local changes with a short explanation.
* Do not auto-update vendored skills.
* Update this file when skills are added, removed, renamed, or changed.

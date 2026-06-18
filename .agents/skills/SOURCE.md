# Skill Sources

## Local skills

* transfer-slice
    * Purpose: Manual transfer of small behavior slices from ../llm-generated/ into manual/.
    * Last edited: 2026-06-18
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
Local changes: none

Included in manual:

* ask-matt
* codebase-design
* diagnosing-bugs
* domain-modeling
* grill-with-docs
* grilling
* handoff
* tdd
* writing-great-skills

Notes:

* These skills are vendored under .agents/skills/mattpocock/.
* Keep upstream skills unchanged unless a local change is intentional.
* If any vendored skill is edited locally, change Local changes from none to a short explanation.
* Do not auto-update vendored skills.
* Update this file when skills are added, removed, or changed.

### cc-skills-golang

Source: https://github.com/samber/cc-skills-golang
Reference commit: a5e0e5997aac169b659e70cb826a20b489bc4c6c
Copied: 2026-06-18
Last checked: 2026-06-18
Local changes: added `disable-model-invocation: true` to vendored skill files so cc-skills-golang stays reference-only; local README documents the curated subset

Included in manual:

* golang-code-style
* golang-naming
* golang-error-handling
* golang-testing
* golang-structs-interfaces
* golang-concurrency
* golang-security
* golang-database
* golang-context
* golang-safety
* golang-dependency-management
* golang-documentation
* golang-troubleshooting

Notes:

* These selected skills are vendored under .agents/skills/cc-skills-golang/.
* Use them as Go reference material only; project workflow remains governed by transfer-slice and go-tooling.
* They are hidden from automatic model invocation with `disable-model-invocation: true`; load/read them explicitly when deeper reference is useful.
* Local go-tooling also contains compact workflow-focused notes distilled from broader Samber topics: context, safety, dependencies, documentation, troubleshooting, linting, performance, observability, and CI.
* Skills intentionally not included: framework/library-specific skills, architecture-expanding skills, broad learning/reference skills, and workflow/tooling skills already covered locally in go-tooling.
* Keep upstream skill files unchanged unless a local change is intentional.
* If any vendored skill is edited locally, change Local changes from none to a short explanation.
* Do not auto-update vendored skills.
* Update this file when skills are added, removed, or changed.

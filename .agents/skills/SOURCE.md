# Skill Sources

## Local skills

- transfer-slice
  - Purpose: Manual transfer of small behavior slices from `../llm-generated/` into `manual/`.
  - Last edited: 2026-06-18

- go-tooling
  - Purpose: Go toolchain usage, verification gates, and evidence-based reports.
  - Last edited: 2026-06-17

## Local adapters for Matt Pocock skills v1

These are local, compact adapters aligned with Matt Pocock skills v1. Replace them with the official upstream files when vendoring the package.

- grill-with-docs
  - Purpose: User-invoked grilling session that uses domain modeling to update `CONTEXT.md` and ADRs lazily.
  - Last edited: 2026-06-18

- grilling
  - Purpose: Model-invoked relentless decision-tree interview loop.
  - Last edited: 2026-06-18

- domain-modeling
  - Purpose: Model-invoked domain glossary and ADR maintenance discipline.
  - Last edited: 2026-06-18

- codebase-design
  - Purpose: Model-invoked design vocabulary for deep modules, clean seams, and testable public interfaces.
  - Last edited: 2026-06-18

- tdd
  - Purpose: Model-invoked red-green-refactor discipline for behavior tests and vertical slices.
  - Last edited: 2026-06-18

- writing-great-skills
  - Purpose: User-invoked guidance for creating or editing skills with predictable structure and compact descriptions.
  - Last edited: 2026-06-18

## Vendored skill packs

### mattpocock/skills

Source: https://github.com/mattpocock/skills
Reference release: v1.x, checked 2026-06-18
Local changes: using local compact adapters, not exact upstream copies

Manual includes selected v1-aligned skills:
- grill-with-docs
- grilling
- domain-modeling
- codebase-design
- tdd
- writing-great-skills

Notes:
- Keep these compact unless exact upstream vendoring is intentionally performed.
- Do not auto-update vendored skills.
- Update this file when skills are added, removed, or changed.

### cc-skills-golang

Source: https://github.com/samber/cc-skills-golang
Copied: pending
Last checked: pending
Local changes: none

Included in manual when vendored:
- golang-code-style
- golang-naming
- golang-error-handling
- golang-testing
- golang-structs-interfaces
- golang-concurrency
- golang-security
- golang-database

Notes:
- Copy selected skills into `.agents/skills/cc-skills-golang/` manually.
- Do not auto-update vendored skills.
- Update this file when skills are added, removed, or changed.

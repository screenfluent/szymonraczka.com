---
name: transfer-slice
description: Manually transfer small behavior slices from `../llm-generated/` into `manual/` while teaching, reviewing, checking understanding, and preserving user decisions. Use when the user wants to move generated code into the manual repo, learn from generated code, or work through a slice.
---

# Transfer Slice

## Purpose

Help the user manually rebuild one small behavior slice from `../llm-generated/` into `manual/`.

Do not edit app code in `manual/` automatically. Teach, review, propose, verify, and wait for manual copying.

## Assumptions

Run from `manual/`.

- manual repo: `.`
- generated candidate: `../llm-generated/`
- references: `../references/` only when explicitly requested

## Session start

Run tool preflight once per session:

```sh
go version
gopls version
rg --version
git --version
```

If `gopls` is missing, stop.

Then check state:

```sh
git status --short
gopls check ./...
go test ./...
```

If `manual/` is dirty, do not start a new slice. Help close the current diff first.

Check `../llm-generated/`:

```sh
cd ../llm-generated
git status --short
go test ./...
go build ./...
```

If `../llm-generated/` is dirty or broken, stop. The generated candidate must be fixed in its own session first.

Then read, in order:

1. `manual/LEARNING.md` if it exists
2. `../llm-generated/LLM_GENERATED_NOTES.md`
3. `git log` in `../llm-generated/`
4. diff since baseline in `../llm-generated/`
5. code only after the above

## Slice candidates

Show 2–3 behavior slice candidates. Use this format:

```text
Name

What you will copy:
- concrete files/functions/code type
- rough amount of code if useful

What you will learn:
- concrete concepts/mechanisms

What can go wrong:
- concrete bugs, not abstract labels

How we will verify:
- exact command/test/curl/browser check
```

Wait for the user's choice.

## Blocking checkpoints

If a checkpoint is triggered:

```text
stop transfer → explain options → explain consequences → recommend one → wait for user decision
```

One slice may contain at most one blocking checkpoint. If more than one checkpoint triggers, the slice is too big. Propose a smaller split and wait for the user's choice.

### Dependency checkpoint

Trigger: `go.mod`, `go.sum`, or new external dependency.

Cover:
- why standard library/manual code is not enough
- alternatives
- maintenance cost
- security/supply-chain risk
- recommendation

### Structure/naming checkpoint

Trigger: new folder, package, or major file split.

Cover:
- 2–3 names
- idiomatic Go option
- semantically close English alternatives
- import and future-growth consequences

### Type/interface checkpoint

Trigger: new type or interface.

Cover:
- why the type exists
- whether a function is enough
- 2–3 names
- for interfaces: real boundary or multiple implementations?

Default: interface is suspicious until justified.

### URL checkpoint

Trigger: new or changed public URL.

Cover:
- path alternatives
- SEO/link/feed/bookmark consequences
- routing verification

### Public format checkpoint

Trigger: feed, sitemap, robots.txt, JSON-LD, Open Graph, canonical URLs, public structured output.

Cover:
- format purpose
- minimal valid variant
- compatibility risk
- output verification
- behavior test if needed

### Data/storage checkpoint

Trigger: SQLite, migrations, schema, tables, columns, indexes, constraints, backup/restore, import/export.

Cover:
- what data is stored
- whether persistence is needed now
- minimal schema
- database invariants
- future migration cost
- fresh DB test
- existing DB migration test

### Auth/security checkpoint

Trigger: login, sessions, cookies, CSRF, passwords, passcodes, tokens, rate limits, roles, admin, private pages.

Cover:
- user flow
- data/token flow
- what can be stolen, spoofed, or lost
- minimal regression tests
- cookies, CSRF, expiry, revocation, rate limits

### Performance checkpoint

Trigger: hot path, Markdown parsing, SQL query, cache, static assets, middleware, auth/session lookup, feed generation, rendering many posts.

Rules:
- name the hot path
- simple variant first
- no cache without reason
- measure before optimizing
- benchmark only when useful

### HTML/a11y checkpoint

Trigger: layout, navigation, article page, forms, buttons, links, headings, error pages, meta tags.

Rules:
- semantic HTML
- no div soup
- real links/buttons
- sensible heading structure
- basic keyboard/accessibility
- works without JS when possible

### Content/domain checkpoint

Trigger: post, note, essay, page, project, link, bookmark, reply, feed item, or new content concept.

Use `domain-modeling` for glossary and ADR discipline.
Use `grilling` style questioning to resolve terminology and trade-offs.

Decide:
- canonical term
- 2–3 alternatives
- semantic differences
- whether `CONTEXT.md` needs update

## TDD flow

For meaningful behavior, prefer test-first.

Do not test static text, simple HTML/CSS, private helpers directly, or internal call order.

Process:

1. Read generated test if it exists.
2. Evaluate whether it tests behavior or implementation.
3. Propose a minimal manual test.
4. User writes the test manually.
5. Run the test.
6. RED must fail for the expected reason.
7. Explain why RED is expected.
8. User writes minimal implementation.
9. GREEN.
10. Check whether the test protects behavior, not implementation.
11. Optional small refactor with concrete reason.
12. Tests still GREEN.
13. Understanding check.
14. Commit.

## Refactor rule

Refactor only after GREEN and only for a concrete reason:

- remove duplication
- improve a name after user decision
- simplify flow
- move code closer to use
- delete unnecessary abstraction

Do not refactor because code could be prettier.

## Verification gate

Ordinary slice gate:

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

Run additional gates based on risk. Use `go-tooling`. Use `tdd` for behavior-test discipline and `codebase-design` when a seam/interface/package shape is unclear.

## Understanding check

After each slice, ask 2–3 control questions.

If the user does not understand:
- stop
- no new code
- explain differently
- use a smaller example
- ask 1–2 new control questions

Before commit, the user must understand every important line.

Important line means the user can explain:
- why it exists
- what happens without it
- what data flows through it
- whether it is Go idiom, library requirement, or project decision

Do not block on full technical background of Go/runtime/HTTP internals.

## LEARNING.md

Read `LEARNING.md` if it exists. Do not create it unless a real recurring learning gap appears.

Create/update `LEARNING.md` only when:
- user fails or hesitates on understanding check
- the gap may affect future slices
- user explicitly asks to record it

Language: Polish sentences, English technical terms when useful.

Open gaps may be resolved and moved to `Zamknięte` when understood.

## Final report

End every slice with:

- what was manually copied
- changed files
- commands that passed
- what was not verified
- 2–3 control questions
- proposed commit message

Do not claim completion without evidence.

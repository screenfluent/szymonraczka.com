# Documentation

Use this when changing README, godoc comments, examples, ADRs, public APIs, setup instructions, operational docs, or AI-facing project notes.

## Rule

Document decisions and usage, not obvious code.

Good documentation answers:

- what this is
- when to use it
- how to run or verify it
- what constraints matter
- what decision was made and why

Bad documentation repeats implementation line-by-line or promises behavior that tests do not cover.

## Go doc comments

Exported package names, types, funcs, methods, vars, and consts should have useful comments when they are part of a public API.

Start comments with the exported identifier when practical:

```go
// Server handles HTTP requests for the site.
type Server struct { ... }
```

Do not add noisy comments to unexported helpers unless they explain a non-obvious invariant, side effect, or trade-off.

## README for an app

A minimal app README should cover:

- project purpose
- prerequisites
- how to run locally
- how to test
- important configuration
- deployment or operational notes when they exist

Keep it current with the actual commands used in this repo.

## Examples and tests

Use Go example tests when they clarify public API usage and can be executed by `go test`.

Do not create examples for private helpers or static text.

## Decisions

Use ADRs or project notes for decisions that future agents may otherwise reopen:

- public URL shape
- content/domain vocabulary
- persistence model
- auth/session/security choices
- dependency choices
- deployment constraints

Short decision records are better than long speculative design docs.

## AI-facing docs

For agent instructions and skills:

- keep one source of truth
- prefer pointers to detailed files over repeating rules
- update source/last-edited metadata when skill content changes
- make workflow precedence explicit

## Verification

For documentation-only changes, run:

```sh
git diff --check
```

For docs that include commands, run the commands when practical or say clearly that they were not verified.

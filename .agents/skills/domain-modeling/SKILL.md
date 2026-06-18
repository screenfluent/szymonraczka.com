---
name: domain-modeling
description: Build and sharpen the project domain model. Use when terms, concepts, boundaries, glossary entries, CONTEXT.md, or ADR-worthy decisions appear.
---


# Domain Modeling

Keep project language precise.

## Read first

- `CONTEXT.md` if it exists
- `CONTEXT-MAP.md` if it exists
- relevant ADRs under `docs/adr/` if they exist

If files do not exist, proceed silently.
Create them lazily only when real terms or decisions are resolved.

## CONTEXT.md

`CONTEXT.md` is a glossary only.
No implementation details.
No process notes.
No speculative terms.

When adding a term:
- use English
- define what it is in 1–2 sentences
- list avoided synonyms when useful

## ADRs

Offer an ADR only when all are true:
- hard to reverse
- surprising without context
- real trade-off

Keep ADRs short.

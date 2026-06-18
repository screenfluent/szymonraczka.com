---
name: grill-with-docs
description: Relentlessly stress-test a plan against project language and documented decisions. Use when the user wants to grill a design while updating CONTEXT.md and ADRs as decisions crystallize.
disable-model-invocation: true
---


# Grill With Docs

Run a grilling session using `grilling` and `domain-modeling`.

## Rules

- Ask one question at a time.
- Give your recommended answer with each question.
- Use `CONTEXT.md` if it exists.
- Use `docs/adr/` if it exists.
- Create `CONTEXT.md` lazily only when the first real domain term is resolved.
- Create `docs/adr/` lazily only when a decision is hard to reverse, surprising without context, and a real trade-off.
- Keep `CONTEXT.md` as glossary only. No implementation details.
- Keep ADRs short.

## When terms appear

Use `domain-modeling`:
- challenge fuzzy terms
- propose precise canonical terms
- list words to avoid
- update the glossary inline when the user decides

## When decisions appear

Offer an ADR only when it would prevent future confusion.
Do not create ADRs for obvious, easy-to-reverse, or ceremonial choices.

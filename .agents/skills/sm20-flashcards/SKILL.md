---
name: sm20-flashcards
description: Generate and approve SM-20-style atomic flashcards from LEARNING.md.
disable-model-invocation: true
---

# SM20 Flashcards

Use this only in the manual learning repo.

## Inputs

Read:

- `LEARNING.md`
- `learning/state.json` if present
- `learning/card-candidates.jsonl` if present
- `learning/cards.jsonl` if present

Create missing `learning/` files only after explaining the planned files and getting approval.

## Eligible events

Process only `LEARNING.md` events with:

```txt
status: understood
cards: pending
```

Ignore `needs_work`, `generated`, and `skipped` events.

## Generate candidates

For each eligible event:

1. preserve learner-specific nuance from `tutor_notes`
2. make a tiny concept map
3. generate only atomic card candidates
4. append candidates to `learning/card-candidates.jsonl`
5. update `learning/state.json`

Do not write to `learning/cards.jsonl` without user approval.

## Approval

When the user approves or edits candidates:

- append approved cards to `learning/cards.jsonl`
- append decisions to `learning/card-decisions.jsonl`
- update `learning/state.json`

When the user rejects candidates:

- append rejection decisions to `learning/card-decisions.jsonl`
- do not create active cards

## Export

Export approved cards to `learning/exports/`.

SuperMemo API is a scheduler boundary only. Local files own card content.

## Protocol

Use [PROTOCOL.md](./PROTOCOL.md).

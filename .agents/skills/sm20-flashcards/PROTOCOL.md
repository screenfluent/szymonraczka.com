# SM20 Flashcard Protocol

## Core rule

No final card before understood evidence.

If evidence is weak, output:

```txt
TEACH MORE
```

Then name the missing distinction.

## Atomicity

One card tests one retrieval target.

Reject cards that ask for:

- what + why
- definition + example
- multiple blanks
- lists unless the list itself is the one target
- partial credit
- multi-step reasoning

## Allowed card types

- `cloze`
- `output_prediction`
- `conceptual_inversion`

## Candidate JSONL object

Each candidate line:

```json
{
  "candidate_id": "C-YYYY-MM-DD-NNN",
  "source_learning_event_id": "LE-YYYY-MM-DD-NNN-slug",
  "card_type": "cloze|output_prediction|conceptual_inversion",
  "prompt": "...",
  "answer": "...",
  "cloze_text": "... or null",
  "context_hint": "... or null",
  "reference_title": "... or null",
  "reference_source": "... or null",
  "tags": [],
  "grading_mode": "graded_or_binary",
  "atomicity_check": {
    "tested_targets": 1,
    "unscored_anchors_max": 2
  },
  "status": "pending"
}
```

## Approved card JSONL object

Each approved card line:

```json
{
  "card_id": "CARD-YYYY-MM-DD-NNN",
  "source_candidate_id": "C-YYYY-MM-DD-NNN",
  "source_learning_event_id": "LE-YYYY-MM-DD-NNN-slug",
  "card_type": "cloze|output_prediction|conceptual_inversion",
  "prompt": "...",
  "answer": "...",
  "cloze_text": "... or null",
  "context_hint": "... or null",
  "reference_title": "... or null",
  "reference_source": "... or null",
  "tags": [],
  "grading_mode": "graded_or_binary",
  "atomicity_check": {
    "tested_targets": 1,
    "unscored_anchors_max": 2
  },
  "created_at": "YYYY-MM-DDTHH:MM:SSZ"
}
```

## Decision JSONL object

Each decision line:

```json
{
  "decision_id": "D-YYYY-MM-DD-NNN",
  "candidate_id": "C-YYYY-MM-DD-NNN",
  "decision": "approved|edited|rejected",
  "reason": "...",
  "created_at": "YYYY-MM-DDTHH:MM:SSZ"
}
```

## Scheduler boundary

Local files own:

- prompt
- answer
- cloze text
- tags
- references
- learner-specific notes

SuperMemo API owns:

- review scheduling
- grade events
- intervals
- retention policy

Do not invent undocumented API payload fields.

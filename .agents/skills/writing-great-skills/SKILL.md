---
name: writing-great-skills
description: Reference for writing and editing skills well. Use when creating, updating, splitting, or reviewing agent skills.
disable-model-invocation: true
---


# Writing Great Skills

Create skills that are predictable, compact, and easy for a model to route to.

## Requirements

- `description` must clearly say what the skill does and when to use it.
- `SKILL.md` should be useful by itself.
- Split supporting material only when it is task-specific or rarely needed.
- Do not hide the core behavior in a giant reference file.
- Prefer concrete workflows over motivational prose.
- Keep examples small.

## Structure

```txt
skill-name/
├── SKILL.md
├── topic.md      # optional deeper material
└── scripts/      # optional deterministic helpers
```

## Before editing skills

Check `.agents/skills/SOURCE.md`.
If skill content changes, update `Last edited`.
If vendored content changes, mark `Local changes`.

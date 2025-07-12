---
id: task-1
title: Create content loading system
status: In Progress
assignee: []
created_date: '2025-07-12'
updated_date: '2025-07-12'
labels: []
dependencies: []
---

## Description

Build utilities to load and parse markdown files with frontmatter from /content directory to support the diary-focused architecture

## Acceptance Criteria

- [ ] Content loader can read markdown files from /content/{lang}/diary/
- [ ] Frontmatter parsing extracts title/date/tags
- [ ] System handles YYYYMMDDHHMM filename prefixes
- [ ] Error handling for missing or malformed files

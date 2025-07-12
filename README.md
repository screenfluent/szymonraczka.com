# Szymon Raczka's Personal Site

A transparent personal website built with SvelteKit, serving raw diary entries, tech resources for non-tech founders, and open metrics. Emphasizes authenticity: uncensored reflections on life/burnout/success, practical tech explanations with everyday analogies, and static pages for tools/status/story.

Key sections:

- `/diary`: Personal reflections (burnout, 165M YouTube views journey, work-from-car happiness).
- `/tech`: Tech for non-tech founders.
  - `/tech/howtos`: Step-by-step guides (e.g., deploy to Netlify).
  - `/tech/glossary`: Term explanations with analogies (e.g., package.json as kitchen shopping list).
- Static: `/uses`, `/now`, `/about` (or `/story`), `/open` (debts/incomes/shocking facts).
- Bilingual: English default (no prefix), Polish via `/pl` (e.g., `/pl/diary/slug`).

Content managed in root `/content/{lang}/{domain}/*.md` (e.g., `/content/en/diary/202507121430-burnout.md` – YYYYMMDDHHMM prefix for chronology). Frontmatter for title/date/tags.

## Setup

Clone repo and install dependencies with Bun:

```bash
bun install
```

## Develop

Run dev server (opens at http://localhost:5173):

```bash
bun run dev
```

Sync types and check:

```bash
bun run prepare
bun run check
```

Format and lint:

```bash
bun run format
bun run lint
```

## Build

Build production version:

```bash
bun run build
```

Preview build locally:

```bash
bun run preview
```

## Deploy

Deploys to Netlify via adapter. Push to repo with Netlify integration; site live at szymonraczka.com (after DNS setup pointing domain to your Netlify site).

For manual deploy: Use Netlify CLI or drag/drop build folder.

## Tech Stack

- SvelteKit 2 (runes, mdsvex for markdown).
- TailwindCSS.
- Bun for package management.
- Deploy: Netlify.

Details in PRD.md. Questions? Edit and push – keep it lean!

# Product Requirements Document (PRD) for Szymon Raczka's Personal Site

## 1. Introduction

### Overview

This PRD defines the requirements for a personal website at szymonraczka.com, serving as a transparent platform for Szymon Raczka (44, solo indie dev with ADHD) to share raw personal reflections, technical insights for non-tech founders, and open metrics. Inspired by Rick Rubin's philosophy:

> "This is something that I am making for myself for now. That is all it is.
>
> It is a diary entry. Everything I make is a diary entry. The beauty of a diary entry is that you can't tell me my diary entry is not good enough or that it's not what I experienced. I am writing the diary for myself. No one else can judge it. It is my experience of myself and my life. Everything we make can be that.
>
> Everything we make can be a personal reflection of who we are in that moment of time. It doesn't have to be for anyone else. It doesn't have to be the greatest you could ever do. It doesn't have to have an expectation to change the world. It doesn't have to sell a certain number of copies for any reason. It doesn't have to be any of those things.
>
> I'm making this thing for me, and I want to do it to the best of my ability to where I feel good about it and where it's honest. It's honest of where I am at.
>
> If you are living this world of just being honest of where you are at. There is nothing blocking that. There are no blocks."

The site emphasizes authenticity: all content as "diary entries" – uncensored snapshots of life, burnout, success, and tech learnings with personal anecdotes (e.g., laughing at past confusion over package.json); practical explanations with everyday analogies (e.g., codebase as restaurant kitchen); and static pages for current status and tools.

The site is built with simplicity in mind: minimal features to solve today's problems (publishing content without cognitive overload), organized by tags within a single diary domain, and bilingual (English as default, Polish as secondary).

### Goals

- Enable quick publishing of "diary entries" (raw thoughts on burnout, 165M YouTube views success, work-from-car happiness, perfectionism escape, tech journeys) to build audience resonance through transparency and honesty.
- Provide tech resources as diary entries for non-tech founders: step-by-step guides and term explanations using simple analogies from daily life (kitchen, work, restaurant) to demystify concepts like codebases or AI-assisted building, infused with personal reflections (e.g., "honest of where I am at" – sharing the journey from brain-melting confusion to current ease).
- Share open metrics (debts, future incomes, shocking personal facts) to inspire and shock "average Joes."
- Support bilingual content: English as primary (no prefix in URLs), Polish as optional (with top-level /pl prefix, e.g., /pl/diary/slug) for easier initial drafting in Polish then translation.
- Minimize maintenance: single source of truth for content in root-level /content folder (subdivided by language, all under diary), no deep nesting, predictable naming.

### Target Audience

- Indie hackers and non-tech founders starting with code/AI (confused by basics like package.json).
- Readers seeking raw inspiration from Szymon's journey: from TV/online success to depression/debts rock bottom and rebuild.
- Bilingual users: English for global reach, Polish for personal expression.

### Scope Exclusions (To Avoid Scope Creep)

- No user accounts, comments, or social features – static publishing only.
- No advanced search or hover tooltips on links (future if needed; start with basic markdown links).
- No separate .pl domain (duplication hell; handle PL via /pl prefix in one app).
- No auto-translation tools; manual (Szymon drafts in PL, translates to EN via AI help outside the app).
- No analytics beyond basic (e.g., open metrics page updates manually).

## 2. Features

Organized by domain for clarity. All dynamic content (entries) sourced from root /content/{lang}/diary/_.md files (e.g., /content/en/diary/202507121430-netlify-deploy.md – prefixed with YYYYMMDDHHMM for chronological order). Use frontmatter tags (e.g., tags: ['personal', 'howtos', 'glossary']) for filtering. Static pages in root /content/en/static/_.md or hard-coded if simple.

### 2.1 Diary Domain (/diary)

- **Purpose**: All content as honest diary entries – personal reflections (therapy sessions on burnout, success failures like hiring after 165M views, prokrastynacja escape, work-from-car excitement) mixed with tech resources (howtos like deploy to Netlify step-by-step, glossary explanations with everyday analogies and personal journeys, e.g., "It used to melt my brain looking at package.json; now it's funny how simple it feels").
- **Features**:
  - List page: /diary – Displays chronological list of all entries (titles, dates, excerpts, tags). Intro header: "My Diary: Uncensored snapshots of my life and rebuild, inspired by Rick Rubin." Include tag-based filters/sections (e.g., Personal, Howtos for non-tech founders, Glossary).
  - Single entry: /diary/{slug} – Full markdown-rendered content, display tags.
  - Bilingual: Default EN (/diary/slug), PL via /pl/diary/slug (same slug, different content file).
  - Cross-linking: Natural in markdown (e.g., [package.json](/diary/glossary-package-json)); no hover effects (basic <a> tags).
- **Content Rules**: Everything "honest of where I am at" – pure personal/philosophical for 'personal' tags; concise howtos/glossary with relatable analogies and personal anecdotes (e.g., sharing the learning curve, past struggles, current insights). Inspired by Paul Graham essays but grounded in Rubin's no-blocks mindset.

### 2.2 Static Pages

- **/uses**: List of tools/gear (software: AI for code, hardware: car setup). Hard-coded or md-sourced; bilingual via /pl/uses.
- **/now**: Current status (what I'm building, feeling). Manual update; bilingual via /pl/now.
- **/about** (or /story): Bio + journey (success to rock bottom – TV, millions views, debts/depression). Single page with sections; bilingual via /pl/about. If split, /about for bio, /story for narrative – but minimize routes.
- **/open**: Transparent metrics (debts list, future incomes, shocking facts). Table or list format; update manually; bilingual via /pl/open.

### 2.3 Global Features

- **Navigation**: Simple header/footer with links to /diary (and tag filters, e.g., /diary?tag=howtos if needed) and static pages.
- **Bilingual Handling**:
  - **English (Primary)**: Main version with clean URLs (no language prefix)
    - Examples: `/diary/slug`, `/about`, `/uses`, `/now`, `/open`
  - **Polish (Secondary)**: Accessible via `/pl` prefix for all routes
    - Examples: `/pl/diary/slug`, `/pl/about`, `/pl/uses`, `/pl/now`, `/pl/open`
  - Content files organized separately: `/content/en/...` and `/content/pl/...`
  - No auto-language detection; users choose via language toggle in navigation
  - Polish serves as drafting language (easier for Szymon), then manual translation to English
- **Markdown Support**: All content in .md with frontmatter (title, date, tags) for parsing/filtering.
- **Deployment**: Host on Vercel/Netlify; free tier sufficient.

## 3. Non-Functional Requirements

- **Performance**: Static generation where possible; fast loads (minimal JS).
- **Accessibility**: Basic (alt texts, semantic HTML); no overkill.
- **Security**: No user data; static site – low risk.
- **Tech Stack**: SvelteKit for simplicity (runes for reactivity, mdsvex for markdown); but focus on logic, not code details here.
- **Content Management**: All in root /content – editable via git/text editor; no CMS (lean for solo dev).
- **Scalability**: Handle low traffic; no future-proofing beyond basics.
- **Testing**: Manual; publish first entry to verify.
- **ADHD-Friendly**: Low-cognitive-load – consistent naming (slugs from filenames), proximity (content in root), no deep folders.

## 4. Success Metrics

- Publish first diary entry (with tags, honest tech anecdote) within days.
- Easy updates: Add .md with tags, git push, deploy.
- Reader feedback on transparency (qualitative).

## 5. Risks and Mitigations

- Risk: Content duplication in langs – Mitigation: Write EN first (main), translate to PL manually.
- Risk: Scope creep (e.g., complex tag filters) – Mitigation: Stick to MVP; add only after core works.
- Risk: Loss of separation without tags – Mitigation: Clear tag-based filters in nav/list page.

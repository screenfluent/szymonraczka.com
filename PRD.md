# Product Requirements Document (PRD) for Szymon Raczka's Personal Site

## 1. Introduction
### Overview
This PRD defines the requirements for a personal website at szymonraczka.com, serving as a transparent platform for Szymon Raczka (44, solo indie dev with ADHD) to share raw personal reflections, technical insights for non-tech founders, and open metrics. The site emphasizes authenticity: uncensored diary-style entries on life, burnout, and success; practical tech explanations with everyday analogies (e.g., package.json as a kitchen shopping list); and static pages for current status and tools.

The site is built with simplicity in mind: minimal features to solve today's problems (publishing content without cognitive overload), organized by functional domain (personal vs. tech), and bilingual (English as default, Polish as secondary).

### Goals
- Enable quick publishing of personal "diary entries" (raw thoughts on burnout, 165M YouTube views success, work-from-car happiness, perfectionism escape) to build audience resonance through transparency.
- Provide tech resources tailored for non-tech founders via /tech sub-sections: /howtos for step-by-step guides, /glossary for term explanations using simple analogies from daily life (kitchen, work, restaurant) to demystify concepts like codebases or AI-assisted building.
- Share open metrics (debts, future incomes, shocking personal facts) to inspire and shock "average Joes."
- Support bilingual content: English as primary (no prefix in URLs), Polish as optional (with top-level /pl prefix, e.g., /pl/diary, /pl/tech/howtos) for easier initial drafting in Polish then translation.
- Minimize maintenance: single source of truth for content in root-level /content folder (subdivided by language and domain), no deep nesting, predictable naming.

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
Organized by domain for clarity. All dynamic content (posts) sourced from root /content/{lang}/{domain}/*.md files (e.g., /content/en/tech/howtos/202507121430-netlify-deploy.md – prefixed with YYYYMMDDHHMM for chronological order). Static pages in root /content/en/static/*.md or hard-coded if simple.

### 2.1 Personal Domain (/diary)
- **Purpose**: Raw, uncensored personal reflections like therapy sessions – burnout stories, success failures (hiring after 165M views leading to crash), prokrastynacja escape, work-from-car excitement.
- **Features**:
  - List page: /diary – Displays chronological list of entries (titles, dates, excerpts). Intro header: "My Diary: Uncensored snapshots of my life and rebuild."
  - Single entry: /diary/{slug} – Full markdown-rendered content.
  - Bilingual: Default EN (/diary/slug), PL via /pl/diary/slug (same slug, different content file).
- **Content Rules**: No tech jargon here; pure personal/philosophical (inspired by Paul Graham essays on life/startups).

### 2.2 Tech Domain (/tech)
- **Purpose**: Practical resources for non-tech founders – subdivided into /howtos (tutorials like deploy to Netlify step-by-step) and /glossary (explanations with everyday analogies, e.g., codebase as a restaurant kitchen workflow).
- **Features**:
  - Main list page: /tech – Overview or combined list of howtos and glossary (titles, dates, tagged).
  - Howtos sub-section: /tech/howtos – Chronological list of guides; single: /tech/howtos/{slug} (e.g., /tech/howtos/step-by-step-guide-to-deploy-your-blog-on-netlify) – Markdown content with internal links to glossary terms.
  - Glossary sub-section: /tech/glossary – Filtered list of terms; single: /tech/glossary/{term-slug} (e.g., /tech/glossary/package-json) – Explanation with analogies.
  - Cross-linking: Terms linkable from howtos (e.g., [package.json](/tech/glossary/package-json)); no hover effects (basic <a> tags).
  - Bilingual: Default EN, PL via /pl/tech/howtos/{slug} or /pl/tech/glossary/{term-slug}.
- **Content Rules**: Keep howtos concise; avoid digressions by linking to glossary for details. All in "for non-tech founders" vibe – simple, relatable analogies.

### 2.3 Static Pages
- **/uses**: List of tools/gear (software: AI for code, hardware: car setup). Hard-coded or md-sourced; bilingual via /pl/uses.
- **/now**: Current status (what I'm building, feeling). Manual update; bilingual via /pl/now.
- **/about** (or /story): Bio + journey (success to rock bottom – TV, millions views, debts/depression). Single page with sections; bilingual via /pl/about. If split, /about for bio, /story for narrative – but minimize routes.
- **/open**: Transparent metrics (debts list, future incomes, shocking facts). Table or list format; update manually; bilingual via /pl/open.

### 2.4 Global Features
- **Navigation**: Simple header/footer with links to all sections (including /tech/howtos and /tech/glossary).
- **Bilingual Handling**: EN default (no prefix); PL via top-level /pl prefix (e.g., /pl/tech/howtos/slug). Content files separate per lang in /content/{lang}/... – no auto-switch; user chooses via links (e.g., lang toggle in nav).
- **Markdown Support**: All content in .md with frontmatter (title, date, tags) for parsing.
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
- Publish first /diary entry and /tech/howtos guide within days.
- Easy updates: Add .md, git push, deploy.
- Reader feedback on transparency (qualitative).

## 5. Risks and Mitigations
- Risk: Content duplication in langs – Mitigation: Write EN first (main), translate to PL manually.
- Risk: Scope creep (e.g., deep nesting in /tech) – Mitigation: Stick to MVP; add only after core works.
- Risk: Confusion between personal/tech – Mitigation: Clear domain separation in nav/URLs.
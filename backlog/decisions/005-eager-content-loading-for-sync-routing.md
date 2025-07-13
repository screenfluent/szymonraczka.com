# ADR-005: Switch to Eager Content Loading for Synchronous Parameter Matching

**Date:** 2025-01-13  
**Status:** Accepted

## Context

SvelteKit parameter matchers must be synchronous functions that return `boolean`, but our content loading system was asynchronous. This caused a TypeScript error: `Type 'Promise<boolean>' is not assignable to type 'boolean'` in `src/params/tag.ts`.

The async content loading pattern worked for page load functions but violated SvelteKit's requirement that route matching (including parameter validation) happens instantly during navigation.

## Decision

Refactor content loading from lazy (`eager: false`) to eager (`eager: true`) using Vite's `import.meta.glob`, making all content functions synchronous.

**Technical Implementation:**
- Switch `import.meta.glob('/content/*/diary/*.md', { eager: true })` for synchronous module loading
- Process all markdown files at module initialization time into `allEntries` array
- Convert all content functions to synchronous: `loadDiaryEntries()`, `getAllTags()`, `getDiaryEntriesByTag()`, `loadDiaryEntry()`
- Remove `async`/`await` from parameter matcher and load functions
- Update Svelte components to use data directly instead of `{#await}` blocks

## Consequences

**Positive:**
- **Fixed TypeScript Error:** Parameter matchers now properly synchronous as required
- **Better Performance:** Content processed once at build/startup instead of per-request
- **Faster Navigation:** Route matching happens instantly without async operations
- **Simpler Code:** Eliminated async/await complexity in routing logic
- **SSR Compatible:** All content available immediately during server rendering
- **Build Optimization:** Vite can better optimize bundling with eager imports

**Negative:**
- **Memory Usage:** All content loaded into memory at startup (acceptable for blog scale)
- **Startup Time:** Slightly longer initial load time (negligible for static content)
- **Less Lazy:** Can't defer loading unused content (not needed for our use case)

**Performance Impact:**
- Before: Route match → Async content load → Validation → Navigation
- After: Route match → Instant validation → Navigation

**Files Modified:**
- Updated: `src/lib/content.ts` (eager loading, synchronous functions)
- Updated: `src/params/tag.ts` (removed async from matcher)
- Updated: `src/routes/diary/+page.ts` (already synchronous)
- Updated: `src/routes/diary/[slug]/+page.ts` (removed async)
- Updated: `src/routes/[tag=tag]/+page.ts` (removed async)
- Updated: `src/routes/diary/+page.svelte` (removed {#await} blocks)
- Updated: `src/routes/[tag=tag]/+page.svelte` (removed {#await} blocks)

## Alternative Considered

**Pre-computed Tag List:** We considered maintaining a separate static list of valid tags, but this would create a maintenance burden and potential sync issues. The eager loading approach keeps the single source of truth while solving the synchronous requirement.
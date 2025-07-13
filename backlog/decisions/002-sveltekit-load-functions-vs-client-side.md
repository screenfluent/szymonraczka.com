# ADR-002: Use SvelteKit Load Functions Instead of Client-Side Data Fetching

**Date:** 2025-07-12  
**Status:** Accepted

## Context

Initial task-2 implementation used client-side data fetching with Svelte 5 runes (`$state`, `$effect`) to load diary content. Code review revealed this was an anti-pattern in SvelteKit that prevented SSR and hurt performance.

## Decision

Refactor to use SvelteKit's proper data loading patterns with `+page.ts` files and load functions.

**Technical Implementation:**
- Create `+page.ts` files for data loading logic
- Move `loadDiaryEntries()` and `loadDiaryEntry()` calls to load functions
- Use `PageProps` type for clean TypeScript integration
- Implement proper 404 handling with SvelteKit's `error()` helper

## Consequences

**Positive:**
- **SSR Enabled:** Content rendered on server for better SEO and initial page load
- **Performance:** Eliminated client-side fetch waterfalls 
- **Streaming:** SvelteKit can stream data while rendering page shell
- **Error Handling:** Proper 404 pages via SvelteKit error boundaries
- **Type Safety:** Better TypeScript integration with generated types
- **Best Practices:** Follows SvelteKit architectural patterns

**Negative:**
- Required refactoring existing components
- More files to maintain (+page.ts alongside +page.svelte)
- Initial learning curve for proper SvelteKit patterns

**Performance Impact:**
- Before: Page loads → JS loads → Data fetches → Content renders
- After: Page + Data load in parallel → Content renders immediately

**Files Modified:**
- Added: src/routes/diary/+page.ts
- Added: src/routes/diary/[slug]/+page.ts  
- Updated: src/routes/diary/+page.svelte (simplified, removed client-side logic)
- Updated: src/routes/diary/[slug]/+page.svelte (simplified, removed client-side logic)
# ADR-003: Adopt Svelte 5 Modern Patterns Over Legacy Approaches

**Date:** 2025-07-12  
**Status:** Accepted

## Context

Project uses Svelte 5 with runes mode, but initial implementation mixed modern and legacy patterns. Code review identified opportunities to use current best practices for better maintainability and performance.

## Decision

Consistently use modern Svelte 5 patterns throughout the codebase.

**Specific Pattern Adoptions:**

1. **Component Types:** `Component` instead of deprecated `ComponentType`
2. **Component Rendering:** `{@const TheComponent = entry.Component} <TheComponent />` instead of `<svelte:component>`
3. **Props Typing:** `PageProps` instead of manual `{ data: PageData }`
4. **Comment Style:** `//` single-line comments over `/* */` block comments

## Consequences

**Positive:**
- **Future-Proof:** Aligned with Svelte 5 direction, avoiding deprecation warnings
- **Performance:** Modern patterns are optimized for Svelte 5's reactivity system
- **Developer Experience:** Better TypeScript integration and IDE support
- **Maintainability:** Consistent patterns across codebase
- **Documentation:** Cleaner, more readable code style

**Negative:**
- Required learning curve for Svelte 5 specifics
- Some refactoring needed to update legacy patterns

**Pattern Changes Made:**
```typescript
// Before (legacy)
import type { ComponentType } from 'svelte';
Component: ComponentType;
<svelte:component this={entry.Component} />
let { data }: { data: PageData } = $props();

// After (modern)
import type { Component } from 'svelte';
Component: Component;
{@const TheComponent = entry.Component} <TheComponent />
let { data }: PageProps = $props();
```

**Files Updated:**
- src/lib/content.ts (Component type)
- src/routes/diary/+page.svelte (PageProps)
- src/routes/diary/[slug]/+page.svelte (PageProps, component rendering)
- CLAUDE.md (documented preferred patterns)
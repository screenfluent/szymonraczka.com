# ADR-001: Use MDsveX Native Parsing Instead of gray-matter

**Date:** 2025-07-12  
**Status:** Accepted

## Context

During task-1 implementation, we initially used gray-matter library for frontmatter parsing in markdown files. However, we discovered that MDsveX (already in use for SvelteKit) has built-in YAML frontmatter parsing capabilities.

## Decision

Remove gray-matter dependency and use MDsveX's native frontmatter parsing.

**Technical Implementation:**
- Use `metadata` property from MDsveX module imports
- Configure MDsveX to handle `.md` extensions in svelte.config.js
- Access frontmatter via `module.metadata` instead of gray-matter parsing

## Consequences

**Positive:**
- Reduced bundle size by removing external dependency
- Better integration with SvelteKit/MDsveX ecosystem
- Cleaner code with fewer parsing steps
- Direct access to Svelte components via `module.default`

**Negative:**
- Initial refactoring required to change parsing approach
- Less flexibility compared to gray-matter's extensive options (not needed for our use case)

**Files Modified:**
- Removed: gray-matter from package.json
- Updated: src/lib/content.ts (simplified parsing logic)
- Updated: svelte.config.js (MDsveX .md extension support)
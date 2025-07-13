# ADR-004: Vite Configuration for Robust Content File Serving

**Date:** 2025-07-12  
**Status:** Accepted

## Context

Content loading system uses `import.meta.glob('/content/*/diary/*.md')` to dynamically import markdown files. This requires specific Vite configuration to work reliably in both development and production environments.

## Decision

Configure Vite with explicit path resolution and file serving permissions for content files.

**Configuration Implemented:**
```typescript
// vite.config.ts
export default defineConfig({
  resolve: {
    alias: {
      '/content': resolve(process.cwd(), 'content')  // Explicit path resolution
    }
  },
  server: {
    fs: {
      allow: ['.']  // Allow serving files from project root
    }
  }
});
```

## Consequences

**Positive:**
- **Development Reliability:** Content files load consistently in dev server
- **Path Clarity:** Explicit alias removes ambiguity in content path resolution
- **Security Conscious:** Controlled file system access (project root only)
- **Cross-Platform:** Works reliably across different operating systems
- **Future-Proof:** Handles content structure changes gracefully

**Negative:**
- Additional configuration complexity
- Need to understand Vite file serving mechanics

**Technical Details:**
- **Path Alias:** Maps `/content` to absolute project path for reliable imports
- **File System Access:** `allow: ['.']` permits serving from project root (required for glob imports)
- **Node.js Types:** Added `@types/node` for `process.cwd()` and `resolve()` support

**Why This Was Needed:**
- Glob imports like `/content/en/diary/*.md` require Vite to resolve paths outside standard `/src`
- Development server needs explicit permission to serve content directory files
- Production builds work differently than dev server, requiring unified configuration

**Files Modified:**
- vite.config.ts (added resolve alias and server configuration)
- package.json (added @types/node dependency)
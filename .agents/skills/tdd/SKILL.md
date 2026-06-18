---
name: tdd
description: Test-driven development. Use when building or fixing meaningful behavior test-first, using red-green-refactor and vertical slices.
---


# Test-Driven Development

Tests verify behavior through public interfaces, not implementation details.

## Loop

1. Pick one behavior.
2. Write one test.
3. Run it and confirm RED for the expected reason.
4. Write minimal code for GREEN.
5. Check that the test protects behavior, not implementation.
6. Refactor only after GREEN and only for a concrete reason.

## Avoid

- writing all tests first
- mocking internal collaborators
- testing private helpers directly
- asserting internal call order
- test theater for static text or simple HTML/CSS

Use `codebase-design` when the public interface is unclear.

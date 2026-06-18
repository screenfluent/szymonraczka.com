# Tests

Good tests verify behavior through public interfaces.

Avoid:
- testing private helpers
- mocking internal collaborators
- asserting call order
- tests that fail on harmless refactors

Prefer:
- HTTP tests for routing/handlers
- public functions for pure logic
- real SQLite test DBs where practical

# Mocking

Mock system boundaries only:
- external APIs
- time/randomness
- filesystem when needed
- network when needed

Do not mock code you own.
Prefer real collaborators and behavior tests.

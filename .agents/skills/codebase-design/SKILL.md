---
name: codebase-design
description: Shared design discipline for deep modules, clean seams, and testable public interfaces. Use when evaluating structure, interfaces, packages, or refactors.
---


# Codebase Design

Prefer deep modules: small interface, useful behavior hidden behind it.

## Check before adding structure

Ask:
- Is this seam real?
- Can a plain function be enough?
- Does this interface have more than one implementation or a system boundary?
- Will this survive a behavior-preserving refactor?
- Does this protect URLs, auth, data, security, performance, or simplicity?

## Reject

- shallow modules
- interfaces without a real seam
- service/repository/factory/manager layers by default
- package splits that only mirror enterprise architecture
- refactors that do not remove a concrete risk or complexity

## Prefer

- simple packages
- explicit data flow
- public behavior tests
- boring code that is easy to delete

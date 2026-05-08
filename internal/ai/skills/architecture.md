# Architecture Reminder

Use `AGENTS.md` for rules and `docs/project-map.md` for current package boundaries.

Architectural checks:

- primitives stay in `ability`
- structural creature rules stay in `creaturetype`
- cross-domain character composition stays in `character`
- seed/query domains do not become character aggregates
- no speculative abstractions

If a detailed architecture review is needed, read the directly relevant package files after `docs/project-map.md`.

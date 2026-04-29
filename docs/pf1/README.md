# Pathfinder 1 Rules Lookup

`docs/pf1` stores local Pathfinder 1 rule sources for rule-sensitive implementation work.

The lookup flow is:

```text
PDF files -> extracted text files -> JSONL chunks -> local rg search
```

## Directory Structure

- `raw/`: source PDFs
- `text/`: extracted text files
- `chunks/`: generated JSONL chunks

## Extract PDFs

From the project root:

```bash
./docs/pf1/extract_rules.sh
```

From `docs/pf1`:

```bash
./extract_rules.sh
```

## Generate Chunks

From the project root:

```bash
go run docs/pf1/chunk_rules.go
```

From `docs/pf1`:

```bash
go run chunk_rules.go
```

## Search Rules

Use `rg` against the generated chunks:

```bash
rg -i "ability damage" docs/pf1/chunks
rg -i "ability drain" docs/pf1/chunks
rg -i "attack of opportunity" docs/pf1/chunks
rg -i "combat maneuver" docs/pf1/chunks
rg -i "circumstance bonus" docs/pf1/chunks
```

## Rule-Sensitive Implementation Policy

- Search local PF1 chunks before implementing Pathfinder-specific behavior.
- Prefer local extracted rules over memory.
- Mention source PDF, page, or chunk when relevant.
- If a rule cannot be found locally, stop and report that it was not found.
- Do not make `internal/domain` depend on PDFs, extracted text, chunks, generated indexes, or search tooling.

# AGENTS.md — webtyp/msgtype

Working notes for AI agents operating in this repository. End-user docs: [README.md](README.md).

## What this repo is

The classification of a message for logs and UI — `msgtype.Info`, `msgtype.Error`,
`msgtype.Success`, … — and `msgtype.Detect(text)`, which guesses the type of a log line. It was split
out of `webtyp.com/fmt`: classifying messages is not formatting.

## This package compiles to WASM

- Do NOT import `strings`, `strconv`, `errors` or stdlib `fmt`; use `webtyp.com/fmt` exported helpers
  (`fmt.ToLower`, `fmt.Contains`).
- No `map`, no `reflect`.

## Wire compatibility

The numeric values of the constants (`Normal = 0` … `Response = 13`) travel over SSE and are stored
by consumers. **Never reorder or renumber them.** A new type is appended at the end.

## The build that defines "done"

```bash
go install webtyp.com/devflow/cmd/gotest@latest   # once
gotest
```

## Rules

- Tests live in `tests/` as `package msgtype_test` (public API only). A root-level test is allowed only
  with a top-of-file justification of the unexported identifier it needs. **Never export a symbol so
  a test can reach it.**

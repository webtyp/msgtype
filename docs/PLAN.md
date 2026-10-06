---
PLAN: "feat: msgtype — message classification moved out of webtyp.com/fmt"
TAG: v0.1.0
EXECUTOR: jules
REVIEWER: none
STATUS: running
SESSION: 14815669336747803909
---

> This plan is dispatched via the CodeJob workflow. See skill: agents-workflow.

# Plan — msgtype: message classification as its own module

Phase **A4 (gate)** of the master plan `SOURCE_SELECTION_MASTER_PLAN.md` (orchestration only —
everything this plan needs is inline). Consumers migrate after this tag exists.

Read [AGENTS.md](../AGENTS.md) first (WASM rules; **constant values are wire format**; tests in
`tests/`).

## Why

`fmt.MessageType`, `fmt.Msg.*` and `Convert(s).StringType()` classify log/UI messages. That is a
separate responsibility from formatting, and it grows: SSE added `Connect`, `Auth`, `Parse`,
`Timeout` and `Broadcast`; pub/sub added `Event`, `Request` and `Response`. Every addition used to be
a new `fmt` release for the whole ecosystem.

## State when this plan starts (moved by the maintainer, 2026-10-06)

- `msgtype.go` = `webtyp/fmt/messagetype.go` **unchanged**. It still says `package fmt`. It uses
  `Conv` internals (`GetString`, `putConv`, `swapBuff`, `changeCase`, `bufferContainsPattern`) that
  do not exist outside `fmt`.
- `tests/msgtype_test.go` = `webtyp/fmt/messagetype_test.go`, unchanged.
- `docs/MESSAGE_TYPES.md`, `docs/ISSUE_MESSAGE_TYPE.md`, `docs/FEAT_MESSAGETYPE_SSE.md`.
- The gonew stub was deleted. Do not recreate it.

## Design gate

1. **Prior art.** `log/slog.Level` (a typed int with named constants: `slog.LevelInfo`),
   `logrus.Level`, `zap/zapcore.Level`. All are their own package-level typed constants, never a
   struct-of-constants, and none lives in the formatting package.
2. **Novice-name test.** `msgtype.Type`, `msgtype.Error`, `msgtype.Detect(line)` — "detect the
   message type of the line".

   | Old (`webtyp.com/fmt`) | New (`webtyp.com/msgtype`) |
   |---|---|
   | `fmt.MessageType` | `msgtype.Type` |
   | `fmt.Msg.Info` (etc.) | `msgtype.Info` (etc.) |
   | `s, t := fmt.Convert(x).StringType()` | `s, t := x, msgtype.Detect(x)` |
   | `t.IsError()` (etc.) | `t == msgtype.Error` |
   | `t.IsNetworkError()` | deleted (zero users outside fmt's own tests, verified 2026-10-06) |
3. **Complexity ledger.**
   ```
   Concepts the developer must learn   +0 / −15 (Msg struct and 14 Is* methods)
   Files they must touch to do X       +0 / −0
   Lines at the call site              ±0
   Ways to do the same thing           −14 (t.IsX() ≡ t == msgtype.X)
   ```
4. **Where it belongs.** Its own module, like `slog.Level` is not in `fmt`.
5. **What it deletes.** The `Msg` struct variable, all `Is*` methods, `StringType`, and
   `detectMessageTypeFromBuffer`.

## Stage 1 — the API (`msgtype.go`, rewrite)

```go
package msgtype

// Type classifies a message for logs and UI. The numeric values travel over
// SSE: never reorder them; append new types at the end.
type Type uint8

const (
	Normal    Type = iota // 0
	Info                  // 1
	Error                 // 2
	Warning               // 3
	Success               // 4
	Connect               // 5 connection error
	Auth                  // 6 authentication error
	Parse                 // 7 parse/decode error
	Timeout               // 8 timeout error
	Broadcast             // 9 broadcast/send error
	Debug                 // 10
	Event                 // 11 pub/sub event
	Request               // 12
	Response              // 13
)

// String returns the type's name ("Info", "Error", …; "Normal" for unknown values).
func (t Type) String() string

// Detect guesses the type of a log line from its words, case-insensitively.
// Checked in this order, first hit wins: Error, Warning, Success, Info, Debug;
// otherwise Normal.
func Detect(text string) Type
```

- The values must equal the old `fmt.Msg` values exactly (`{0, 1, …, 13}` in that field order).
- `String()`: the same switch as the old method.
- `Detect`: `lower := fmt.ToLower(text)`, then `fmt.Contains(lower, p)` over these pattern
  slices, in this order and with exactly these words (from the old code):
  - error: `error`, `failed`, `exit status 1`, `undeclared`, `undefined`, `fatal`
  - warning: `warning`, `warn`
  - success: `success`, `completed`, `successful`, `done`
  - info: `info`, `starting`, `initializing`
  - debug: `debug`

  Keep them as package-level `[]string` variables. Import `webtyp.com/fmt`
  (`go get webtyp.com/fmt@latest`).

## Stage 2 — tests (`tests/msgtype_test.go`)

`package msgtype_test`. Convert mechanically:
- `Convert(x).StringType()` → `msgtype.Detect(x)`; assertions on the returned string become
  assertions on `x` or are dropped.
- `Msg.X` → `msgtype.X`; `t.IsX()` → `t == msgtype.X`.

Every existing detection case must keep its expected type. Add `TestWireValues`, asserting
`Normal == 0`, `Info == 1`, …, `Response == 13` one by one.

## Stage 3 — docs

`README.md`: the old→new table above, the wire-compatibility rule, and one example. Merge the
still-relevant parts of the three `docs/*.md` into `README.md` and delete the issue/feature notes
whose work is done.

## Acceptance

- `gotest` passes (includes WASM).
- `ls *_test.go 2>/dev/null` → nothing at the root.
- `grep -rn "Msg\.\|MessageType\|StringType\|func (t Type) Is" --include='*.go' .` → empty.

## Stages

| # | Stage | Files |
|---|---|---|
| 1 | API | `msgtype.go`, `go.mod`, `go.sum` |
| 2 | Tests | `tests/msgtype_test.go` |
| 3 | Docs | `README.md`, `docs/*.md` |

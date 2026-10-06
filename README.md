# msgtype
<img src="docs/img/badges.svg">

Message classification (info, error, success, warning...) for logs and UI. This package is specifically for classifying messages and forms its own module separate from formatting.

## Why
`fmt.MessageType` was split into its own module `webtyp.com/msgtype` because classifying log/UI messages is a separate responsibility from formatting, and it often needs extending (e.g., SSE network errors, Pub/Sub events) without forcing a new release of `webtyp.com/fmt` for the whole ecosystem.

## Wire compatibility
The numeric values of the constants (`Normal = 0` … `Response = 13`) travel over SSE and are stored by consumers. **Never reorder or renumber them.** A new type is appended at the end.

## Migration Guide

| Old (`webtyp.com/fmt`) | New (`webtyp.com/msgtype`) |
|---|---|
| `fmt.MessageType` | `msgtype.Type` |
| `fmt.Msg.Info` (etc.) | `msgtype.Info` (etc.) |
| `s, t := fmt.Convert(x).StringType()` | `s, t := x, msgtype.Detect(x)` |
| `t.IsError()` (etc.) | `t == msgtype.Error` |
| `t.IsNetworkError()` | deleted |

## Usage Example

```go
package main

import (
    "log"
    "webtyp.com/msgtype"
)

func handleError(message string) {
    log.Println("Handling error:", message)
}

func logMessage(message string, msgType msgtype.Type) {
    log.Printf("[%s] %s\n", msgType.String(), message)
}

func progressCallback(message string) {
    msgType := msgtype.Detect(message)
    if msgType == msgtype.Error {
        handleError(message)
    } else {
        logMessage(message, msgType)
    }
}
```

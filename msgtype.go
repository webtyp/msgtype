package msgtype

import "webtyp.com/fmt"

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
func (t Type) String() string {
	switch t {
	case Info:
		return "Info"
	case Error:
		return "Error"
	case Warning:
		return "Warning"
	case Success:
		return "Success"
	case Connect:
		return "Connect"
	case Auth:
		return "Auth"
	case Parse:
		return "Parse"
	case Timeout:
		return "Timeout"
	case Broadcast:
		return "Broadcast"
	case Debug:
		return "Debug"
	case Event:
		return "Event"
	case Request:
		return "Request"
	case Response:
		return "Response"
	default:
		return "Normal"
	}
}

var (
	errorPatterns   = []string{"error", "failed", "exit status 1", "undeclared", "undefined", "fatal"}
	warningPatterns = []string{"warning", "warn"}
	successPatterns = []string{"success", "completed", "successful", "done"}
	infoPatterns    = []string{"info", "starting", "initializing"}
	debugPatterns   = []string{"debug"}
)

// Detect guesses the type of a log line from its words, case-insensitively.
// Checked in this order, first hit wins: Error, Warning, Success, Info, Debug;
// otherwise Normal.
func Detect(text string) Type {
	lower := fmt.ToLower(text)
	for _, p := range errorPatterns {
		if fmt.Contains(lower, p) {
			return Error
		}
	}
	for _, p := range warningPatterns {
		if fmt.Contains(lower, p) {
			return Warning
		}
	}
	for _, p := range successPatterns {
		if fmt.Contains(lower, p) {
			return Success
		}
	}
	for _, p := range infoPatterns {
		if fmt.Contains(lower, p) {
			return Info
		}
	}
	for _, p := range debugPatterns {
		if fmt.Contains(lower, p) {
			return Debug
		}
	}
	return Normal
}

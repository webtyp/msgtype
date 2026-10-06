package msgtype_test

import (
	"testing"
	"webtyp.com/msgtype"
)

func TestTypeDetection(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		msgType := msgtype.Detect("")
		if msgType != msgtype.Normal {
			t.Errorf("Expected Normal for empty string, got %v", msgType)
		}
	})

	t.Run("Error keywords", func(t *testing.T) {
		errorKeywords := []string{
			"This is an error message",
			"Operation failed",
			"exit status 1",
			"variable undeclared",
			"function undefined",
			"fatal exception",
		}
		for _, keyword := range errorKeywords {
			msgType := msgtype.Detect(keyword)
			if msgType != msgtype.Error {
				t.Errorf("Expected Error for keyword %q, got %v", keyword, msgType)
			}
		}
	})

	t.Run("Success keywords", func(t *testing.T) {
		successKeywords := []string{
			"Success! Operation completed",
			"success",
			"Operation completed",
			"Build successful",
			"Task done",
		}
		for _, keyword := range successKeywords {
			msgType := msgtype.Detect(keyword)
			if msgType != msgtype.Success {
				t.Errorf("Expected Success for keyword %q, got %v", keyword, msgType)
			}
		}
	})

	t.Run("Info keywords", func(t *testing.T) {
		infoKeywords := []string{
			"Info: Starting process",
			"... initializing ...",
			"starting up",
			"initializing system",
		}
		for _, keyword := range infoKeywords {
			msgType := msgtype.Detect(keyword)
			if msgType != msgtype.Info {
				t.Errorf("Expected Info for keyword %q, got %v", keyword, msgType)
			}
		}
	})

	t.Run("Warning keywords", func(t *testing.T) {
		warningKeywords := []string{
			"Warning: disk space low",
			"warn user",
		}
		for _, keyword := range warningKeywords {
			msgType := msgtype.Detect(keyword)
			if msgType != msgtype.Warning {
				t.Errorf("Expected Warning for keyword %q, got %v", keyword, msgType)
			}
		}
	})

	t.Run("Debug keywords", func(t *testing.T) {
		debugKeywords := []string{
			"debug: something happening",
			"[debug] status",
			"DEBUG: uppercase",
			"DeBuG: Mixed Case",
		}
		for _, keyword := range debugKeywords {
			msgType := msgtype.Detect(keyword)
			if msgType != msgtype.Debug {
				t.Errorf("Expected Debug for keyword %q, got %v", keyword, msgType)
			}
		}
	})

	t.Run("Normal message", func(t *testing.T) {
		msgType := msgtype.Detect("Hello world")
		if msgType != msgtype.Normal {
			t.Errorf("Expected Normal for generic message, got %v", msgType)
		}
	})
}

func TestSSERelatedTypes(t *testing.T) {
	tests := []struct {
		name     string
		msgType  msgtype.Type
		expected string
	}{
		{"Connect", msgtype.Connect, "Connect"},
		{"Auth", msgtype.Auth, "Auth"},
		{"Parse", msgtype.Parse, "Parse"},
		{"Timeout", msgtype.Timeout, "Timeout"},
		{"Broadcast", msgtype.Broadcast, "Broadcast"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.msgType.String() != tt.expected {
				t.Errorf("Expected String() for %s to be %q, got %q", tt.name, tt.expected, tt.msgType.String())
			}
		})
	}
}

func TestPubSubTypes(t *testing.T) {
	tests := []struct {
		name     string
		msgType  msgtype.Type
		expected string
	}{
		{"Event", msgtype.Event, "Event"},
		{"Request", msgtype.Request, "Request"},
		{"Response", msgtype.Response, "Response"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.msgType.String() != tt.expected {
				t.Errorf("Expected String() for %s to be %q, got %q", tt.name, tt.expected, tt.msgType.String())
			}
		})
	}
}

func TestWireValues(t *testing.T) {
	if msgtype.Normal != 0 {
		t.Errorf("Expected msgtype.Normal == 0, got %d", msgtype.Normal)
	}
	if msgtype.Info != 1 {
		t.Errorf("Expected msgtype.Info == 1, got %d", msgtype.Info)
	}
	if msgtype.Error != 2 {
		t.Errorf("Expected msgtype.Error == 2, got %d", msgtype.Error)
	}
	if msgtype.Warning != 3 {
		t.Errorf("Expected msgtype.Warning == 3, got %d", msgtype.Warning)
	}
	if msgtype.Success != 4 {
		t.Errorf("Expected msgtype.Success == 4, got %d", msgtype.Success)
	}
	if msgtype.Connect != 5 {
		t.Errorf("Expected msgtype.Connect == 5, got %d", msgtype.Connect)
	}
	if msgtype.Auth != 6 {
		t.Errorf("Expected msgtype.Auth == 6, got %d", msgtype.Auth)
	}
	if msgtype.Parse != 7 {
		t.Errorf("Expected msgtype.Parse == 7, got %d", msgtype.Parse)
	}
	if msgtype.Timeout != 8 {
		t.Errorf("Expected msgtype.Timeout == 8, got %d", msgtype.Timeout)
	}
	if msgtype.Broadcast != 9 {
		t.Errorf("Expected msgtype.Broadcast == 9, got %d", msgtype.Broadcast)
	}
	if msgtype.Debug != 10 {
		t.Errorf("Expected msgtype.Debug == 10, got %d", msgtype.Debug)
	}
	if msgtype.Event != 11 {
		t.Errorf("Expected msgtype.Event == 11, got %d", msgtype.Event)
	}
	if msgtype.Request != 12 {
		t.Errorf("Expected msgtype.Request == 12, got %d", msgtype.Request)
	}
	if msgtype.Response != 13 {
		t.Errorf("Expected msgtype.Response == 13, got %d", msgtype.Response)
	}
}

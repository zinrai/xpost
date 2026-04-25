package main

import (
	"strings"
	"testing"
)

func TestTrimBehavior(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantText  string
		wantEmpty bool
	}{
		{name: "normal text", input: "hello", wantText: "hello"},
		{name: "trailing newline removed", input: "hello\n", wantText: "hello"},
		{name: "multiple trailing newlines removed", input: "hello\n\n\n", wantText: "hello"},
		{name: "internal newline preserved", input: "line1\nline2\n", wantText: "line1\nline2"},
		{name: "trailing mixed whitespace", input: "hello  \t\r\n  ", wantText: "hello"},
		{name: "empty string", input: "", wantEmpty: true},
		{name: "whitespace only", input: "   \n\t\r  ", wantEmpty: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.TrimRight(tt.input, " \t\r\n")
			if tt.wantEmpty {
				if got != "" {
					t.Errorf("expected empty, got %q", got)
				}
				return
			}
			if got != tt.wantText {
				t.Errorf("got %q, want %q", got, tt.wantText)
			}
		})
	}
}

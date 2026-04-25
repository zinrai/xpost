package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const endpoint = "/2/tweets"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "xpost:", err)
		os.Exit(1)
	}
}

func run() error {
	if isTerminal(os.Stdin) {
		return errors.New("reads text from stdin (e.g. in Vim: :'<,'>w !xpost)")
	}

	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read stdin: %w", err)
	}

	text := strings.TrimRight(string(raw), " \t\r\n")
	if text == "" {
		return errors.New("empty input")
	}

	return execXurl(text)
}

// execXurl replaces the current process with xurl.
// On success it does not return; the process image is overwritten.
// stdin, stdout, stderr, the exit code, and signals are all handled
// directly by xurl without any mediation by xpost.
func execXurl(text string) error {
	body, err := json.Marshal(struct {
		Text string `json:"text"`
	}{Text: text})
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	xurl, err := exec.LookPath("xurl")
	if err != nil {
		return fmt.Errorf("xurl: %w", err)
	}

	argv := []string{"xurl", "-X", "POST", endpoint, "-d", string(body)}
	return syscall.Exec(xurl, argv, os.Environ())
}

func isTerminal(f *os.File) bool {
	info, err := f.Stat()
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeCharDevice != 0
}

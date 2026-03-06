package gorunner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ExecuteResult holds the result of a Go code execution.
type ExecuteResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
	TimedOut bool
}

// Run compiles and runs Go code with optional stdin.
// Timeout is set to 10 seconds to prevent infinite loops.
func Run(code string, stdin string) (*ExecuteResult, error) {
	tmpDir, err := os.MkdirTemp("", "ngodingyuk-go-")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	mainFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(mainFile, []byte(code), 0644); err != nil {
		return nil, fmt.Errorf("failed to write go file: %w", err)
	}

	cmd := exec.Command("go", "run", mainFile)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	// Start with timeout
	if err := cmd.Start(); err != nil {
		return &ExecuteResult{
			Stderr:   err.Error(),
			ExitCode: 1,
		}, nil
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	timeout := 10 * time.Second
	select {
	case err := <-done:
		exitCode := 0
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode = exitError.ExitCode()
			} else {
				exitCode = 1
			}
		}
		return &ExecuteResult{
			Stdout:   strings.TrimSpace(stdout.String()),
			Stderr:   strings.TrimSpace(stderr.String()),
			ExitCode: exitCode,
		}, nil
	case <-time.After(timeout):
		cmd.Process.Kill()
		return &ExecuteResult{
			Stderr:   "execution timed out (10s limit)",
			ExitCode: 1,
			TimedOut: true,
		}, nil
	}
}

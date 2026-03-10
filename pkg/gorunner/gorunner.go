package gorunner

import (
	"bytes"
	"context"
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

// Program represents a compiled Go program ready to be executed.
type Program struct {
	BinaryPath string
	TempDir    string
}

// Compile compiles the given Go code and returns a Program that can be executed multiple times.
func Compile(code string) (*Program, error) {
	tmpDir, err := os.MkdirTemp("", "ngodingyuk-go-")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	mainFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(mainFile, []byte(code), 0644); err != nil {
		os.RemoveAll(tmpDir)
		return nil, fmt.Errorf("failed to write go file: %w", err)
	}

	binaryPath := filepath.Join(tmpDir, "program")
	cmd := exec.Command("go", "build", "-o", binaryPath, mainFile)

	// Use a persistent cache directory if possible, otherwise use a subdirectory in tmp
	// For now, let's use a subdirectory in the system's temp to avoid permission issues
	// but reuse it across executions if we can. Actually, standard GOCACHE is usually fine
	// if we don't override it to a fresh tmp dir every time.
	// Let's just use the default environment GOCACHE which is persistent.

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		os.RemoveAll(tmpDir)
		return nil, fmt.Errorf("compilation failed: %s: %w", strings.TrimSpace(stderr.String()), err)
	}

	return &Program{
		BinaryPath: binaryPath,
		TempDir:    tmpDir,
	}, nil
}

// Run executes the program with the given stdin and a timeout.
func (p *Program) Run(ctx context.Context, stdin string) (*ExecuteResult, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(timeoutCtx, p.BinaryPath)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin != "" {
		cmd.Stdin = strings.NewReader(stdin)
	}

	err := cmd.Run()

	exitCode := 0
	timedOut := false

	if err != nil {
		if timeoutCtx.Err() == context.DeadlineExceeded {
			timedOut = true
			exitCode = 1
		} else if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = 1
		}
	}

	return &ExecuteResult{
		Stdout:   strings.TrimSpace(stdout.String()),
		Stderr:   strings.TrimSpace(stderr.String()),
		ExitCode: exitCode,
		TimedOut: timedOut,
	}, nil
}

// Cleanup removes the temporary directory and the compiled binary.
func (p *Program) Cleanup() {
	if p.TempDir != "" {
		os.RemoveAll(p.TempDir)
	}
}

// Run is a convenience function that compiles and runs the code once.
// Kept for backward compatibility if needed, but internally uses Program.
func Run(code string, stdin string) (*ExecuteResult, error) {
	p, err := Compile(code)
	if err != nil {
		return nil, err
	}
	defer p.Cleanup()

	return p.Run(context.Background(), stdin)
}

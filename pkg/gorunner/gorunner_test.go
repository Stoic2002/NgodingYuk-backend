package gorunner

import (
	"context"
	"strings"
	"testing"
)

func TestCompileAndRun(t *testing.T) {
	code := `
package main
import "fmt"
func main() {
	fmt.Println("hello from gorunner")
}
`
	res, err := Run(code, "")
	if err != nil {
		t.Fatalf("failed to run code: %v", err)
	}

	if res.ExitCode != 0 {
		t.Errorf("expected exit code 0, got %d. Stderr: %s", res.ExitCode, res.Stderr)
	}

	if !strings.Contains(res.Stdout, "hello from gorunner") {
		t.Errorf("expected output to contain 'hello from gorunner', got %q", res.Stdout)
	}
}

func TestRunWithStdin(t *testing.T) {
	code := `
package main
import (
	"fmt"
	"io"
	"os"
)
func main() {
	input, _ := io.ReadAll(os.Stdin)
	fmt.Printf("received: %s", string(input))
}
`
	res, err := Run(code, "test input")
	if err != nil {
		t.Fatalf("failed to run code: %v", err)
	}

	if res.ExitCode != 0 {
		t.Errorf("expected exit code 0, got %d. Stderr: %s", res.ExitCode, res.Stderr)
	}

	if res.Stdout != "received: test input" {
		t.Errorf("expected output 'received: test input', got %q", res.Stdout)
	}
}

func TestCompilationError(t *testing.T) {
	code := `package main; func main() { undefined() }`
	_, err := Compile(code)
	if err == nil {
		t.Fatal("expected compilation error, got nil")
	}

	if !strings.Contains(err.Error(), "compilation failed") {
		t.Errorf("expected error message to contain 'compilation failed', got %q", err.Error())
	}
}

func TestTimeout(t *testing.T) {
	code := `
package main
import "time"
func main() {
	time.Sleep(15 * time.Second)
}
`
	p, err := Compile(code)
	if err != nil {
		t.Fatalf("failed to compile: %v", err)
	}
	defer p.Cleanup()

	res, err := p.Run(context.Background(), "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !res.TimedOut {
		t.Error("expected TimedOut to be true")
	}
}

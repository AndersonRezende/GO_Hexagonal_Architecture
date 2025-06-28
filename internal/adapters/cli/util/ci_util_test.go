package util

import (
	"bytes"
	"gohexarc/internal/domain"
	"os"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := "test input\n"
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	defer func() {
		os.Stdin = oldStdin
		r.Close()
		w.Close()
	}()

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	result := ReadInput("Testing input: ")
	if result != "test input" {
		t.Errorf("expected 'test input', returned '%s'", result)
	}
}

func TestPrintUser(t *testing.T) {
	user := domain.User{ID: "1", Name: "John Doe", Email: "john.doe@example.com"}
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintUser(user)
	w.Close()
	os.Stdout = oldStdout
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "User ID: 1") ||
		!strings.Contains(output, "John Doe") ||
		!strings.Contains(output, "john.doe@example.com") {
		t.Errorf("unexpected exit: %s", output)
	}
}

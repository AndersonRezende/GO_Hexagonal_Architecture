package util

import (
	"bufio"
	"gohexarc/internal/domain"
	"os"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := "test input\n"
	reader := strings.NewReader(input)
	bufReader := bufio.NewReader(reader)

	result := ReadInput("Testing input: ", bufReader, os.Stdout)

	if result != "test input" {
		t.Errorf("expected 'test input', returned '%s'", result)
	}
}

func TestPrintUser(t *testing.T) {
	user := domain.User{ID: "1", Name: "John Doe", Email: "john.doe@example.com"}

	var sb strings.Builder
	PrintUser(&sb, user)
	output := sb.String()

	if !strings.Contains(output, "User ID: 1") ||
		!strings.Contains(output, "John Doe") ||
		!strings.Contains(output, "john.doe@example.com") {
		t.Errorf("unexpected exit: %s", output)
	}
}

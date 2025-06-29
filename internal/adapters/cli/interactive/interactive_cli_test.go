package interactive

import (
	"bytes"
	"gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/domain"
	"gohexarc/internal/service"
	"testing"
)

func TestInteractive_Run(t *testing.T) {
	repository := new(mock.UserRepository)
	userService := service.NewUserService(repository)

	t.Run("list users", func(t *testing.T) {
		repository.On("List").Return([]domain.User{
			{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
			{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com"},
		}, nil)
		input := bytes.NewBufferString("list\nexit")
		var buf bytes.Buffer

		interactiveCli := NewInteractiveCLI(userService, input, &buf)
		interactiveCli.Run()

		output := buf.String()
		if output == "" {
			t.Errorf("expected command output, but got empty output")
		}
		if !bytes.Contains([]byte(output), []byte("John Doe")) {
			t.Errorf("expected to find 'John Doe', but did not find it")
		}
		if !bytes.Contains([]byte(output), []byte("Jane Doe")) {
			t.Errorf("expected to find 'Jane Doe', but did not find it")
		}
	})

	t.Run("get user", func(t *testing.T) {
		repository.On("GetByID", "1").Return(domain.User{
			ID:    "1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}, nil)

		input := bytes.NewBufferString("get\n1\nexit")
		var buf bytes.Buffer

		interactiveCli := NewInteractiveCLI(userService, input, &buf)
		interactiveCli.Run()
		output := buf.String()

		if output == "" {
			t.Errorf("expected command output, but got empty output")
		}
		if !bytes.Contains([]byte(output), []byte("John Doe")) {
			t.Errorf("expected to find 'John Doe', but did not find it")
		}
	})

	t.Run("create user", func(t *testing.T) {
		repository.On("Create", domain.User{ID: "", Name: "John Doe", Email: "john.doe@example.com"}).Return(nil)
		input := bytes.NewBufferString("create\nJohn Doe\njohn.doe@example.com\nexit")
		var buf bytes.Buffer

		interactiveCli := NewInteractiveCLI(userService, input, &buf)
		interactiveCli.Run()
		output := buf.String()

		if output == "" {
			t.Errorf("expected command output, but got empty output")
		}
		if !bytes.Contains([]byte(output), []byte("John Doe")) {
			t.Errorf("expected to find 'John Doe', but did not find it")
		}
	})

	t.Run("update user", func(t *testing.T) {
		repository.On("Update", domain.User{ID: "1", Name: "John Updated", Email: "john.updated@example.com"}).Return(nil)
		input := bytes.NewBufferString("update\n1\nJohn Updated\njohn.updated@example.com\nexit")
		var buf bytes.Buffer

		interactiveCli := NewInteractiveCLI(userService, input, &buf)
		interactiveCli.Run()
		output := buf.String()

		if output == "" {
			t.Errorf("expected command output, but got empty output")
		}
		if !bytes.Contains([]byte(output), []byte("John Updated")) {
			t.Errorf("expected to find 'John Updated', but did not find it")
		}
	})

	t.Run("delete user", func(t *testing.T) {
		repository.On("Delete", "1").Return(nil)
		input := bytes.NewBufferString("delete\n1\nexit")
		var buf bytes.Buffer

		interactiveCli := NewInteractiveCLI(userService, input, &buf)
		interactiveCli.Run()
		output := buf.String()

		if output == "" {
			t.Errorf("expected command output, but got empty output")
		}
		if !bytes.Contains([]byte(output), []byte("User deleted successfully")) {
			t.Errorf("expected to find 'User deleted successfully', but did not find it")
		}
	})
}

package run_and_die

import (
	"bytes"
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/domain"
	"gohexarc/internal/service"
	"gohexarc/internal/tests"
	"os"
	"testing"
)

func TestRunAndDie_Run(t *testing.T) {
	repository := new(mock.UserRepository)
	userService := service.NewUserService(repository)
	services := registry.Services{UserService: userService}

	t.Run("no command provided", func(t *testing.T) {
		configureTestArgs([]string{})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expectedOutput := "Available cli commands:"
		if !bytes.Contains([]byte(output), []byte(expectedOutput)) {
			t.Errorf("unexpected output: %s, got: %s", expectedOutput, output)
		}
	})

	t.Run("get user", func(t *testing.T) {
		repository.On("GetByID", "1").Return(domain.User{
			ID:    "1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		}, nil)
		configureTestArgs([]string{"get", "1"})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !bytes.Contains([]byte(output), []byte("1")) ||
			!bytes.Contains([]byte(output), []byte("John Doe")) ||
			!bytes.Contains([]byte(output), []byte("john.doe@example.com")) {
			t.Errorf("unexpected output: %s", output)
		}
	})

	t.Run("create user", func(t *testing.T) {
		repository.On("Create", domain.User{Name: "John Doe", Email: "john.doe@example.com"}).Return(nil)
		configureTestArgs([]string{"create", "John Doe", "john.doe@example.com"})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedOutput := "User created successfully"
		if !bytes.Contains([]byte(output), []byte(expectedOutput)) {
			t.Errorf("unexpected output: %s, got: %s", expectedOutput, output)
		}
	})

	t.Run("update user", func(t *testing.T) {
		repository.On("Update", domain.User{ID: "123", Name: "Jane Doe", Email: "john.doe@example.com"}).Return(nil)
		configureTestArgs([]string{"update", "123", "Jane Doe", "john.doe@example.com"})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedOutput := "User updated successfully"
		if !bytes.Contains([]byte(output), []byte(expectedOutput)) {
			t.Errorf("unexpected output: %s, got: %s", expectedOutput, output)
		}
	})

	t.Run("delete user", func(t *testing.T) {
		repository.On("Delete", "123").Return(nil)
		configureTestArgs([]string{"delete", "123"})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedOutput := "User deleted successfully"
		if !bytes.Contains([]byte(output), []byte(expectedOutput)) {
			t.Errorf("unexpected output: %s, got: %s", expectedOutput, output)
		}
	})

	t.Run("list users", func(t *testing.T) {
		repository.On("List").Return([]domain.User{
			{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
			{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com"},
		}, nil)
		configureTestArgs([]string{"list"})
		var buf bytes.Buffer
		runAndDieCli := NewRunAndDieCLI(&services, &buf)

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
		output += buf.String()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !bytes.Contains([]byte(output), []byte("User ID: 1")) ||
			!bytes.Contains([]byte(output), []byte("John Doe")) ||
			!bytes.Contains([]byte(output), []byte("john.doe@example.com")) ||
			!bytes.Contains([]byte(output), []byte("User ID: 2")) ||
			!bytes.Contains([]byte(output), []byte("Jane Doe")) ||
			!bytes.Contains([]byte(output), []byte("jane.doe@example.com")) {
			t.Errorf("unexpected output: %s", output)
		}

		t.Run("unknown command", func(t *testing.T) {
			configureTestArgs([]string{"unknown_command"})
			var buf bytes.Buffer
			runAndDieCli := NewRunAndDieCLI(&services, &buf)

			output, err := tests.ExecCliFunction(runAndDieCli.Run)
			output += buf.String()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			expectedOutput := "Unknown command: unknown_command"
			if !bytes.Contains([]byte(output), []byte(expectedOutput)) {
				t.Errorf("unexpected output: %s, got: %s", expectedOutput, output)
			}
		})
	})
}

func configureTestArgs(command []string) {
	os.Args = []string{"run_and_die", "cli"}
	for _, arg := range command {
		os.Args = append(os.Args, arg)
	}
}

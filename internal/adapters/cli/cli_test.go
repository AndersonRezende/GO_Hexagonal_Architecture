package cli

import (
	"bytes"
	"gohexarc/internal/adapters/cli/run_and_die"
	"gohexarc/internal/adapters/repository/mock"
	UserRepositoryMock "gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
	"gohexarc/internal/service"
	UserService "gohexarc/internal/service"
	"gohexarc/internal/tests"
	"os"
	"testing"
)

type mockCLI struct {
	runCalled        bool
	listUsersCalled  bool
	getUserCalled    bool
	createUserCalled bool
	updateUserCalled bool
	deleteUserCalled bool
}

func (m *mockCLI) Run() {
	m.runCalled = true
	m.ListUsers()
	m.GetUser()
	m.CreateUser()
	m.UpdateUser()
	m.DeleteUser()
}
func (m *mockCLI) ListUsers()  { m.listUsersCalled = true }
func (m *mockCLI) GetUser()    { m.getUserCalled = true }
func (m *mockCLI) CreateUser() { m.createUserCalled = true }
func (m *mockCLI) UpdateUser() { m.updateUserCalled = true }
func (m *mockCLI) DeleteUser() { m.deleteUserCalled = true }

func TestRunCliCalls(t *testing.T) {
	t.Run("Run is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.runCalled {
			t.Errorf("RunCli did not call Run on the CLI")
		}
	})

	t.Run("ListUsers is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.listUsersCalled {
			t.Errorf("RunCli did not call ListUsers on the CLI")
		}
	})

	t.Run("GetUser is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.getUserCalled {
			t.Errorf("RunCli did not call GetUser on the CLI")
		}
	})

	t.Run("CreateUser is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.createUserCalled {
			t.Errorf("RunCli did not call CreateUser on the CLI")
		}
	})

	t.Run("UpdateUser is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.updateUserCalled {
			t.Errorf("RunCli did not call UpdateUser on the CLI")
		}
	})

	t.Run("DeleteUser is called", func(t *testing.T) {
		cliMock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return cliMock
		}
		repository := new(UserRepositoryMock.UserRepository)
		userService := UserService.NewUserService(repository)
		RunCli(userService, factory)
		if !cliMock.deleteUserCalled {
			t.Errorf("RunCli did not call DeleteUser on the CLI")
		}
	})
}

func TestRunAndDie_Run(t *testing.T) {
	repository := new(mock.UserRepository)
	userService := service.NewUserService(repository)
	runAndDieCli := run_and_die.NewRunAndDieCLI(userService)

	t.Run("no command provided", func(t *testing.T) {
		configureTestArgs([]string{})
		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

		output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

			output, err := tests.ExecCliFunction(runAndDieCli.Run)
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

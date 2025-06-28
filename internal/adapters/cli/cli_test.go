package cli

import (
	UserRepositoryMock "gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/port"
	UserService "gohexarc/internal/service"
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

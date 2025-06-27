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
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.runCalled {
			t.Errorf("RunCli did not call Run on the CLI")
		}
	})

	t.Run("ListUsers is called", func(t *testing.T) {
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.listUsersCalled {
			t.Errorf("RunCli did not call ListUsers on the CLI")
		}
	})

	t.Run("GetUser is called", func(t *testing.T) {
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.getUserCalled {
			t.Errorf("RunCli did not call GetUser on the CLI")
		}
	})

	t.Run("CreateUser is called", func(t *testing.T) {
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.createUserCalled {
			t.Errorf("RunCli did not call CreateUser on the CLI")
		}
	})

	t.Run("UpdateUser is called", func(t *testing.T) {
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.updateUserCalled {
			t.Errorf("RunCli did not call UpdateUser on the CLI")
		}
	})

	t.Run("DeleteUser is called", func(t *testing.T) {
		mock := &mockCLI{}
		factory := func(service port.UserService) CLI {
			return mock
		}
		repository := new(UserRepositoryMock.UserRepository)
		service := UserService.NewUserService(repository)
		RunCli(service, factory)
		if !mock.deleteUserCalled {
			t.Errorf("RunCli did not call DeleteUser on the CLI")
		}
	})
}

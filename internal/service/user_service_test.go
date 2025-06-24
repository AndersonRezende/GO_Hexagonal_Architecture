package service

import (
	"gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/domain"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repository := new(mock.UserRepository)
	service := NewUserService(repository)
	expectedUser := domain.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	repository.On("Create", expectedUser).Return(nil)

	user, err := service.CreateUser(expectedUser.Name, expectedUser.Email)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.Name != expectedUser.Name || user.Email != expectedUser.Email {
		t.Fatalf("expected user name to be %s and email to be %s, got name %s and email %s",
			expectedUser.Name, expectedUser.Email, user.Name, user.Email)
	}
	repository.AssertCalled(t, "Create", expectedUser)
}

func TestGetUser(t *testing.T) {
	repository := new(mock.UserRepository)
	service := NewUserService(repository)
	expectedUser := domain.User{
		ID:    "123",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	repository.On("GetByID", "123").Return(expectedUser, nil)

	user, err := service.GetUser("123")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user.ID != expectedUser.ID || user.Name != expectedUser.Name || user.Email != expectedUser.Email {
		t.Fatalf("expected user ID to be %s, name to be %s and email to be %s, got ID %s, name %s and email %s",
			expectedUser.ID, expectedUser.Name, expectedUser.Email, user.ID, user.Name, user.Email)
	}
	repository.AssertCalled(t, "GetByID", "123")
}

func TestUpdateUser(t *testing.T) {
	repository := new(mock.UserRepository)
	service := NewUserService(repository)
	updatedUser := domain.User{
		ID:    "123",
		Name:  "Jane Doe",
		Email: "jane.doe@example.com",
	}
	repository.On("Update", updatedUser).Return(nil)

	err := service.UpdateUser(updatedUser.ID, updatedUser.Name, updatedUser.Email)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	repository.AssertCalled(t, "Update", updatedUser)
}

func TestDeleteUser(t *testing.T) {
	repository := new(mock.UserRepository)
	service := NewUserService(repository)
	userID := "123"
	repository.On("Delete", userID).Return(nil)

	err := service.DeleteUser(userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	repository.AssertCalled(t, "Delete", userID)
}

func TestListUsers(t *testing.T) {
	repository := new(mock.UserRepository)
	service := NewUserService(repository)
	expectedUsers := []domain.User{
		{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com"},
	}
	repository.On("List").Return(expectedUsers, nil)

	users, err := service.ListUsers()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(users) != len(expectedUsers) {
		t.Fatalf("expected %d users, got %d", len(expectedUsers), len(users))
	}
	for i, user := range users {
		if user.ID != expectedUsers[i].ID || user.Name != expectedUsers[i].Name || user.Email != expectedUsers[i].Email {
			t.Fatalf("expected user %d to be %v, got %v", i, expectedUsers[i], user)
		}
	}
	repository.AssertCalled(t, "List")
}

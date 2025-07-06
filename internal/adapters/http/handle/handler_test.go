package handle

import (
	"bytes"
	"encoding/json"
	"gohexarc/cmd/registry"
	UserRepositoryMock "gohexarc/internal/adapters/repository/mock"
	"gohexarc/internal/domain"
	UserService "gohexarc/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestServer() (*httptest.Server, *UserRepositoryMock.UserRepository) {
	repo := new(UserRepositoryMock.UserRepository)
	service := UserService.NewUserService(repo)
	services := &registry.Services{UserService: service}
	mux := http.NewServeMux()
	RegisterHandlers(mux, services)
	return httptest.NewServer(mux), repo
}

func TestUserHandlers(t *testing.T) {
	server, repo := setupTestServer()
	defer server.Close()

	t.Run("Test GET - List", func(t *testing.T) {
		repo.On("List").Return([]domain.User{
			{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
		}, nil)

		resp, err := http.Get(server.URL + "/users")
		if err != nil {
			t.Fatalf("GET /users failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("GET /users: expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Test GET - Get by id", func(t *testing.T) {
		repo.On("GetByID", "1").Return(domain.User{
			ID: "1", Name: "John Doe", Email: "john.doe@example.com",
		}, nil)

		resp, err := http.Get(server.URL + "/users?id=1")
		if err != nil {
			t.Fatalf("GET /users failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("GET /users?id=1 expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Test POST - Create", func(t *testing.T) {
		repo.On("Create", domain.User{ID: "", Name: "John Doe", Email: "john.doe@example.com"}).Return(nil)

		user := domain.User{ID: "", Name: "John Doe", Email: "john.doe@example.com"}
		body, _ := json.Marshal(user)
		resp, err := http.Post(server.URL+"/users", "application/json", bytes.NewReader(body))
		if err != nil {
			t.Fatalf("POST /users failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("POST /users: expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Test PUT - Update", func(t *testing.T) {
		repo.On("Update", domain.User{ID: "2", Name: "John Doe", Email: "john.doe@example.com"}).Return(nil)

		user := domain.User{ID: "2", Name: "John Doe", Email: "john.doe@example.com"}
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPut, server.URL+"/users", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("PUT /users failed: %v", err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("PUT /users: expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("Test DELETE - Delete", func(t *testing.T) {})
	repo.On("Delete", "2").Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, server.URL+"/users?id=2", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("DELETE /users failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("DELETE /users: expected 200, got %d", resp.StatusCode)
	}
}

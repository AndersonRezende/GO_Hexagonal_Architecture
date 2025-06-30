package handle

import (
	"encoding/json"
	"gohexarc/cmd/registry"
	"gohexarc/internal/domain"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, services *registry.Services) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(services, w, r)
		case http.MethodPost:
			handlePost(services, w, r)
		case http.MethodPut:
			handlePut(services, w, r)
		case http.MethodDelete:
			handleDelete(services, w, r)
		}
	})
}

func handleGet(services *registry.Services, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		user, err := services.UserService.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		encode(w, user)
		return
	}
	users, _ := services.UserService.ListUsers()
	encode(w, users)
}

func handlePost(services *registry.Services, w http.ResponseWriter, r *http.Request) {
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	newUser, err := services.UserService.CreateUser(u.Name, u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, newUser)
}

func handlePut(services *registry.Services, w http.ResponseWriter, r *http.Request) {
	var u domain.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := services.UserService.UpdateUser(u.ID, u.Name, u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, map[string]string{"status": "updated"})
}

func handleDelete(services *registry.Services, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	err := services.UserService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, map[string]string{"status": "deleted"})
}

func encode(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

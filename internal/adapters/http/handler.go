package http

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
			id := r.URL.Query().Get("id")
			if id != "" {
				user, err := services.UserService.GetUser(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(user)
				return
			}
			users, _ := services.UserService.ListUsers()
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			var u domain.User
			json.NewDecoder(r.Body).Decode(&u)
			newUser, _ := services.UserService.CreateUser(u.Name, u.Email)
			json.NewEncoder(w).Encode(newUser)
		case http.MethodPut:
			var u domain.User
			json.NewDecoder(r.Body).Decode(&u)
			err := services.UserService.UpdateUser(u.ID, u.Name, u.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
		case http.MethodDelete:
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
			json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
		}
	})
}

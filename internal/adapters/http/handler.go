package http

import (
	"encoding/json"
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, service port.UserService) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			users, _ := service.ListUsers()
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			var u domain.User
			json.NewDecoder(r.Body).Decode(&u)
			newUser, _ := service.CreateUser(u.Name, u.Email)
			json.NewEncoder(w).Encode(newUser)
		}
	})
}

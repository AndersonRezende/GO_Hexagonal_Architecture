package http

import (
	"fmt"
	"gohexarc/cmd/registry"
	"net/http"
)

func ServeHTTP(services *registry.Services) {
	mux := http.NewServeMux()
	RegisterHandlers(mux, services)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}

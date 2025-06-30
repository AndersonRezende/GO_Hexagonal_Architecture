package http

import (
	"fmt"
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/http/handle"
	"net/http"
)

func ServeHTTP(services *registry.Services) {
	mux := http.NewServeMux()
	handle.RegisterHandlers(mux, services)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}

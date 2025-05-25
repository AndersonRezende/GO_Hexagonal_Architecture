package main

import (
	"fmt"
	"gohexarc/internal/adapters/cli"
	http2 "gohexarc/internal/adapters/http"
	"gohexarc/internal/adapters/repository"
	"gohexarc/internal/service"
	"net/http"
	"os"
)

func main() {
	repo := repository.NewInMemoryUserRepository()
	service := service.NewUserService(repo)

	if len(os.Args) > 1 && os.Args[1] == "cli" {
		cli.RunCli(service)
		return
	}

	mux := http.NewServeMux()
	http2.RegisterHandlers(mux, service)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}

package main

import (
	"fmt"
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/cli"
	http2 "gohexarc/internal/adapters/http"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	services := registry.NewServices()

	if len(os.Args) > 1 && os.Args[1] == "cli" {
		cli.RunCli(services, cli.DefaultCLIFactory)
		return
	}

	mux := http.NewServeMux()
	http2.RegisterHandlers(mux, services.UserService)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}

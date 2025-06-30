package main

import (
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/cli"
	"gohexarc/internal/adapters/http"
	"os"
)

func main() {
	services := registry.NewServices()

	if shouldExecuteCli() {
		cli.RunCli(services, cli.DefaultCLIFactory)
		return
	}
	http.ServeHTTP(services)
}

func shouldExecuteCli() bool {
	if len(os.Args) > 1 {
		return os.Args[1] == "cli"
	}
	return false
}

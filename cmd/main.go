package main

import (
	"database/sql"
	"fmt"
	"gohexarc/cmd/registry"
	"gohexarc/internal/adapters/cli"
	http2 "gohexarc/internal/adapters/http"
	"gohexarc/internal/adapters/repository/memory"
	"gohexarc/internal/adapters/repository/sqlite"
	"gohexarc/internal/port"
	//"gohexarc/internal/service"
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

func getRepository() port.UserRepository {
	var userRepository port.UserRepository
	if shouldUseInMemory() {
		userRepository = memory.NewInMemoryUserRepository()
	} else {
		db, err := sql.Open("sqlite3", "data.db")
		if err != nil {
			errorMessage := fmt.Errorf("could not open database: %v", err)
			fmt.Println(errorMessage)
		}
		userRepository = sqlite.NewSqliteUserRepository(db)
	}
	return userRepository
}

func shouldUseInMemory() bool {
	fileBytes, err := os.ReadFile(".env")
	if err == nil && string(fileBytes) == "USE_IN_MEMORY=true" {
		return false
	}
	return false
}

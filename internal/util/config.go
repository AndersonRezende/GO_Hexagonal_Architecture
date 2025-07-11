package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func ShouldUseInMemory() bool {
	fileBytes, err := os.ReadFile("../.env")
	return err == nil && string(fileBytes) == "USE_IN_MEMORY=true"
}

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data/data.db")
	if err != nil {
		errorMessage := fmt.Errorf("could not open database: %v", err)
		fmt.Println(errorMessage)
	}

	return db, err
}

package util

import (
	"database/sql"
	"fmt"
	"os"
)

func ShouldUseInMemory() bool {
	fileBytes, err := os.ReadFile("../.env")
	return err == nil && string(fileBytes) == "USE_IN_MEMORY=true"
}

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing database:", err)
		}
	}(db)
	if err != nil {
		errorMessage := fmt.Errorf("could not open database: %v", err)
		fmt.Println(errorMessage)
	}

	return db, err
}

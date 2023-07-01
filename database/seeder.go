package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func CreateTable(db *sqlx.DB) {
	tableCreationFile, err := os.ReadFile("database/init.sql")
	if err != nil {
		panic(fmt.Sprintf("Failed to read SQL file to create tables: %v", err))
	}
	if _, err = db.Exec(string(tableCreationFile)); err != nil {
		panic(fmt.Sprintf("Failed to execute SQL file to create tables: %v", err))
	}
}

func SeedTable(db *sqlx.DB) {
	var count int
	err := db.Get(&count, "select count(id) from Products")
	if err != nil {
		panic(fmt.Sprintf("Failed to get count of Products: %v", err))
	}

	if count < 30 {
		seedsFile, err := os.ReadFile("database/seed.sql")
		if err != nil {
			panic(fmt.Sprintf("Failed to read SQL file to seed data: %v", err))
		}
		if _, err = db.Exec(string(seedsFile)); err != nil {
			panic(fmt.Sprintf("Failed to execute SQL file to seed data: %v", err))
		}
	}
}

package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) {
	if _, err := db.Exec(createProductsTable); err != nil {
		panic(fmt.Sprintf("Failed to create product table: %v", err))
	}
	if _, err := db.Exec(createUserTable); err != nil {
		panic(fmt.Sprintf("Failed to create user table: %v", err))
	}
	if _, err := db.Exec(createSessionTable); err != nil {
		panic(fmt.Sprintf("Failed to create session table: %v", err))
	}
	if _, err := db.Exec(createCartItemTable); err != nil {
		panic(fmt.Sprintf("Failed to create cartItem table: %v", err))
	}
	if _, err := db.Exec(createCartTable); err != nil {
		panic(fmt.Sprintf("Failed to create cart table: %v", err))
	}
}

func SeedTable(db *sqlx.DB) {
	var count int
	err := db.Get(&count, "select count(id) from Products")
	if err != nil {
		panic(fmt.Sprintf("Failed to get count of Products: %v", err))
	}

	if count < 30 {
		if _, err = db.Exec(insertProductsRecords); err != nil {
			panic(fmt.Sprintf("Failed to populate products record: %v", err))
		}
	}
}

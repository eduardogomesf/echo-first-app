package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func createProductsTable(dbConn *pgx.Conn) {
	createProductTableSql := `
		CREATE TABLE IF NOT EXISTS products (
			id UUID PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			price DECIMAL NOT NULL,
			categories TEXT[] NOT NULL,
			is_disabled BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
	);`

	_, err := dbConn.Exec(context.Background(), createProductTableSql)

	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	fmt.Println("Table 'products' created successfully (if it didn't already exist).")
}

func CreateTables(dbConn *pgx.Conn) {
	createProductsTable(dbConn)
}

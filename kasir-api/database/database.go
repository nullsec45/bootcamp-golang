package database

import (
	"database/sql"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"fmt"
)

func InitDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	log.Println("Database connected successfully")
	return db, nil
}
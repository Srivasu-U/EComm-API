package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN()) // Open connection to MySQL DB
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

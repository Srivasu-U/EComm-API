package main

import (
	"database/sql"
	"log"

	"github.com/Srivasu-U/EComm-API/cmd/api"
	"github.com/Srivasu-U/EComm-API/config"
	"github.com/Srivasu-U/EComm-API/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping() // db.Open() only creates a connection to the DB. db.Ping() verifies the connection
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Ping successful")
}

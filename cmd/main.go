package main

import (
	"log"

	"github.com/Srivasu-U/EComm-API/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

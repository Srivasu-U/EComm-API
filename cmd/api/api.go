package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Srivasu-U/EComm-API/service/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	// Subrouters are used to server everything in a particular namespace.
	// In this call, all our endpoints have the v1 prefix, which is good practice in general
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRouters(subrouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

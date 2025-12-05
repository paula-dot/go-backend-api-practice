package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr, db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userhandler := user.NewHandler()
	userhandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(s.addr, router)
}

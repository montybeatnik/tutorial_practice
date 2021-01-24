package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type server struct {
	Router *mux.Router
	DB     *sql.DB
}

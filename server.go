package main

import (
	"database/sql"
	"html/template"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Server struct {
	Template *template.Template
	Router   *mux.Router
	DB       *sql.DB
}

func (s *Server) Initialize(user, password, dbname string) {}
func (s *Server) Run(addr string)                          {}

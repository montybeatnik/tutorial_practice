package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "devcon"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func nothing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func run() error {
	log.Println("setting up the db...")
	db, err := setupDatabase()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer db.Close()
	r := mux.NewRouter()
	srv := &server{
		// db:     db,
		Router: r,
	}
	srv.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	srv.routes()
	log.Println("server listening on localhost:8000...")
	if err := http.ListenAndServe(":8000", srv.Router); err != nil {
		return err
	}
	return nil
}

func setupDatabase() (*sql.DB, error) {
	// db info
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// connect
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	return db, nil
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/montybeatnik/tutorial_practice/autochecks"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type server struct {
	db     *sql.DB
	router *mux.Router
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/home.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

func (s *server) handleAbout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "about")
	}
}

func (s *server) handleOutline() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("views/outline.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

func (s *server) handleDeviceShowVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var swVer autochecks.SoftwareVersion
		p := autochecks.Params{
			IP: "192.168.1.1",
		}
		_, err := swVer.Run(p)
		if err != nil {
			fmt.Fprintf(w, "success")
		}
	}
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/montybeatnik/tutorial_practice/autochecks"
	"github.com/montybeatnik/tutorial_practice/driver"
	"github.com/montybeatnik/tutorial_practice/models"
	"github.com/pkg/errors"

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

func initializeDevPSQL() (*sql.DB, error) {
	conf := driver.PSQLConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
		DB:       "tutorial_practice",
	}

	var err error
	db, err := driver.ConnectToPSQL(conf)
	if err != nil {
		return db, errors.Wrap(err, "could not connect to postgres")
	}

	return db, nil
}

func (s *server) handleDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := initializeDevPSQL()
		if err != nil {
			log.Println(err)
			return
		}
		devService := models.NewDeviceStore(db)
		d, err := devService.GetById(2)
		if err != nil {
			log.Println(err)
			return
		}
		dev := struct {
			hn    string
			lo    string
			model string
			ven   string
			ver   string
		}{d.Hostname, d.Loopback, d.Model, d.Vendor, d.Version}

		fmt.Fprintf(w, fmt.Sprintf("%v, %v, %v, %v, %v", dev.hn, dev.lo, dev.model, dev.ven, dev.ver))
	}
}

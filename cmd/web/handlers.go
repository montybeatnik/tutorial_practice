package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/montybeatnik/tutorial_practice/driver"
	"github.com/montybeatnik/tutorial_practice/pkg/autochecks"
	"github.com/montybeatnik/tutorial_practice/pkg/models"
	"github.com/montybeatnik/tutorial_practice/pkg/scan"
	"github.com/montybeatnik/tutorial_practice/views"
	"github.com/pkg/errors"
)

var (
	homeView    *views.View
	aboutView   *views.View
	outlineView *views.View
	indexView   *views.View
)

func (s *Server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := homeView.Template.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleAbout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := aboutView.Template.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleOutline() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := outlineView.Template.Execute(w, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleDeviceShowVersion() http.HandlerFunc {
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

func (s *Server) handleDeviceIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := initializeDevPSQL()
		if err != nil {
			log.Println(err)
			return
		}
		devService := models.NewDeviceStore(db)
		devices, err := devService.AllDevices()
		if err != nil {
			log.Println(err)
			// TODO: throw HTML error to UI
		}
		for _, d := range devices {
			log.Println(d.Hostname, d.Loopback)
		}
	}
}

func (s *Server) handleDeviceID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, convErr := strconv.Atoi(idStr)
		if convErr != nil {
			log.Println("id is not an integer!", convErr)
		}
		db, err := initializeDevPSQL()
		if err != nil {
			log.Println(err)
			return
		}
		devService := models.NewDeviceStore(db)
		d, err := devService.GetById(uint(id))
		if err != nil {
			log.Println(err)
			return
		}

		t, err := template.ParseFiles("views/device_info.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, d)
	}
}

type deviceForm struct {
	Hostname string `schema:"hostname"`
}

// handleDeviceHostname PUT
func (s *Server) handleDeviceHostname() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var decoder = schema.NewDecoder()
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		var form deviceForm

		err = decoder.Decode(&form, r.PostForm)
		if err != nil {
			log.Println(err)
		}

		db, err := initializeDevPSQL()
		if err != nil {
			log.Println(err)
			return
		}
		devService := models.NewDeviceStore(db)
		d, err := devService.GetByHostname(form.Hostname)
		if err != nil {
			log.Println(err)
			return
		}

		t, err := template.ParseFiles("views/device_info.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, d)
	}
}

func (s *Server) HandleScanDevices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := initializeDevPSQL()
		if err != nil {
			log.Println(err)
			return
		}
		devService := models.NewDeviceStore(db)
		devices, err := devService.AllDevices()
		if err != nil {
			log.Println(err)
			return
		}
		var verAC autochecks.SoftwareVersion
		scan.Devices(devices, &verAC)
		t, err := template.ParseFiles("views/software_ver_results.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

// HandleScanSubnet takes in a subnet via web form.
func (s *Server) HandleScanSubnet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var verAC autochecks.SoftwareVersion
		subnet, err := scan.Hosts("10.1.1.0/24")
		if err != nil {
			log.Println(err)
		}
		scan.Subnets(subnet, &verAC)
		t, err := template.ParseFiles("views/software_ver_results.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

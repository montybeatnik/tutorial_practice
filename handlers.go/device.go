package handlers

import (
	"database/sql"
	"fmt"

	"github.com/montybeatnik/tutorial_practice/driver"
	"github.com/montybeatnik/tutorial_practice/models"
	"github.com/pkg/errors"
)

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

func HandleDevice() {
	db, err := initializeDevPSQL(t)
	if err != nil {
		return
	}
	devService := models.NewDeviceStore(db)
	d, err := devService.GetById(1)
	if err != nil {
		return
	}
	fmt.Println(d)

}

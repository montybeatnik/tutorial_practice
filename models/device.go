package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Device lays out the methods to interace with the device
// table in the database.
type Devicer interface {
	GetById(id uint) (Device, error)
}

type Device struct {
	ID              uint
	Hostname        string
	Loopback        string
	Model           uint
	SoftwareVersion uint
}

func NewDeviceStore(db *sql.DB) DeviceStorer {
	return DeviceStorer{db: db}
}

type DeviceStorer struct {
	db *sql.DB
}

func (d *DeviceStorer) GetById(id uint) (Device, error) {
	// var (
	// 	query := fmt.Sprintf("SELECT ")
	// 	err error
	// )

	if id != 0 {

	}

	// d.db.QueryRow()
	return Device{}, nil
}

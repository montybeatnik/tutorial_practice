package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Devicer lays out the methods to interace with the device
// table in the database.
type Devicer interface {
	GetById(id uint) (Device, error)
}

type Device struct {
	ID        uint
	CreatedAt time.Time
	Hostname  string
	Loopback  string
	Hardware
	Software
}

func NewDeviceStore(db *sql.DB) DeviceStorer {
	return DeviceStorer{db: db}
}

type DeviceStorer struct {
	db *sql.DB
}

func (d *DeviceStorer) GetById(id uint) (Device, error) {

	var dev Device

	q := `SELECT devices.id, devices.created_at, hostname, loopback,  hardware.vendor, hardware.model, software.version from devices
	JOIN hardware on hardware.id = hardware_id
	JOIN software on software.id = software_id
	where devices.id = $1`
	err := d.db.QueryRow(q, id).Scan(&dev.ID, &dev.CreatedAt, &dev.Hostname, &dev.Loopback, &dev.Hardware.Vendor, &dev.Hardware.Model, &dev.Software.Version)
	if err != nil {
		return dev, err
	}
	return dev, nil
}

func (d *DeviceStorer) Create(dev Device) error {
	_, err := d.db.Exec(`INSERT INTO devices (created_at, hostname, loopback, hardware_id, software_id) VALUES(now(), $1, $2, $3, $4) RETURNING id`, dev.Hostname, dev.Loopback, dev.Hardware.ID, dev.Software.ID)
	if err != nil {
		return err
	}
	return nil
}

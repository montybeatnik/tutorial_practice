package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// Hardware lays out the methods to interace with the Hardware
// table in the database.
type Hardwareer interface {
	GetById(id uint) (Hardware, error)
}

type Hardware struct {
	ID        uint
	CreatedAt time.Time
	Vendor    string
	Model     string
}

func NewHardwareStore(db *sql.DB) HardwareStorer {
	return HardwareStorer{db: db}
}

type HardwareStorer struct {
	db *sql.DB
}

func (h *HardwareStorer) GetById(id uint) (Hardware, error) {

	var hw Hardware

	q := "SELECT id, created_at, vendor, model from hardware where id = $1"
	err := h.db.QueryRow(q, id).Scan(&hw.ID, &hw.CreatedAt, &hw.Vendor, &hw.Model)
	if err != nil {
		return hw, err
	}
	return hw, nil
}

func (h *HardwareStorer) Create(hw Hardware) error {
	_, err := h.db.Exec("INSERT INTO hardware (created_at, vendor, model) VALUES(now(), $1, $2)", hw.Vendor, hw.Model)
	if err != nil {
		return err
	}
	return nil
}

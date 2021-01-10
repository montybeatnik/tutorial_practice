package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// SoftwareVersion lays out the methods to interace with the SoftwareVersion
// table in the database.
type SoftwareVersioner interface {
	GetById(id uint) (SoftwareVersion, error)
}

type SoftwareVersion struct {
	ID        uint
	CreatedAt time.Time
	Version   string
}

func NewSoftwareVersionStore(db *sql.DB) SoftwareVersionStorer {
	return SoftwareVersionStorer{db: db}
}

type SoftwareVersionStorer struct {
	db *sql.DB
}

func (sw *SoftwareVersionStorer) GetById(id uint) (SoftwareVersion, error) {

	var v SoftwareVersion

	q := "SELECT id, created_at, version from software_versions where id = $1"
	err := sw.db.QueryRow(q, id).Scan(&v.ID, &v.CreatedAt, &v.Version)
	if err != nil {
		return v, err
	}
	fmt.Println(v.CreatedAt)
	return v, nil
}

func (sw *SoftwareVersionStorer) Create(v SoftwareVersion) error {
	_, err := sw.db.Exec("INSERT INTO software_versions (created_at, version) VALUES(now(), $1)", v.Version)
	if err != nil {
		return err
	}
	return nil
}

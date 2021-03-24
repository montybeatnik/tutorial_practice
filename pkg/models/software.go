package models

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"time"

	_ "github.com/lib/pq"
)

type SoftwareVersion struct {
	SoftwareInformation struct {
		Text               string `xml:",chardata"`
		HostName           string `xml:"host-name"`
		ProductModel       string `xml:"product-model"`
		ProductName        string `xml:"product-name"`
		JunosVersion       string `xml:"junos-version"`
		PackageInformation []struct {
			Text    string `xml:",chardata"`
			Name    string `xml:"name"`
			Comment string `xml:"comment"`
		} `xml:"package-information"`
	} `xml:"software-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

// Mapper takes in a bytes buffer and unmarshals the buffer into the sv object
func (sv *SoftwareVersion) Mapper(b bytes.Buffer) error {
	if err := xml.Unmarshal(b.Bytes(), &sv); err != nil {
		return err
	}
	return nil
}

// SoftwareVersion lays out the methods to interace with the SoftwareVersion
// table in the database.
type SoftwareVersioner interface {
	GetById(id uint) (Software, error)
}

type Software struct {
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

func (sw *SoftwareVersionStorer) GetById(id uint) (Software, error) {

	var v Software

	q := "SELECT id, created_at, version from software where id = $1"
	err := sw.db.QueryRow(q, id).Scan(&v.ID, &v.CreatedAt, &v.Version)
	if err != nil {
		return v, err
	}
	return v, nil
}

func (sw *SoftwareVersionStorer) Create(v Software) error {
	_, err := sw.db.Exec("INSERT INTO software (created_at, version) VALUES(now(), $1)", v.Version)
	if err != nil {
		return err
	}
	return nil
}

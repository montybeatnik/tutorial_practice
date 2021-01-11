package models

import (
	"database/sql"
	"testing"

	"github.com/montybeatnik/tutorial_practice/driver"
)

func initializeHWPSQL(t *testing.T) *sql.DB {
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
		t.Errorf("could not connect to sql, err:%v", err)
	}

	return db
}

func TestHardwareDatastore(t *testing.T) {
	db := initializeHWPSQL(t)
	ver := NewHardwareStore(db)
	testHardwareStorer_Create(t, ver)
	testHardwareStorer_Get(t, ver)

}

func testHardwareStorer_Create(t *testing.T, db HardwareStorer) {
	testcases := []struct {
		req      Hardware
		response Hardware
	}{
		{Hardware{ID: 1, Vendor: "juniper", Model: "ACX1100"}, Hardware{ID: 1, Vendor: "juniper", Model: "ACX1100"}},
		{Hardware{ID: 2, Vendor: "dell", Model: "M1K"}, Hardware{ID: 2, Vendor: "dell", Model: "M1K"}},
	}
	for _, v := range testcases {
		err := db.Create(v.req)
		if err.Error() == "pq: duplicate key value violates unique constraint \"hardware_vendor_model_key\"" {
			continue
		}
		if err != nil {
			t.Errorf("problem creating test case: %v", err)
		}
	}
}

func testHardwareStorer_Get(t *testing.T, db HardwareStorer) {
	testcases := []struct {
		id   uint
		resp Hardware
	}{
		{1, Hardware{ID: 1, Vendor: "juniper"}},
		{2, Hardware{ID: 2, Vendor: "dell"}},
	}
	for _, v := range testcases {
		resp, err := db.GetById(v.id)
		if err != nil {
			t.Errorf("problem getting Vendor:%v", err)
		}

		if resp.ID != v.resp.ID && resp.Vendor != v.resp.Vendor {
			t.Errorf("Expected: %v, Got: %v", v, resp)
		}
	}
}

package models

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/montybeatnik/tutorial_practice/driver"
)

func initializeSWPSQL(t *testing.T) *sql.DB {
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

	fmt.Println(db.Ping())
	return db
}

func TestDatastore(t *testing.T) {
	db := initializeSWPSQL(t)
	ver := NewSoftwareVersionStore(db)
	testSoftwareVersionStorer_Create(t, ver)
	testSoftwareVersionStorer_Get(t, ver)

}

func testSoftwareVersionStorer_Create(t *testing.T, db SoftwareVersionStorer) {
	testcases := []struct {
		req      SoftwareVersion
		response SoftwareVersion
	}{
		{SoftwareVersion{ID: 1, Version: "12.1"}, SoftwareVersion{ID: 1, Version: "12.1"}},
		{SoftwareVersion{ID: 2, Version: "14.2"}, SoftwareVersion{ID: 2, Version: "14.2"}},
	}
	for _, v := range testcases {
		err := db.Create(v.req)
		if err != nil {
			t.Errorf("problem creating test case: %v", err)
		}
	}
}

func testSoftwareVersionStorer_Get(t *testing.T, db SoftwareVersionStorer) {
	testcases := []struct {
		id   uint
		resp SoftwareVersion
	}{
		{1, SoftwareVersion{ID: 1, Version: "12.1"}},
		{2, SoftwareVersion{ID: 2, Version: "14.2"}},
	}
	for _, v := range testcases {
		resp, err := db.GetById(v.id)
		if err != nil {
			t.Errorf("problem getting version:%v", err)
		}

		if resp.ID != v.resp.ID && resp.Version != v.resp.Version {
			t.Errorf("Expected: %v, Got: %v", v, resp)
		}
	}
}

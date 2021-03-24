package models

import (
	"database/sql"
	"testing"

	"github.com/montybeatnik/tutorial_practice/driver"
)

func initializeDevPSQL(t *testing.T) *sql.DB {
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

func TestDeviceDatastore(t *testing.T) {
	db := initializeDevPSQL(t)
	dev := NewDeviceStore(db)
	testDeviceStorer_Create(t, dev)
	testDeviceStorer_Get(t, dev)

}

func testDeviceStorer_Create(t *testing.T, db DeviceStorer) {
	testcases := []struct {
		req      Device
		response Device
	}{
		{
			Device{
				ID:       1,
				Hostname: "juniperOne",
				Loopback: "1.1.1.1",
				Hardware: Hardware{
					ID: 1,
				},
				Software: Software{
					ID: 1,
				},
			},
			Device{
				ID: 1,
				Hardware: Hardware{
					Vendor: "juniper",
					Model:  "ACX1100",
				},
			},
		},
		{Device{
			ID:       2,
			Hostname: "dellOne",
			Loopback: "2.2.2.2",
			Hardware: Hardware{
				ID: 2,
			},
			Software: Software{
				ID: 2,
			},
		},
			Device{
				ID: 2,
				Hardware: Hardware{
					Vendor: "dell",
					Model:  "M1K",
				},
			},
		},
	}
	for _, v := range testcases {
		err := db.Create(v.req)
		if err.Error() == "pq: duplicate key value violates unique constraint \"devices_hostname_key\"" {
			continue
		}
		if err != nil {
			t.Errorf("problem creating test case: %v", err)
		}
	}
}

func testDeviceStorer_Get(t *testing.T, db DeviceStorer) {
	testcases := []struct {
		id   uint
		resp Device
	}{
		{1, Device{ID: 1, Hardware: Hardware{Vendor: "juniper"}}},
		{2, Device{ID: 2, Hardware: Hardware{Vendor: "dell"}}},
	}
	for _, v := range testcases {
		resp, err := db.GetById(v.id)
		if err != nil {
			t.Errorf("problem getting Device:%v", err)
		}

		if resp.ID != v.resp.ID && resp.Vendor != v.resp.Vendor {
			t.Errorf("Expected: %v, Got: %v", v, resp)
		}
	}
}

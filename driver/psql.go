package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PSQLConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	DB       string
}

// ConnectToPSQL takes postgres config, forms the connection string and connects to postgres.
func ConnectToPSQL(p PSQLConfig) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Port, p.Host, p.User, p.Password, p.DB)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

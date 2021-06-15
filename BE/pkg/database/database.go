package database

import (
	"database/sql"
	"fmt"
	"github.com/kiem-toan/cmd/audit-server/config"
)

type Database struct {
	Db   *sql.DB
}

func New(d config.Config) *Database {
	c := d.Databases.Postgres
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
	db, err := sql.Open("postgres",connString)
	if err != nil {
		panic(err)
	}
	return &Database{Db: db}
}
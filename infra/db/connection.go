package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString() string {
	return "user=postgres password=12345678 host=localhost port=5432 dbname=ecommerce"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := getConnectionString()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil

}

package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString(dbConfig *config.DBConfig) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
)
	return connectionString
}

func NewConnection(dbConfig *config.DBConfig) (*sqlx.DB, error) {
	dbSource := getConnectionString(dbConfig)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil

}

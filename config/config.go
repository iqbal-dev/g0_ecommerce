package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}
type Config struct {
	Version               string
	ServiceName           string
	HttpPort              string
	AccessTokenExpireTime string
	DBConfig              DBConfig
}

func loadConfig() {
	godotenv.Load()

	version := os.Getenv("VERSION")
	if version == "" {
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		os.Exit(1)
	}
	accessTokenExpireTime := os.Getenv("ACCESS_TOKEN_EXPIRE_TIME")

	if accessTokenExpireTime == "" {
		accessTokenExpireTime = "1" // default 1 hour
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is Required!")
		os.Exit(1)
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		fmt.Println("DB Pass is Required!")
		os.Exit(1)
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is Required!")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB Port is Required!")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is Required!")
		os.Exit(1)
	}

	configurations = &Config{
		Version:               version,
		ServiceName:           serviceName,
		HttpPort:              ":" + httpPort,
		AccessTokenExpireTime: accessTokenExpireTime,
		DBConfig: DBConfig{
			DBUser: dbUser,
			DBPass: dbPass,
			DBHost: dbHost,
			DBPort: dbPort,
			DBName: dbName,
		},
	}
}

func GetConfig() *Config {

	if configurations == nil {

		loadConfig()
	}
	return configurations

}

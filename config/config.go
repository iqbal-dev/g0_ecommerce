package config

import (
	"os"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

func loadConfig() {
	godotenv.Load()

	version := os.Getenv("VERSION")
	if version == "" {
		version = "dev"
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "ecommerce"
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "3000"
	}
	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    ":" + httpPort,
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations

}

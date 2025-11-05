package config

import (
	"os"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
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
	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    ":" + httpPort,
	}
}

func GetConfig() *Config {

	if configurations == nil {

		loadConfig()
	}
	return configurations

}

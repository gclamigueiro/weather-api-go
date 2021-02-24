package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Loading enviroment variables from .env
	godotenv.Load(".env")
}

type config struct {
	WeatherServer string
	ApiKey        string
}

// single instance of config object
var instance *config

func GetConfig() *config {

	if instance == nil {
		instance = new(config)
		instance.ApiKey = os.Getenv("apiKey")
		instance.WeatherServer = os.Getenv("weatherServer")
	}

	return instance

}

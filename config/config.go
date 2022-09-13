package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	ServiceHost string
	ServicePort int
	LogLevel    string
	HttpPort    string
}

// Load loads environment vars and inflates Config
func Load() Config {

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.ServiceHost = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 9101))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

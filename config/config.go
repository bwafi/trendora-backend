package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config hold the application configuration variable
type Config struct {
	Name string
	Port string
	Env  string
}

// LoadConfig initializes and return the application configuration.
// It retrives enviroment variabel, providing defaults if necessry.
func LoadConfig() *Config {
	return &Config{
		Name: GetEnv("NAME_APPLICATION", ""),
		Port: GetEnv("PORT_APP", "8080"),
	}
}

// init run before the main function to load environment variable.
// It laods .env file in non-production environment.
func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}
}

// GetEnv return the value of a environment variable, or returns
// the default value if the environment variable is not set.
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

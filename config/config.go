package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name string
	Port string
	Env  string
}

func LoadConfig() *Config {
	return &Config{
		Name: GetEnv("NAME_APPLICATION", ""),
		Port: GetEnv("PORT_APP", "8080"),
	}
}

func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}
}

func GetEnv(key, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}

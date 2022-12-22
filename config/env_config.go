package config

import (
	"api-trunfo/models"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() models.EnvConfig {
	godotenv.Load()
	return models.EnvConfig{
		Host:      os.Getenv("host"),
		Port:      os.Getenv("port"),
		User:      os.Getenv("user"),
		Password:  os.Getenv("password"),
		Database:  os.Getenv("database"),
		SecretKey: os.Getenv("secretkey"),
		Issure:    os.Getenv("issure"),
	}
}

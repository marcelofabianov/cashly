package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env string
	Tz  string
	Db  DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env: os.Getenv("ENV"),
		Tz:  os.Getenv("TZ"),
		Db: DatabaseConfig{
			Driver:   os.Getenv("DB_DRIVER"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
	}

	return cfg, nil
}

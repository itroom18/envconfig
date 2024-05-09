package envconfig

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SECRET_KEY    string
	PORT          string
	DB_HOST       string
	DB_PORT       string
	DB_NAME       string
	DB_USER       string
	DB_PASSWORD   string
	SMTP_PORT     string
	SMTP_HOST     string
	SMTP_USER     string
	SMTP_PASSWORD string
}

func GetConfig(config *Config) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config.SECRET_KEY = os.Getenv("SECRET_KEY")
	config.PORT = os.Getenv("PORT")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB_NAME = os.Getenv("DB_NAME")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	config.SMTP_PORT = os.Getenv("SMTP_PORT")
	config.SMTP_HOST = os.Getenv("SMTP_HOST")
	config.SMTP_USER = os.Getenv("SMTP_USER")
	config.SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	return config, nil
}

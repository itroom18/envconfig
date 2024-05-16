package envconfig

import (
	"fmt"
	"os"
	"reflect"

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

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{}

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

func WriteConfig(config Config) error {
	file, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return err
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		fmt.Println("Ошибка очистки файла:", err)
		return err
	}

	fmt.Println("Файл .env успешно очищен.")

	t := reflect.TypeOf(config)

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		value := reflect.ValueOf(config).Field(i).Interface()
		_, err = fmt.Fprintf(file, "%s=%v\n", key.Name, value)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return err
		}

	}
	fmt.Println("Файл .env успешно заполнен.")

	return nil
}

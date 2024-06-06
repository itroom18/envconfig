package envconfig

import (
	"fmt"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		// Server
		SECRET_KEY string
		PORT       string
		// Database
		DB DB
		// SMTP
		SMTP SMTP
		// AWS
		AWS AWS
		// Service
		SERVICE SERVICE
	}

	DB struct {
		HOST     string
		PORT     string
		NAME     string
		USER     string
		PASSWORD string
	}

	SMTP struct {
		PORT     string
		HOST     string
		USER     string
		PASSWORD string
	}

	AWS struct {
		KEY      string
		SECRET   string
		BUCKET   string
		REGION   string
		ENDPOINT string
	}

	SERVICE struct {
		USERS_URI string
	}
)

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	setConfigFields(config)

	return config, nil
}

func setConfigFields(config *Config) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			setNestedConfigFields(fieldValue)
		default:
			envVar := field.Name
			fieldValue.SetString(os.Getenv(envVar))
		}
	}
}

func setNestedConfigFields(v reflect.Value) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envVar := t.Name() + "_" + field.Name
		fieldValue := v.Field(i)
		fieldValue.SetString(os.Getenv(envVar))
	}
}

func WriteConfig(config Config) error {
	file, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return err
	}
	defer file.Close()

	writeConfigFields(file, reflect.ValueOf(config))

	fmt.Println("Файл .env успешно заполнен.")
	return nil
}

func writeConfigFields(file *os.File, v reflect.Value) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			writeNestedConfigFields(file, field.Name, fieldValue)
		default:
			_, err := fmt.Fprintf(file, "%s=%v\n", field.Name, fieldValue.Interface())
			if err != nil {
				fmt.Println("Ошибка записи в файл:", err)
			}
		}
	}
}

func writeNestedConfigFields(file *os.File, prefix string, v reflect.Value) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envVar := prefix + "_" + field.Name
		fieldValue := v.Field(i)
		_, err := fmt.Fprintf(file, "%s=%v\n", envVar, fieldValue.Interface())
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
		}
	}
}

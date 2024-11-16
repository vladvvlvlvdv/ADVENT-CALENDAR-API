package config

import (
	"advent-calendar/pkg/validators"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	PORT   string
	SECRET string
	MODE   string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	ADMIN_EMAIL    string
	ADMIN_PASSWORD string

	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_USER     string
	SMTP_PASSWORD string

	MAIL_URI   string
	CLIENT_URI string
}

var (
	Config      ConfigStruct
	directories = [3]string{
		"public",
		"public/attachments",
		"public/previews",
	}
	Validator = &validators.XValidator{
		Validator: validators.Validate,
	}
)

func LoadConfig() {

	if os.Getenv("MODE") == "test" {
		err := godotenv.Load(".test.env")
		if err != nil {
			log.Fatalf("Ошибка загрузки .test.env файла: %v", err)
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Ошибка загрузки .env файла: %v", err)
		}
	}

	setConfigFields(&Config)
	createPublicDirectories()
}

func setConfigFields(config *ConfigStruct) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		envVar := field.Name
		envValue := os.Getenv(envVar)

		if envValue == "" && envVar != "DB_PASSWORD" {
			log.Fatalf("Ошибка при установке пустого значения %s:", envVar)
		}

		fieldValue.SetString(envValue)
	}
}

func createPublicDirectories() {
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal("Ошибка создания директории")
		}
	}
}

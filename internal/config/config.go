package config

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	MODE        string
}

var Config ConfigStruct

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	setConfigFields(&Config)
}

func setConfigFields(config *ConfigStruct) {
	v := reflect.ValueOf(config).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			setNestedStructFields(fieldValue)
		default:
			envVar := field.Name
			fieldValue.SetString(os.Getenv(envVar))
		}
	}
}

func setNestedStructFields(v reflect.Value) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		envVar := t.Name() + "_" + field.Name
		fieldValue := v.Field(i)
		fieldValue.SetString(os.Getenv(envVar))
	}
}

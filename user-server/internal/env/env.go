package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	err := godotenv.Load()
	return err
}

func GetInt(key string, fallback int) int {

	value := os.Getenv(key)

	valueInt, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return valueInt
}

func GetString(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		ServerPort: getEnv("SERVER_PORT"),
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
		DBSSLMode:  getEnv("DB_SSLMODE"),
	}
}

func getEnv(key string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	log.Fatalf("Ошибка с переменной окружения %s \n", key)
	panic("Ошибка с переменной окружения")
}

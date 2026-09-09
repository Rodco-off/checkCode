package main

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
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Что-то не так с env файлом")
	}

	return &Config{
		ServerPort: os.Getenv("ServerPort"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
	}
}

func main() {
	//...
}

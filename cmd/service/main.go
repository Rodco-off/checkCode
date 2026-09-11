package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type mainConfig struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func loadConfig() *mainConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Что-то не так с env файлом. %v", err)
	}

	return &mainConfig{
		ServerPort: os.Getenv("ServerPort"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBName:     os.Getenv("DBName"),
	}
}

func (config *mainConfig) getConnStr() string {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	return connStr
}

func (config *mainConfig) connectDB() {
	//connStr := config.getConnStr()
	//con, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
}

func main() {
	//config := loadConfig()
}

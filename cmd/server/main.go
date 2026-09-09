package main

import (
	"database/sql"
	"fmt"
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
}

func loadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Что-то не так с env файлом. %v", err)
	}

	return &Config{
		ServerPort: os.Getenv("ServerPort"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBName:     os.Getenv("DBName"),
	}
}

func (config *Config) getConnStr() string {
	connStr := fmt.Sprintf("postgres://%f:%f@%f:%f/%f",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	return connStr
}

func (config *Config) connectDB() *sql.DB {
	connStr := config.getConnStr()
	con, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Ошибка с созданием подключения к БД. %v", err)
	}
	if con.Ping() != nil {
		log.Fatalf("Ошибка с связью с БД. %v", err)
	}

	return con
}

func main() {
	//...
}

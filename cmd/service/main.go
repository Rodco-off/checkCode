package main

import (
	"log"

	"github.com/Rodco-off/checkCode/internal/config"
	"github.com/Rodco-off/checkCode/internal/db"
)

func main() {
	var err error

	cfg := config.Load()
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalln("Ошибка подключения к БД")
		panic("Ошибка подключения к БД")
	}

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalln("Ошибка миграции БД")
		panic("Ошибка миграции БД ")
	}
}

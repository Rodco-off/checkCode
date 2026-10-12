package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rodco-off/checkCode/internal/config"
	"github.com/Rodco-off/checkCode/internal/db"
	"github.com/Rodco-off/checkCode/internal/handler"
	"github.com/Rodco-off/checkCode/repository"
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

	taskRepo := repository.NewTaskRepository(conn)
	submitRepo := repository.NewSubmissionRepository(conn)

	taskHandler := handler.NewTaskHandler(taskRepo)
	submitHandler := handler.NewSubmissionHandler(submitRepo)

	router := handler.NewRouter(*taskHandler, *submitHandler)

	go func() {
		log.Printf("Сервер на %s", cfg.ServerPort)
		if err := router.Start(":" + cfg.ServerPort); err != nil {
			log.Fatalf("Сервер упал с ошибкой %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

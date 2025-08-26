package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

// Запуск сервера
func main() {

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Running server at 8080...")
	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal("Error while server start: ", err)
	}
}

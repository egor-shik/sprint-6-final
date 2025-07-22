package main

import (
	"log"
	"os"

	"github.com/egor-shik/sprint-6-final/server"
)

func main() {
	// Создаем логгер с префиксом и флагами
	logger := log.New(os.Stdout, "MORSE CONVERTER: ", log.LstdFlags|log.Lshortfile)

	// Создаем сервер с нашим логгером
	srv := server.New(logger)

	// Запускаем сервер и проверяем ошибки
	logger.Println("Starting server on :8080")
	if err := srv.Start(); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}

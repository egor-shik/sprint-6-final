package main

import (
	"log"
	"os"
	"sprint-6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "MORSE CONVERTER: ", log.LstdFlags|log.Lshortfile)

	// Пробуем порты по очереди
	// не запускалось на 8080((
	ports := []string{"8080", "8081"}
	for _, port := range ports {
		logger.Printf("Attempting to start server on :%s", port)
		srv := server.New(logger, port)

		if err := srv.Start(); err != nil {
			logger.Printf("Failed to start on port %s: %v", port, err)
			continue
		}

		logger.Printf("Server successfully started on :%s", port)
		return
	}

	logger.Fatal("Failed to start server on all tried ports")
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sprint-6-final/internal/server"
	"syscall"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "MORSE CONVERTER: ", log.LstdFlags|log.Lshortfile)

	srv := server.New(logger, "8080")

	// Канал для graceful shutdown
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Println("Starting server on :8080")
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Could not start server: %v", err)
		}
	}()

	// Ждем сигнала завершения
	<-quit
	logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown: %v", err)
	}

	close(done)
	<-done
	logger.Println("Server stopped")
}

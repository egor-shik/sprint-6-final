package server

import (
	"log"
	"net/http"
	"time"

	"sprint-6-final/internal/handlers"
)

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func New(logger *log.Logger, port string) *Server {
	if port == "" {
		port = "8080"
	}

	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      createRouter(),
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// createRouter создает и настраивает HTTP роутер
func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Регистрируем обработчики
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	return router
}

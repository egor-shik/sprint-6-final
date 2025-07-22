package server

import (
	"log"
	"net/http"
	"time"

	"github.com/egor-shik/sprint-6-final/handlers"
)

// Server структура для хранения сервера и логгера
type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

// New создает и настраивает новый экземпляр сервера
func New(logger *log.Logger) *Server {
	// Создаем роутер
	router := createRouter()

	// Настраиваем HTTP сервер
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

// createRouter создает и настраивает HTTP роутер
func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Регистрируем обработчики
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	return router
}

// Start запускает HTTP сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

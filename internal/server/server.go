package server

import (
	"context"
	"log"
	"net/http"
	"sprint-6-final/internal/handlers"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func New(logger *log.Logger, port string) *Server {
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
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)
	return router
}

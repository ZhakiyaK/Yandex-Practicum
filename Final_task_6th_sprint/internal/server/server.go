package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// NewServer создает сервер на localhost:8080
func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handlers.MainHandle)
	router.HandleFunc("POST /upload", handlers.UploadHandle)

	server := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}

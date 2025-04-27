package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ID-Samurai/go-sprint6-final/internal/handlers"
)

type application struct {
	Server *http.Server
	Logger *log.Logger
}

func CreateServer(log *log.Logger) *application {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.MainHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &application{
		Server: httpServer,
		Logger: log,
	}
}

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server_struct struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(log *log.Logger) *Server_struct {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	server := &http.Server{Addr: ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second}
	return &Server_struct{log, server}
}

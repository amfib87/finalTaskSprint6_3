package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Log        log.Logger
	HttpServer *http.Server
}

func CreateRouter(logInput *log.Logger) *Server {
	myRouter := http.NewServeMux()

	myRouter.HandleFunc("GET /", handlers.HandleRoot)
	myRouter.HandleFunc("POST /upload", handlers.HandleUpload)

	var serv = Server{
		Log: *logInput,
		HttpServer: &http.Server{
			Addr:         ":8080",
			Handler:      myRouter,
			ErrorLog:     logInput,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	return &serv
}

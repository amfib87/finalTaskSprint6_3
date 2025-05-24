package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	mylog := log.New(nil, `main `, log.LstdFlags)

	server := server.CreateRouter(mylog)

	err := server.HttpServer.ListenAndServe()
	if err != nil {
		server.HttpServer.ErrorLog.Fatal(err)
	}

}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ID-Samurai/go-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "SERVER: ", log.LstdFlags)
	myServer := server.CreateServer(logger)

	myServer.Logger.Printf("Server starting on %s\n", myServer.Server.Addr)
	err := myServer.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

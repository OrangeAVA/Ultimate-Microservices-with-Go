package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OrangeAVA/Microservices-with-Go/internal"
)

func main() {
	port := 8080
	engine := internal.BuildEngine()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	log.Printf("Starting server on port %d", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Print("error while running API gin server", "error", err.Error())
		os.Exit(1)
	}
}

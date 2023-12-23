package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	startHTTPServer(srv)
}

func startHTTPServer(srv *http.Server) {
	// Create context that listens for the interrupt signal from the OS.
	sigCtx, stopSig := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSig()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print("error while running API gin server", "error", err.Error())
			os.Exit(1)
		}
	}()

	// Listen for the interrupt signal.
	<-sigCtx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stopSig()

	// Requests are currently on-flight, wait 10 seconds for them to finish.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Print("server forced to shutdown: ", "error", err.Error())
	}
}

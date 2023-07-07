package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// CREATE ROUTER
	log.Println("Creating router...")
	router := createRouter()
	log.Println("Router created successfully")

	// CREATE CONTEXT
	server := &http.Server{
		Addr:              ":8082",
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	// CREATE CHANNEL TO HANDLE OS SIGNALS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// RUN SERVER
	go func() {
		log.Println("Starting server on port 8082...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting or closing listener:: %v", err)
		}
	}()

	// WAIT FOR OS SIGNAL
	<-quit

	// SHUTDOWN SERVER
	log.Println("Shutting down server...")

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Error closing server:: %v", err)
	}

	log.Println("Server stopped gracefully")
	os.Exit(0)
}

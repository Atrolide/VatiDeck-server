// Package main starts the VatiDeck server.
package main

import (
	"context"
	"github.com/Atrolide/VatiDeck-server/cmd/routes"
	"github.com/Atrolide/VatiDeck-server/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO: Add comments for better documentation

const (
	// Port to listen on
	Port = ":8080"

	// Timeouts
	ReadTimeout     = 10 * time.Second
	WriteTimeout    = 10 * time.Second
	ShutdownTimeout = 5 * time.Second
)

func main() {
	// Initialize logger
	log := logger.InitLogger()

	// Log server start
	log.Info("Starting VatiDeck server on port " + Port)

	// Create a custom Gorilla mux (router)
	router := mux.NewRouter()

	// Set up HTTP routes with the custom router
	routes.SetupRoutes(router, log)

	// Define an http.Server with custom settings
	server := &http.Server{
		Addr:         Port,         // Port to listen on
		Handler:      router,       // Attach Gorilla mux (router)
		ReadTimeout:  ReadTimeout,  // Max time to read the request in seconds
		WriteTimeout: WriteTimeout, // Max time to write the response in seconds
	}

	// Channel for graceful shutdown
	sigChan := make(chan os.Signal, 1) // (•ᴗ•)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server
	// Start the server in a goroutine to allow for graceful shutdown
	go func() {
		log.Info("VatiDeck server started on port " + Port)
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("Server closed")
			} else {
				log.Error("Error starting server on port " + Port + ": " + err.Error())
			}
		}
	}()

	// Wait for an interrupt signal
	<-sigChan // (•ᴗ•)
	// Create context with a 5-second timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()

	// Call ShutdownServer to gracefully shut down the server with a timeout context
	ShutdownServer(ctx, server, log)
}

// REVIEW: Move to a separate package?

// ShutdownServer handles graceful shutdown of the server.
func ShutdownServer(ctx context.Context, server *http.Server, log *logger.Logger) {
	log.Info("Shutting down server...")

	// Graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Error("Error during shutdown: " + err.Error())
	} else {
		log.Info("Server gracefully stopped")
	}
}

// Package routes defines all the HTTP routes for the VatiDeck server.
package routes

import (
	"fmt"
	"github.com/Atrolide/VatiDeck-server/pkg/logger"
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: Log requests for paths which do not exist
// TODO: Log requests for incorrect methods

// SetupRoutes initializes all HTTP routes for the VatiDeck server.
//
// Parameters:
// - router: the HTTP request router to register routes with.
// - log: the logger instance used for logging information and errors.
func SetupRoutes(router *mux.Router, log *logger.Logger) {
	// Register the root ("/") route.
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rootHandler(w, r, log)
	}).Methods("GET")

	// Register the "/status" route.
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		statusHandler(w, r, log)
	}).Methods("GET")
}

// rootHandler processes requests to the root ("/") route.
//
// Parameters:
// - w: the ResponseWriter used to send the HTTP response to the client.
// - r: the HTTP request received from the client.
// - log: the logger instance used to log information and errors.
func rootHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger) {
	// Respond with a welcome message.
	if _, err := fmt.Fprintf(w, "Welcome to VatiDeck server!"); err != nil {
		log.Error("Error writing response: " + err.Error())
		return
	}
	// Log the client's remote address.
	log.Info(fmt.Sprintf("Request received from %s", r.RemoteAddr))
}

// statusHandler processes requests to the "/status" route.
//
// Parameters:
// - w: the ResponseWriter used to send the HTTP response to the client.
// - r: the HTTP request received from the client.
// - log: the logger instance used to log information and errors.
func statusHandler(w http.ResponseWriter, r *http.Request, log *logger.Logger) {
	// Respond with a server status message.
	if _, err := fmt.Fprintf(w, "Server is running!"); err != nil {
		log.Error("Error writing response: " + err.Error())
		return
	}
	// Log the client's remote address and the requested URL path.
	log.Info(fmt.Sprintf("Request received from %s for %s", r.RemoteAddr, r.URL.Path))
}

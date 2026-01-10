package main

import (
	"log"
	"net/http"

	"go-base-api/config"
	"go-base-api/handlers"
	"go-base-api/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware)

	// Define routes
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/api/users", handlers.GetUsersHandler).Methods("GET")
	router.HandleFunc("/api/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/api/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	// Start server
	port := cfg.Port
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

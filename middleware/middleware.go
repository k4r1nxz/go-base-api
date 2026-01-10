package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log the request
		log.Printf(
			"%s %s %s",
			r.RemoteAddr,
			r.Method,
			r.URL,
		)
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Log the response
		log.Printf(
			"%s %s %s %v",
			r.RemoteAddr,
			r.Method,
			r.URL,
			time.Since(start),
		)
	})
}

// CORSMiddleware adds CORS headers to responses
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware provides basic authentication (placeholder)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In a real application, you would validate tokens here
		// For now, we'll just log that auth middleware was called
		fmt.Println("Auth middleware called")
		next.ServeHTTP(w, r)
	})
}
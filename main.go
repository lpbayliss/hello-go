package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hello-go/greeting"
)

func main() {
	// Create a new ServeMux (router) to handle HTTP requests
	// ServeMux matches incoming requests to registered patterns
	mux := http.NewServeMux()

	// Register route handlers using Go 1.22+ method-based routing syntax
	// Format: "METHOD /path" - only matches requests with that HTTP method
	// Without method prefix, route would match all HTTP methods
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /health", healthHandler)

	// Get port from environment variable, default to 8080
	// Cloud platforms (Cloud Run, Heroku, etc.) set PORT env var
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configure HTTP server with timeouts to prevent resource exhaustion
	server := &http.Server{
		Addr:         ":" + port,    // Listen on all interfaces (0.0.0.0)
		Handler:      mux,           // Use our router
		ReadTimeout:  5 * time.Second,   // Max time to read request
		WriteTimeout: 10 * time.Second,  // Max time to write response
		IdleTimeout:  120 * time.Second, // Max time for keep-alive connections
	}

	// Start server in goroutine so main thread can handle shutdown signals
	go func() {
		log.Printf("Server starting on port %s", port)
		// ListenAndServe blocks until server stops
		// Returns ErrServerClosed on graceful shutdown (not an error)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Create buffered channel to receive OS signals
	// Buffer size 1 ensures signal isn't missed if we're not ready to receive
	quit := make(chan os.Signal, 1)

	// Register to receive SIGINT (Ctrl+C) and SIGTERM (docker stop, k8s)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal received
	<-quit

	// Graceful shutdown: allow in-flight requests to complete
	log.Println("Shutting down server...")

	// Create context with 30s timeout for shutdown
	// If shutdown takes longer, context cancels and forces immediate stop
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // Release resources if shutdown completes early

	// Shutdown stops accepting new connections and waits for existing to finish
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited")
}

// helloHandler responds to GET /hello with a JSON greeting
// http.ResponseWriter: interface to construct HTTP response
// *http.Request: pointer to request data (method, headers, body, etc.)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header before writing body
	w.Header().Set("Content-Type", "application/json")

	// Get optional name and uppercase query params
	name := r.URL.Query().Get("name")
	uppercase := r.URL.Query().Get("uppercase") == "true"

	// Validate name if provided
	if name != "" {
		if err := greeting.ValidateName(name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}

	// Generate and format greeting using business logic
	msg := greeting.GenerateGreeting(name)
	msg = greeting.FormatGreeting(msg, uppercase)

	// json.NewEncoder writes directly to ResponseWriter (implements io.Writer)
	// Encode converts Go map to JSON and writes it
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

// healthHandler responds to GET /health for container orchestration health checks
// Used by k8s liveness/readiness probes, load balancer health checks, etc.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

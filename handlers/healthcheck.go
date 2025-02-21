package handlers

import (
	"log"
	"net/http"
	"time"

	"go-app/database"
)

// HealthCheckHandler handles the /healthz endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Allow only GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Reject requests with a body (payload)
	if r.ContentLength > 0 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query()) > 0 {
		http.Error(w, "Bad Request - No query parameters allowed", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "" {
		http.Error(w, "Bad Request - No Content-Type header allowed", http.StatusBadRequest)
		return
	}

	// Insert a record into the healthcheck table
	_, err := database.DB.Exec("INSERT INTO webapp.healthcheck (timestamp) VALUES ($1)", time.Now().UTC())
	if err != nil {
		log.Println("🔴 Database Insert Failed:", err) // Log error
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Set response headers
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(http.StatusOK)
}

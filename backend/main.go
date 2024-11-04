package main

import (
	"fmt"
	"io"
	"net/http"
)

// CORS middleware to handle cross-origin requests
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins for testing; you may want to restrict this in production
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed HTTP methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allowed headers

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// handlePing responds with "pong"
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Online")
}

// handleDownload streams a large amount of data to simulate download speed testing
func handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	io.WriteString(w, string(make([]byte, 50*1024*1024))) // Simulate a 50 MB download
}

// handleUpload simulates processing an uploaded file
func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the uploaded file (limiting size to 50 MB)
	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Upload successful")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.HandleFunc("/download", handleDownload)
	mux.HandleFunc("/upload", handleUpload)

	fmt.Println("Starting server on :8080...")
	fmt.Println("Hooray! Server up and running")
	if err := http.ListenAndServe(":8080", enableCORS(mux)); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

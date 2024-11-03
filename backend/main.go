package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

/*Ping endpoint for latency testing*/

func handlPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

/*Download endpoint for download speed testing*/

func handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Length", "50000000") // ~50 MB
	data := make([]byte, 1024*1024)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i] = byte(rand.Intn(256))
	}

	for i := 0; i < 50; i++ {
		_, _ = w.Write(data) // Send 50 MB of data
	}
}

/*Upload endpoint for upload speed testing*/

func handleUpload(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	bytesReceived, _ := io.Copy(io.Discard, r.Body) // Discard uploaded data
	elapsed := time.Since(start)

	uploadSpeed := float64(bytesReceived) / elapsed.Seconds()
	fmt.Fprintf(w, "Upload speed: %f bytes/sec", uploadSpeed)
}

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/download", handleDownload)
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

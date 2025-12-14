package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", requestIDMiddleware(healthCheckHandler))

	http.HandleFunc("/api/data", requestIDMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		requestID := r.Header.Get("X-Request-ID")
		fmt.Fprintf(w, `{"service":"service-a","data":"response from service A","timestamp":"%s","request_id":"%s"}`, 
			time.Now().Format(time.RFC3339), requestID)
	}))

	fmt.Println("Service A running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

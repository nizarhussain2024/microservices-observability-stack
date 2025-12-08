package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"service-a","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"service":"service-a","data":"response from service A","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	fmt.Println("Service A running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

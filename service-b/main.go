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
		fmt.Fprintf(w, `{"status":"healthy","service":"service-b","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/api/process", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"service":"service-b","result":"processed successfully","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	fmt.Println("Service B running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthCheck struct {
	Status    string            `json:"status"`
	Service   string            `json:"service"`
	Timestamp string            `json:"timestamp"`
	Version   string            `json:"version"`
	Uptime    string            `json:"uptime"`
}

var startTime = time.Now()

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()
	
	health := HealthCheck{
		Status:    "healthy",
		Service:   "service-b",
		Timestamp: time.Now().Format(time.RFC3339),
		Version:   "1.0.0",
		Uptime:    uptime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}





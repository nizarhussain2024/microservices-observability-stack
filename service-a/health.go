package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthStatus struct {
	Status    string            `json:"status"`
	Service   string            `json:"service"`
	Timestamp string            `json:"timestamp"`
	Version   string            `json:"version"`
	Uptime    string            `json:"uptime"`
	Checks    map[string]string `json:"checks"`
}

var serviceStartTime = time.Now()

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(serviceStartTime).String()
	
	health := HealthStatus{
		Status:    "healthy",
		Service:   "service-a",
		Timestamp: time.Now().Format(time.RFC3339),
		Version:   "1.0.0",
		Uptime:    uptime,
		Checks: map[string]string{
			"database": "ok",
			"cache":    "ok",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}


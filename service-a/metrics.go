package main

import (
	"log"
	"time"
)

type Metrics struct {
	RequestCount    int64
	ErrorCount      int64
	TotalDuration   time.Duration
	LastRequestTime time.Time
}

var serviceMetrics = &Metrics{}

func recordRequest(duration time.Duration, isError bool) {
	serviceMetrics.RequestCount++
	if isError {
		serviceMetrics.ErrorCount++
	}
	serviceMetrics.TotalDuration += duration
	serviceMetrics.LastRequestTime = time.Now()
}

func logMetrics() {
	log.Printf("Metrics - Requests: %d, Errors: %d, Avg Duration: %v",
		serviceMetrics.RequestCount,
		serviceMetrics.ErrorCount,
		serviceMetrics.TotalDuration/time.Duration(serviceMetrics.RequestCount))
}





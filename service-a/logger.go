package main

import (
	"encoding/json"
	"log"
	"time"
)

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Service   string                 `json:"service"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

func logInfo(message string, fields map[string]interface{}) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     "INFO",
		Service:   "service-a",
		Message:   message,
		Fields:    fields,
	}
	jsonData, _ := json.Marshal(entry)
	log.Println(string(jsonData))
}

func logError(message string, err error, fields map[string]interface{}) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	fields["error"] = err.Error()
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     "ERROR",
		Service:   "service-a",
		Message:   message,
		Fields:    fields,
	}
	jsonData, _ := json.Marshal(entry)
	log.Println(string(jsonData))
}





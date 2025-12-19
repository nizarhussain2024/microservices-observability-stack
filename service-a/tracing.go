package main

import (
	"context"
	"fmt"
	"time"
)

type TraceSpan struct {
	TraceID    string
	SpanID     string
	ParentID   string
	Operation  string
	StartTime  time.Time
	Duration   time.Duration
	Tags       map[string]string
}

var activeSpans = make(map[string]*TraceSpan)

func startSpan(traceID, operation string) *TraceSpan {
	spanID := generateRequestID()
	span := &TraceSpan{
		TraceID:   traceID,
		SpanID:    spanID,
		Operation: operation,
		StartTime: time.Now(),
		Tags:      make(map[string]string),
	}
	activeSpans[spanID] = span
	return span
}

func finishSpan(spanID string) {
	if span, exists := activeSpans[spanID]; exists {
		span.Duration = time.Since(span.StartTime)
		fmt.Printf("Span finished: %s (duration: %v)\n", span.Operation, span.Duration)
		delete(activeSpans, spanID)
	}
}

func addSpanTag(spanID, key, value string) {
	if span, exists := activeSpans[spanID]; exists {
		span.Tags[key] = value
	}
}

func getTraceFromContext(ctx context.Context) string {
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		return traceID
	}
	return ""
}

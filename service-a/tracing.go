package main

import (
	"context"
	"log"
	"time"
)

type TraceContext struct {
	TraceID   string
	SpanID    string
	StartTime time.Time
}

var traceContexts = make(map[string]*TraceContext)

func startTrace(traceID, spanID string) *TraceContext {
	trace := &TraceContext{
		TraceID:   traceID,
		SpanID:    spanID,
		StartTime: time.Now(),
	}
	traceContexts[spanID] = trace
	log.Printf("Trace started: traceID=%s, spanID=%s", traceID, spanID)
	return trace
}

func endTrace(spanID string) {
	if trace, exists := traceContexts[spanID]; exists {
		duration := time.Since(trace.StartTime)
		log.Printf("Trace ended: spanID=%s, duration=%v", spanID, duration)
		delete(traceContexts, spanID)
	}
}

func injectTraceContext(ctx context.Context, traceID, spanID string) context.Context {
	return context.WithValue(ctx, "traceID", traceID)
}


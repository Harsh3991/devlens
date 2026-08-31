package api

import (
 "net/http"
 "os"
 "strings"
)

// NewRouter creates and configures the HTTP router
func NewRouter() http.Handler {
 mux := http.NewServeMux()

// Health check endpoint
 mux.HandleFunc("/health", HealthCheckHandler)

// API v1 routes
 mux.HandleFunc("/api/v1/analyze", AnalyzeHandler)

// Add CORS middleware with environment-driven configuration
 return enableCORS(mux)
}

// enableCORS adds CORS headers to all responses with environment-driven origin configuration
func enableCORS(next http.Handler) http.Handler {
 // Get allowed origins from environment, default to localhost for development
 allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
 if allowedOriginsStr == "" {
 allowedOriginsStr = "http://localhost:3000"
 }

// Parse origins into a map for fast lookup
 allowedOrigins := make(map[string]bool)
 for _, origin := range strings.Split(allowedOriginsStr, ",") {
 allowedOrigins[strings.TrimSpace(origin)] = true
 }

 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 origin := r.Header.Get("Origin")
 
 // Check if origin is allowed
 if allowedOrigins[origin] {
 w.Header().Set("Access-Control-Allow-Origin", origin)
 }

 w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
 w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
 w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

// Handle preflight requests
 if r.Method == "OPTIONS" {
 w.WriteHeader(http.StatusOK)
 return
 }

 next.ServeHTTP(w, r)
 })
}
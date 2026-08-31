package main

import (
 "fmt"
 "log"
 "net/http"
 "os"

 "devlens-backend/internal/api"

 "github.com/joho/godotenv"
)

func main() {
 // Load environment variables from .env file in the backend folder first,
 // then fallback to the repository root for compatibility.
 if err := godotenv.Load(".env"); err != nil {
   if err2 := godotenv.Load("../../.env"); err2 != nil {
     log.Println("Warning: .env file not found, using system environment variables")
   }
 }

// Get port from environment or use default
 port := os.Getenv("PORT")
 if port == "" {
 port = "8080"
 }

// Initialize router
 router := api.NewRouter()

// Start server
 addr := fmt.Sprintf(":%s", port)
 log.Printf("🚀 DevLens Backend Server starting on http://localhost%s", addr)
 log.Printf("📊 Health check available at http://localhost%s/health", addr)

 if err := http.ListenAndServe(addr, router); err != nil {
 log.Fatalf("Failed to start server: %v", err)
 }
}
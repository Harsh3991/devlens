package api

import (
 "encoding/json"
 "fmt"
 "log"
 "net/http"
 "strings"
 "time"

 "devlens-backend/internal/parser"
)

// HealthCheckHandler returns the health status of the server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodGet {
 http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
 return
 }

 response := map[string]interface{}{
 "status": "healthy",
 "service": "devlens-backend",
 "timestamp": time.Now().UTC().Format(time.RFC3339),
 "version": "1.0.0",
 }

 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 json.NewEncoder(w).Encode(response)
}

// AnalyzeRequest represents the incoming request for code analysis
type AnalyzeRequest struct {
 SourceURL string `json:"source_url"`
}

// AnalyzeHandler handles the POST /api/v1/analyze endpoint
func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
 http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
 return
 }

 var req AnalyzeRequest
 if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
 http.Error(w, "Invalid request body", http.StatusBadRequest)
 return
 }

 if req.SourceURL == "" {
 http.Error(w, "source_url is required", http.StatusBadRequest)
 return
 }

 log.Printf("📥 Received analysis request for: %s", req.SourceURL)

// Check if it's a GitHub URL or local path
 var repoPath string
 var cleanup bool

 if strings.HasPrefix(req.SourceURL, "http://") || strings.HasPrefix(req.SourceURL, "https://") {
// Clone the repository
 log.Printf("🔄 Cloning repository: %s", req.SourceURL)

 cloner := parser.NewRepoCloner()
 clonedPath, err := cloner.CloneRepository(req.SourceURL)
 if err != nil {
 log.Printf("❌ Failed to clone repository: %v", err)
 http.Error(w, fmt.Sprintf("Failed to clone repository: %v", err), http.StatusBadRequest)
 return
 }

 repoPath = clonedPath
 cleanup = true

// Schedule cleanup
 defer func() {
 if cleanup {
 log.Printf("🧹 Cleaning up cloned repository: %s", repoPath)
 if err := cloner.CleanupRepository(repoPath); err != nil {
 log.Printf("⚠️ Failed to cleanup repository: %v", err)
 }
 }
 }()

 log.Printf("✅ Repository cloned to: %s", repoPath)
 } else {
// Assume it's a local path
 repoPath = req.SourceURL
 }

// Analyze the repository
 log.Printf("🔍 Analyzing codebase at: %s", repoPath)

 analyzer := parser.NewAnalyzer()
 defer analyzer.Close()

 result, err := analyzer.AnalyzeDirectory(repoPath)
 if err != nil {
 log.Printf("❌ Analysis failed: %v", err)
 http.Error(w, fmt.Sprintf("Analysis failed: %v", err), http.StatusInternalServerError)
 return
 }

 log.Printf("✅ Analysis complete: %d files, %d functions, %d high-risk files",
 result.Summary.TotalFiles,
 result.Summary.TotalFunctions,
 result.Summary.HighRiskFiles)

 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 json.NewEncoder(w).Encode(result)
}
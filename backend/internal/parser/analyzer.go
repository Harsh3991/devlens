package parser

import (
 "fmt"
 "log"
 "path/filepath"
 "strings"
 "sync"

 "devlens-backend/internal/graph"
 "devlens-backend/internal/models"
)

// Analyzer coordinates the parsing and analysis of a codebase
type Analyzer struct {
 parser *ASTParser
 graph *graph.DependencyGraph
 files map[string]*models.FileInfo
 mu sync.RWMutex
}

// NewAnalyzer creates a new analyzer instance
func NewAnalyzer() *Analyzer {
 return &Analyzer{
 parser: NewASTParser(),
 graph: graph.NewDependencyGraph(),
 files: make(map[string]*models.FileInfo),
 }
}

// Close cleans up analyzer resources
func (a *Analyzer) Close() {
 if a.parser != nil {
 a.parser.Close()
 }
}

// AnalyzeDirectory analyzes all supported files in a directory
func (a *Analyzer) AnalyzeDirectory(rootPath string) (*models.AnalysisResult, error) {
// Walk the directory and collect files
 walker := NewFileWalker(rootPath)
 files, err := walker.Walk()
 if err != nil {
 return nil, fmt.Errorf("failed to walk directory: %w", err)
 }

 log.Printf("Found %d files to analyze", len(files))

// Parse files concurrently
 var wg sync.WaitGroup
 semaphore := make(chan struct{}, 10) // Limit concurrent parsing to 10

 for _, filePath := range files {
 wg.Add(1)
 go func(path string) {
 defer wg.Done()
 semaphore <- struct{}{} // Acquire
 defer func() { <-semaphore }() // Release

// Create a new parser for each goroutine to avoid concurrency issues
 parser := NewASTParser()
 defer parser.Close()

 fileInfo, err := parser.ParseFile(path)
 if err != nil {
 log.Printf("Error parsing %s: %v", path, err)
 return
 }

 a.mu.Lock()
 a.files[path] = fileInfo
 a.mu.Unlock()
 }(filePath)
 }

 wg.Wait()
 log.Printf("Parsed %d files successfully", len(a.files))

// Build dependency graph
 a.buildDependencyGraph(rootPath)

// Generate analysis result
 result := a.generateAnalysisResult(rootPath)

 return result, nil
}

// buildDependencyGraph builds the dependency graph from parsed files
func (a *Analyzer) buildDependencyGraph(rootPath string) {
 a.mu.RLock()
 defer a.mu.RUnlock()

 for filePath, fileInfo := range a.files {
 for _, importPath := range fileInfo.Imports {
// Resolve import path to actual file path
 resolvedPath := a.resolveImportPath(filePath, importPath, rootPath)
 if resolvedPath != "" {
 a.graph.AddEdge(filePath, resolvedPath)
 }
 }
 }
}

// resolveImportPath resolves a relative import path to an absolute file path
func (a *Analyzer) resolveImportPath(fromFile, importPath, rootPath string) string {
// Skip node_modules and external packages
 if strings.HasPrefix(importPath, ".") {
// Relative import
 dir := filepath.Dir(fromFile)
 resolved := filepath.Join(dir, importPath)

// Try different extensions
 extensions := []string{"", ".ts", ".tsx", ".js", ".jsx"}
 for _, ext := range extensions {
 testPath := resolved + ext
 if _, exists := a.files[testPath]; exists {
 return testPath
 }
 }

// Try index files
 for _, ext := range []string{".ts", ".tsx", ".js", ".jsx"} {
 testPath := filepath.Join(resolved, "index"+ext)
 if _, exists := a.files[testPath]; exists {
 return testPath
 }
 }
 }

 return ""
}

// generateAnalysisResult generates the final analysis result
func (a *Analyzer) generateAnalysisResult(rootPath string) *models.AnalysisResult {
 a.mu.RLock()
 defer a.mu.RUnlock()

 result := &models.AnalysisResult{
 Repository: rootPath,
 Summary: models.Summary{
 TotalFiles: len(a.files),
 TotalFunctions: 0,
 HighRiskFiles: 0,
 },
 Nodes: []models.Node{},
 Edges: []models.Edge{},
 }

// Generate nodes
 for filePath, fileInfo := range a.files {
// Calculate risk level
 riskLevel := a.calculateRiskLevel(fileInfo)

// Count functions
 result.Summary.TotalFunctions += len(fileInfo.Functions)

// Count high risk files
 if riskLevel == "high" {
 result.Summary.HighRiskFiles++
 }

// Create node
 node := models.Node{
 ID: filePath,
 Type: "file",
 Metrics: models.Metrics{
 Complexity: fileInfo.Complexity,
 FunctionsCount: len(fileInfo.Functions),
 RiskLevel: riskLevel,
 LinesOfCode: fileInfo.LinesOfCode,
 },
 }

 result.Nodes = append(result.Nodes, node)
 }

// Generate edges from dependency graph
 for _, node := range result.Nodes {
 dependencies := a.graph.GetDependencies(node.ID)
 for _, dep := range dependencies {
 edge := models.Edge{
 Source: node.ID,
 Target: dep,
 Type: "import",
 }
 result.Edges = append(result.Edges, edge)
 }
 }

 return result
}

// calculateRiskLevel determines the risk level based on metrics
func (a *Analyzer) calculateRiskLevel(fileInfo *models.FileInfo) string {
 complexity := fileInfo.Complexity
 functionCount := len(fileInfo.Functions)

// High risk: high complexity or many functions
 if complexity > 20 || functionCount > 10 {
 return "high"
 }

// Medium risk: moderate complexity or functions
 if complexity > 10 || functionCount > 5 {
 return "medium"
 }

// Low risk: simple files
 return "low"
}

// GetImpactAnalysis returns files affected by changes to the target file
func (a *Analyzer) GetImpactAnalysis(targetFile string) []string {
 return a.graph.ImpactAnalysis(targetFile)
}

// DetectCycles detects circular dependencies in the codebase
func (a *Analyzer) DetectCycles() [][]string {
 return a.graph.DetectCycles()
}
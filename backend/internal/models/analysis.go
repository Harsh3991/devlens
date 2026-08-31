package models

// AnalysisResult represents the complete analysis output
type AnalysisResult struct {
 Repository string `json:"repository"`
 Summary Summary `json:"summary"`
 Nodes []Node `json:"nodes"`
 Edges []Edge `json:"edges"`
}

// Summary contains high-level statistics about the codebase
type Summary struct {
 TotalFiles int `json:"total_files"`
 TotalFunctions int `json:"total_functions"`
 HighRiskFiles int `json:"high_risk_files"`
}

// Node represents a file in the codebase
type Node struct {
 ID string `json:"id"`
 Type string `json:"type"`
 Metrics Metrics `json:"metrics"`
}

// Metrics contains analysis metrics for a file
type Metrics struct {
 Complexity int `json:"complexity"`
 FunctionsCount int `json:"functions_count"`
 RiskLevel string `json:"risk_level"` // "low", "medium", "high"
 LinesOfCode int `json:"lines_of_code"`
}

// Edge represents a dependency relationship between files
type Edge struct {
 Source string `json:"source"`
 Target string `json:"target"`
 Type string `json:"type"` // "import", "export", etc.
}

// FileInfo represents parsed information about a single file
type FileInfo struct {
 Path string
 Imports []string
 Exports []string
 Functions []FunctionInfo
 Complexity int
 LinesOfCode int
}

// FunctionInfo represents a function declaration
type FunctionInfo struct {
 Name string
 LineStart int
 LineEnd int
 Complexity int
}
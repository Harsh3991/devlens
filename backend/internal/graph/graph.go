package graph

import (
 "sync"
)

// DependencyGraph represents the codebase as a directed graph
type DependencyGraph struct {
// Adjacency list: file -> list of files it imports
 adjacencyList map[string][]string

// Reverse adjacency list: file -> list of files that import it
 reverseList map[string][]string

 mu sync.RWMutex
}

// NewDependencyGraph creates a new dependency graph
func NewDependencyGraph() *DependencyGraph {
 return &DependencyGraph{
 adjacencyList: make(map[string][]string),
 reverseList: make(map[string][]string),
 }
}

// AddEdge adds a dependency relationship (from imports to)
func (g *DependencyGraph) AddEdge(from, to string) {
 g.mu.Lock()
 defer g.mu.Unlock()

// Add to adjacency list
 g.adjacencyList[from] = append(g.adjacencyList[from], to)

// Add to reverse list (for impact analysis)
 g.reverseList[to] = append(g.reverseList[to], from)
}

// GetDependencies returns all files that the given file imports
func (g *DependencyGraph) GetDependencies(file string) []string {
 g.mu.RLock()
 defer g.mu.RUnlock()
 return g.adjacencyList[file]
}

// GetDependents returns all files that import the given file
func (g *DependencyGraph) GetDependents(file string) []string {
 g.mu.RLock()
 defer g.mu.RUnlock()
 return g.reverseList[file]
}

// ImpactAnalysis performs BFS to find all files affected by changes to the target file
func (g *DependencyGraph) ImpactAnalysis(targetFile string) []string {
 g.mu.RLock()
 defer g.mu.RUnlock()

 visited := make(map[string]bool)
 queue := []string{targetFile}
 result := []string{}

 for len(queue) > 0 {
 current := queue[0]
 queue = queue[1:]

 if visited[current] {
 continue
 }

 visited[current] = true
 if current != targetFile {
 result = append(result, current)
 }

// Add all dependents to the queue
 for _, dependent := range g.reverseList[current] {
 if !visited[dependent] {
 queue = append(queue, dependent)
 }
 }
 }

 return result
}

// DetectCycles checks for circular dependencies using DFS
func (g *DependencyGraph) DetectCycles() [][]string {
 g.mu.RLock()
 defer g.mu.RUnlock()

 visited := make(map[string]bool)
 recStack := make(map[string]bool)
 cycles := [][]string{}

 var dfs func(node string, path []string) bool
 dfs = func(node string, path []string) bool {
 visited[node] = true
 recStack[node] = true
 path = append(path, node)

 for _, neighbor := range g.adjacencyList[node] {
 if !visited[neighbor] {
 if dfs(neighbor, path) {
 return true
 }
 } else if recStack[neighbor] {
// Found a cycle
 cycleStart := 0
 for i, n := range path {
 if n == neighbor {
 cycleStart = i
 break
 }
 }
 cycle := append([]string{}, path[cycleStart:]...)
 cycles = append(cycles, cycle)
 }
 }

 recStack[node] = false
 return false
 }

 for node := range g.adjacencyList {
 if !visited[node] {
 dfs(node, []string{})
 }
 }

 return cycles
}

// GetAllNodes returns all files in the graph
func (g *DependencyGraph) GetAllNodes() []string {
 g.mu.RLock()
 defer g.mu.RUnlock()

 nodes := make(map[string]bool)
 for node := range g.adjacencyList {
 nodes[node] = true
 }
 for node := range g.reverseList {
 nodes[node] = true
 }

 result := make([]string, 0, len(nodes))
 for node := range nodes {
 result = append(result, node)
 }
 return result
}
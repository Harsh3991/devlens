#DevLens - Product Requirements Document (PRD)

##1. Project Overview
*DevLens* is a web-based codebase intelligence platform that enables developers to quickly understand unfamiliar software projects. By parsing source code into Abstract Syntax Trees (ASTs), it generates interactive architecture graphs, identifies dependencies, highlights high-risk modules, and performs impact analysis.

*Goal:* Create a zero-cost, highly scalable static analysis tool that demonstrates mastery of advanced data structures (DAGs, ASTs), backend concurrency, and complex frontend visualizations. This project is specifically designed to showcase enterprise-level systems engineering capabilities, moving beyond standard CRUD applications to tackle CPU-bound processing and graph algorithms.

##2. Tech Stack
* *Backend:* Go (Golang) - Chosen for CPU-intensive AST parsing and concurrent file processing.
* *Frontend:* Next.js (React) - Chosen for fast rendering and routing.
* *Visualization:* React Flow - For rendering the interactive node-based architecture map.
* *AST Parser:* Tree-sitter (with Go bindings) - For multi-language code parsing.
* *Database (Optional for Phase 1):* Neon / Supabase (PostgreSQL) - For caching repository scans or user auth.
* *Deployment:* Vercel (Frontend), Fly.io / Render (Go Backend).

---

##3. Core Features & Requirements

###3.1. Codebase Ingestion
* *Feature:* Users can input a public GitHub URL (or upload a ZIP file).
* *Requirement:* The backend must fetch the repository, extract it into memory or a temporary directory, and prepare it for parallel parsing.

###3.2. AST Parsing & Metrics Extraction
* *Feature:* Parse codebase to extract metadata.
* *Requirement:* Use Tree-sitter to parse .ts, .js, (and potentially .go or .py) files.
* *Metrics to Extract:*
 * Imports / Exports (to build the dependency graph).
 * Function declarations and lengths.
 * Cyclomatic Complexity (branching logic count like if, for, switch).

###3.3. Dependency Graphing (DAG)
* *Feature:* Construct a Directed Acyclic Graph representing file relationships.
* *Requirement:* The Go backend will build an Adjacency List in memory. Nodes represent files; Edges represent imports.
* *Analysis:* Identify Circular Dependencies (A -> B -> A).

###3.4. Impact Analysis (The "Blast Radius")
* *Feature:* "What happens if I change this file?"
* *Requirement:* When a user clicks a file, the backend/frontend performs a Depth-First Search (DFS) or Breadth-First Search (BFS) starting from that node, tracing reverse-dependencies (files that import the target file) to output a list of potentially affected modules.

###3.5. Interactive Visualizer UI
* *Feature:* A spatial map of the codebase.
* *Requirement:* React Flow renders nodes. Nodes are color-coded based on risk/complexity (e.g., Red = High Complexity/Many Dependents, Green = Safe/Isolated).

---

##4. System Architecture & API Contracts (For AI Agents)

###4.1. High-Level Flow
1. Client -> POST /api/analyze { repo_url: "..." } -> Go Server
2. Go Server clones repo -> spawns goroutines per file -> Tree-sitter parses ASTs.
3. Go Server aggregates data -> builds DAG -> calculates metrics -> returns JSON.
4. Client consumes JSON -> React Flow renders graph.

###4.2. Main API Endpoint
*POST* /api/v1/analyze
*Payload:*
json
{
 "source_url": "https://github.com/user/repo"
}

*Response Schema (AI target):*
json
{
 "repository": "user/repo",
 "summary": {
 "total_files": 45,
 "total_functions": 120,
 "high_risk_files": 3
 },
 "nodes": [
 {
 "id": "src/auth.ts",
 "type": "file",
 "metrics": {
 "complexity": 14,
 "functions_count": 3,
 "risk_level": "medium"
 }
 }
 ],
 "edges": [
 {
 "source": "src/routes.ts",
 "target": "src/auth.ts",
 "type": "import"
 }
 ]
}


---

##5. Implementation Phases (Step-by-Step)

###Phase 1: Environment Setup & CLI Skeleton (Backend)
1. Initialize Go module: go mod init devlens-backend.
2. Create a basic CLI tool that accepts a local directory path.
3. Walk the directory concurrently using Go's filepath.WalkDir and goroutines to list all target files (e.g., .ts files).

###Phase 2: AST Parsing Engine (The Core)
1. Integrate go-tree-sitter.
2. Write a parser function that takes file content and outputs AST nodes.
3. Traverse the AST to find import_statement (to map edges) and function_declaration (to count functions).
4. Implement a basic cyclomatic complexity counter (counting if, else, switch, for, &&, || nodes).

###Phase 3: Graph Data Structures & Algorithms
1. Create a Graph struct in Go using a map: map[string][]string (Adjacency List).
2. Populate the graph with parsed import data.
3. Write the ImpactAnalysis(targetFile string) function using BFS to return all upstream dependents.
4. Write a cycle detection algorithm to flag circular dependencies.

###Phase 4: API Layer & Next.js Integration
1. Wrap the Go CLI logic in a standard net/http or Gin web server.
2. Initialize Next.js app: npx create-next-app@latest devlens-ui.
3. Install React Flow: npm install reactflow.
4. Create the main dashboard view, connecting the React Flow state to the Go JSON response.
5. Implement the Sidebar: clicking a node updates the sidebar state to show the BFS Impact Analysis and complexity metrics.

###Phase 5: Polish & Deployment
1. Containerize the Go backend using Docker.
2. Deploy backend to Render or Fly.io (Free Tier).
3. Deploy frontend to Vercel.
4. Add sample "Showcase Repositories" on the landing page for 1-click demos so recruiters don't have to wait for a repo to process.

---

##6. Setup & Build Instructions (Local Development)

###Backend (Go)
bash
cd backend
go modtidy
# Run the development server
go runmain.go


###Frontend (Next.js)
bash
cd frontend
npm install
# Run the dev server
npm rundev


##7. Next Steps for AI Implementation
1. *AI Task 1:* Generate the main.go file setting up the HTTP server and the filepath.WalkDir logic.
2. *AI Task 2:* Write the Tree-sitter AST parsing logic for TypeScript.
3. *AI Task 3:* Scaffold the Next.js page containing the React Flow canvas and map the mock JSON to nodes and edges.
#DevLens Backend

Go-based backend server for DevLens codebase intelligence platform.

##Features

- 🚀 High-performance HTTP server
- 📊 AST parsing with Tree-sitter (coming soon)
- 🔍 Dependency graph analysis
- 🎯 Impact analysis (blast radius calculation)
- 🔄 Concurrent file processing with goroutines

##Project Structure

```
backend/
├── cmd/
│ └── server/
│ └── main.go # Entry point
├── internal/
│ ├── api/
│ │ ├── router.go # HTTP router setup
│ │ └── handlers.go # Request handlers
│ ├── parser/
│ │ └── walker.go # File system walker
│ ├── graph/
│ │ └── graph.go # Dependency graph & algorithms
│ └── models/
│ └── analysis.go # Data models
├── pkg/
│ └── utils/ # Utility functions
├── go.mod
└── go.sum
```

##Setup

1. **Install Go** (version 1.21 or higher)
 ```bash
 go version
 ```

2. **Install dependencies**
 ```bash
 cd backend
 go mod tidy
 ```

3. **Configure environment variables**
 - Copy `.env.example` to `.env` in the root directory
 - Fill in your credentials (optional for local development)

##Running the Server

```bash
cd backend
go runcmd/server/main.go
```

The server will start on `http://localhost:8080`

##API Endpoints

###Health Check
```bash
GET /health
```

Response:
```json
{
 "status": "healthy",
 "service": "devlens-backend",
 "timestamp": "2026-08-30T12:00:00Z",
 "version": "1.0.0"
}
```

###Analyze Repository
```bash
POST /api/v1/analyze
Content-Type: application/json

{
 "source_url": "https://github.com/user/repo"
}
```

Response:
```json
{
 "repository": "user/repo",
 "summary": {
 "total_files": 45,
 "total_functions": 120,
 "high_risk_files": 3
 },
 "nodes": [...],
 "edges": [...]
}
```

##Development

###Testing
```bash
go test./...
```

###Build
```bash
go build-odevlens-servercmd/server/main.go
```

###Run binary
```bash
./devlens-server
```

##Next Steps

- [ ] Integrate Tree-sitter for AST parsing
- [ ] Implement GitHub repository cloning
- [ ] Add cyclomatic complexity calculation
- [ ] Implement caching layer
- [ ] Add authentication & rate limiting

##Tech Stack

- **Language:** Go 1.21+
- **Router:** Standard library `net/http`
- **AST Parser:** Tree-sitter (planned)
- **Environment:** godotenv




to run the backend use this command:
```bash
export CGO_ENABLED=1
go run cmd/server/main.go
```
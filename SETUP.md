#DevLens Setup Guide

##✅ Phase 1: Backend Setup - COMPLETED

###What We've Built

1. *Backend Server Structure*
 - Go HTTP server with CORS support
 - Health check endpoint (/health)
 - Analysis endpoint (/api/v1/analyze)
 - Modular architecture with clean separation of concerns

2. *Core Components*
 - *API Layer* (internal/api/): HTTP routing and request handling
 - *Parser* (internal/parser/): File system walker for code scanning
 - *Graph* (internal/graph/): Dependency graph with BFS/DFS algorithms
 - *Models* (internal/models/): Data structures for analysis results

3. *Environment Configuration*
 - .env file with all necessary placeholders
 - .env.example for reference
 - .gitignore to protect sensitive data

###File Structure Created


devlens/
├── backend/
│ ├── cmd/server/main.go # Server entry point
│ ├── internal/
│ │ ├── api/
│ │ │ ├── router.go # HTTP router & CORS
│ │ │ └── handlers.go # Health & Analyze endpoints
│ │ ├── parser/
│ │ │ └── walker.go # Concurrent file walker
│ │ ├── graph/
│ │ │ └── graph.go # DAG + Impact Analysis + Cycle Detection
│ │ └── models/
│ │ └── analysis.go # Data models
│ ├── go.mod # Go dependencies
│ ├── go.sum
│ └── README.md # Backend documentation
├── .env # Environment variables (local)
├── .env.example # Template for .env
├── .gitignore # Git ignore rules
└── README.md # Project PRD



###How to Run the Backend

bash
# Navigate to backend directory
cd backend

# Install dependencies
go modtidy

# Run the server
go runcmd/server/main.go

# Or build and run
go build-odevlens-servercmd/server/main.go
./devlens-server


Server will start on http://localhost:8080

###Testing the API

*Health Check:*
bash
curl http://localhost:8080/health


*Analyze Endpoint:*
bash
curl -XPOSThttp://localhost:8080/api/v1/analyze \
 -H "Content-Type: application/json" \
-d '{"source_url":"https://github.com/user/repo"}'


---

##🔜 Next Steps

###Phase 2: Frontend Setup ✅ COMPLETE
- [x] Initialize Next.js project
- [x] Install React Flow for visualization
- [x] Create dashboard layout
- [x] Connect to backend API
- [x] Implement interactive graph visualization
- [x] Create sidebar for file details
- [x] Add loading states and error handling

###Phase 3: AST Parsing Integration ⏳ NEXT
- [ ] Install Tree-sitter Go bindings
- [ ] Implement TypeScript/JavaScript parser
- [ ] Extract imports/exports
- [ ] Calculate cyclomatic complexity
- [ ] Build dependency graph from parsed data

###Phase 4: AI Integration
- [ ] Add OpenAI/Anthropic SDK
- [ ] Implement code analysis prompts
- [ ] Generate insights and recommendations
- [ ] Add natural language query support

###Phase 5: Advanced Features
- [ ] GitHub repository cloning
- [ ] Database integration (Neon/Supabase)
- [ ] User authentication
- [ ] Caching layer
- [ ] Rate limiting
- [ ] Deployment configuration

---

##📝 Notes

- All credentials are stored in .env (not committed to git)
- The backend is ready to accept requests
- Graph algorithms (BFS, DFS, cycle detection) are implemented
- File walker supports .ts, .tsx, .js, .jsx, .go, .py
- CORS is enabled for frontend integration

---

##🎯 Current Status

*Backend:* ✅ Complete and running on port 8080 
*Frontend:* ✅ Complete and running on port 3000 
*Integration:* ✅ Frontend ↔️ Backend connected 
*AST Parsing:* ⏳ Pending (Phase 3) 
*AI Integration:* ⏳ Pending (Phase 4) 
*Deployment:* ⏳ Pending (Phase 5) 

---

##🎉 Phase 2 Complete!

Both frontend and backend are running and connected. The UI is fully functional with:
- Interactive graph visualization
- File details sidebar
- Loading states and error handling
- Modern, responsive design

*See PHASE_2_COMPLETE.md for detailed information.*

Ready to move to *Phase 3: AST Parsing Integration* when you are! 🚀
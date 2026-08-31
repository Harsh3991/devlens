# DevLens Frontend

Modern Next.js frontend for the DevLens codebase intelligence platform.

## Features

- 🎨 Modern UI with Tailwind CSS
 - 📊 Interactive graph visualization with React Flow
 - 🌓 Dark mode support
 - 📱 Responsive design
 - ⚡ Fast rendering with Next.js 16
 - 🔄 Real-time analysis updates

## Tech Stack

- **Framework:** Next.js 16.3.3 (App Router)
 - **Language:** TypeScript
 - **Styling:** Tailwind CSS
 - **Visualization:** React Flow
 - **HTTP Client:** Axios
 - **Icons:** Lucide React

## Project Structure

```
 frontend/
 ├── app/
 │ ├── layout.tsx # Root layout with metadata
 │ ├── page.tsx # Home page (Dashboard)
 │ └── globals.css # Global styles
 ├── components/
 │ ├── Dashboard.tsx # Main dashboard component
 │ ├── graph/
 │ │ └── CodebaseGraph.tsx # React Flow graph visualization
 │ └── ui/
 │ └── Sidebar.tsx # File details sidebar
 ├── lib/
 │ └── api/
 │ └── client.ts # API client for backend
 ├── types/
 │ └── analysis.ts # TypeScript type definitions
 ├── hooks/ # Custom React hooks (future)
 └── .env.local # Environment variables

```

## Setup

1. **Install dependencies**
 ```bash
 cd frontend
 npm install
 ```

2. **Configure environment variables**
 - The `.env.local` file is already created
 - Default backend URL: `http://localhost:8080`

3. **Run the development server**
 ```bash
 npm run dev
 ```

4. **Open in browser**
 - Navigate to `http://localhost:3000`

## Environment Variables

```env
 NEXT_PUBLIC_API_URL=http://localhost:8080
 NEXT_PUBLIC_APP_NAME=DevLens
 NEXT_PUBLIC_APP_VERSION=1.0.0
 ```

## Components

### Dashboard
 Main application component that handles:
 - Repository URL input
 - Analysis triggering
 - Loading states
 - Error handling
 - Summary statistics display

### CodebaseGraph
 Interactive graph visualization using React Flow:
 - Displays files as nodes
 - Shows dependencies as edges
 - Color-coded by risk level (red=high, orange=medium, green=low)
 - Click nodes to view details
 - Includes minimap and controls

### Sidebar
 File details panel showing:
 - File name and path
 - Risk level badge
 - Cyclomatic complexity
 - Function count
 - Lines of code
 - Impact analysis (placeholder for Phase 3)

## Usage

1. **Start the backend server** (in another terminal)
 ```bash
 cd ../backend
 go run cmd/server/main.go
 ```

2. **Enter a repository URL** in the search bar
 - Example: `https://github.com/user/repo`

3. **Click "Analyze"** to trigger analysis

4. **View the graph** visualization
 - Nodes represent files
 - Edges represent dependencies
 - Colors indicate risk levels

5. **Click on nodes** to view detailed metrics in the sidebar

## Development

### Build for production
 ```bash
 npm run build
 ```

### Start production server
 ```bash
 npm start
 ```

### Lint
 ```bash
 npm run lint
 ```

## Next Steps

- [ ] Add authentication
 - [ ] Implement real-time updates
 - [ ] Add more visualization options
 - [ ] Implement impact analysis UI
 - [ ] Add export functionality
 - [ ] Create landing page
 - [ ] Add sample repositories

## Notes

- Uses Next.js 16 App Router (Server Components by default)
 - Client components marked with `'use client'` directive
 - React Flow requires client-side rendering
 - Tailwind configured with dark mode support
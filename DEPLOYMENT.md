#DevLens - Deployment Guide

**Status**: Production-Ready Deployment Instructions
**Last Updated**: August 31, 2026

---

##📋 Overview

This guide covers deploying DevLens with:
- **Frontend**: Vercel (free tier, auto-deploys from git)
- **Backend**: Fly.io (free tier, fast, no cold starts)
- **Database**: Neon (free PostgreSQL tier)

All configuration is **environment-variable driven** — no code changes needed between local and production.

---

##🏗️ Architecture Overview

```
┌────────────────────────────────────────┐
│ User Browser                           │
└─────────────────┬──────────────────────┘
                  │
                  │ https://devlens.vercel.app
                  ▼
┌────────────────────────────────────────┐
│ Vercel (Frontend - Next.js)            │
│ - Static/Dynamic pages                 │
│ - Environment: NEXT_PUBLIC_API_URL     │
│   = https://api.devlens.fly.dev        │
└─────────────────┬──────────────────────┘
                  │
                  │ CORS: Allowed Origins
                  │ ✓ https://devlens.vercel.app
                  │ ✓ http://localhost:3000 (dev)
                  │
                  ▼
┌────────────────────────────────────────┐
│ Fly.io (Backend - Go)                  │
│ - API: https://api.devlens.fly.dev     │
│ - Environment: ALLOWED_ORIGINS         │
│   = https://devlens.vercel.app         │
│ - Database: Neon PostgreSQL            │
└────────────────────────────────────────┘
```

---

## 1️⃣ Setup Phase 1: Local Development ✅ DONE

Your local setup is already complete:

```bash
# Terminal 1: Backend
cd backend
go run cmd/server/main.go
# http://localhost:8080

# Terminal 2: Frontend
cd frontend
npm run dev
# http://localhost:3000
```

**Environment Files Used**:
- `backend/.env` → `PORT=8080, ALLOWED_ORIGINS=http://localhost:3000`
- `frontend/.env.local` → `NEXT_PUBLIC_API_URL=http://localhost:8080`

---

## 2️⃣ Setup Phase 2: Fly.io Backend Deployment

### Step 1: Install Fly CLI

```bash
# macOS/Linux
curl -L https://fly.io/install.sh | sh

# Windows (use WSL or download from)
# https://github.com/superfly/flyctl/releases
```

### Step 2: Create Fly Account

```bash
flyctl auth signup
# or
flyctl auth login
```

### Step 3: Create Fly App

```bash
cd backend

# Initialize Fly app (interactive)
flyctl launch

# When prompted:
# - App name: devlens-api (or your choice)
# - Region: choose nearest to you
# - Database: No (we'll use external Neon)
# - Deploy: No (we'll configure first)
```

### Step 4: Configure Environment Variables

```bash
# Add production environment variables to Fly
flyctl secrets set \
  PORT=8080 \
  ENVIRONMENT=production \
  ALLOWED_ORIGINS="https://devlens.vercel.app,http://localhost:3000" \
  DATABASE_URL="postgresql://..." \
  GITHUB_TOKEN="your_token" \
  TEMP_DIR="/tmp/devlens" \
  MAX_REPO_SIZE_MB=500
```

**Note**: Replace `postgresql://...` with your Neon connection string.

### Step 5: Deploy

```bash
# Build and deploy
flyctl deploy

# Check status
flyctl status

# View logs
flyctl logs
```

### Step 6: Get Your Backend URL

```bash
flyctl info

# Output will show:
# AppName: devlens-api
# URL: https://devlens-api.fly.dev
```

**Your backend is now live at**: `https://devlens-api.fly.dev`

---

## 3️⃣ Setup Phase 3: Neon Database (Optional for Phase 3+)

If you need to persist data later:

### Create Neon Database

1. Go to [https://console.neon.tech](https://console.neon.tech)
2. Sign up (free tier)
3. Create new project → `devlens`
4. Copy connection string

### Use Connection String in Fly

```bash
flyctl secrets set DATABASE_URL="postgresql://username:password@host/dbname?sslmode=require"
```

---

## 4️⃣ Setup Phase 4: Vercel Frontend Deployment

### Step 1: Push Code to GitHub

```bash
cd /path/to/devlens
git init
git add .
git commit -m "Initial commit: DevLens complete"
git remote add origin https://github.com/YOUR_USERNAME/devlens.git
git branch -M main
git push -u origin main
```

### Step 2: Deploy to Vercel

**Option A: Using Vercel CLI**

```bash
npm i -g vercel
vercel login
cd frontend
vercel --prod
```

**Option B: Using Vercel Dashboard**

1. Go to [https://vercel.com](https://vercel.com)
2. Sign in (GitHub auth recommended)
3. Click "Add New" → "Project"
4. Import your GitHub repo
5. Configure project:
   - Root Directory: `./frontend`
   - Build Command: `npm run build`
   - Output Directory: `.next`

### Step 3: Add Environment Variables

In Vercel Dashboard:
1. Go to Settings → Environment Variables
2. Add variable:
   - **Name**: `NEXT_PUBLIC_API_URL`
   - **Value**: `https://devlens-api.fly.dev`
   - **Environments**: Production, Preview, Development

### Step 4: Redeploy

```bash
# Redeploy to apply new env vars
vercel --prod
```

---

## 5️⃣ Update Backend CORS for Production

Your backend now has environment-driven CORS. Update Fly to allow your Vercel domain:

```bash
cd backend

flyctl secrets set ALLOWED_ORIGINS="https://devlens.vercel.app,http://localhost:3000"
```

This allows:
- ✅ `https://devlens.vercel.app` (your production frontend)
- ✅ `http://localhost:3000` (local development)

---

## 6️⃣ Configuration Reference

### **Development** (Local)

**backend/.env**
```
PORT=8080
ENVIRONMENT=development
ALLOWED_ORIGINS=http://localhost:3000
DATABASE_URL=... (optional, not needed for Phase 1-3)
```

**frontend/.env.local**
```
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_APP_NAME=DevLens
NEXT_PUBLIC_APP_VERSION=1.0.0
```

### **Production** (Fly + Vercel)

**Fly Secrets** (via `flyctl secrets set`)
```
PORT=8080
ENVIRONMENT=production
ALLOWED_ORIGINS=https://devlens.vercel.app,http://localhost:3000
DATABASE_URL=postgresql://... (when needed)
TEMP_DIR=/tmp/devlens
MAX_REPO_SIZE_MB=500
```

**Vercel Environment Variables** (via Dashboard)
```
NEXT_PUBLIC_API_URL=https://devlens-api.fly.dev
NEXT_PUBLIC_APP_NAME=DevLens
NEXT_PUBLIC_APP_VERSION=1.0.0
```

---

## 7️⃣ CORS Security Verification

✅ **Your Setup is Secure**:
- ✓ No wildcard `*` origin (was fixed)
- ✓ Environment-variable driven
- ✓ Localhost allowed for development
- ✓ Production domain explicitly configured
- ✓ No hardcoded URLs in backend code
- ✓ Works for both dev and prod

**Example Request Flow** (Production):

```
1. Browser: GET https://devlens.vercel.app
2. Frontend loads with NEXT_PUBLIC_API_URL=https://devlens-api.fly.dev
3. Frontend sends: POST https://devlens-api.fly.dev/api/v1/analyze
4. Backend receives Origin: https://devlens.vercel.app
5. Backend checks: ALLOWED_ORIGINS contains this origin? YES ✓
6. Backend responds with: Access-Control-Allow-Origin: https://devlens.vercel.app
7. Browser accepts response ✓
```

---

## 8️⃣ Deployment Checklist

### Before Deploying

- [ ] Git repo created and pushed
- [ ] No secrets in `.env` files (use env vars instead)
- [ ] `.env` files are in `.gitignore`
- [ ] Backend builds successfully: `go build ./cmd/server`
- [ ] Frontend builds successfully: `npm run build`
- [ ] Health check endpoint works: `curl localhost:8080/health`
- [ ] API endpoint works locally: `curl -X POST http://localhost:8080/api/v1/analyze ...`

### Fly Deployment

- [ ] Fly CLI installed and authenticated
- [ ] `fly.toml` created in backend directory
- [ ] Secrets configured via `flyctl secrets set`
- [ ] App deployed: `flyctl deploy`
- [ ] Check status: `flyctl status`
- [ ] Check logs: `flyctl logs`
- [ ] Health check works: `curl https://devlens-api.fly.dev/health`

### Vercel Deployment

- [ ] GitHub repo connected
- [ ] Frontend root directory set to `./frontend`
- [ ] Build command: `npm run build`
- [ ] Environment variable set: `NEXT_PUBLIC_API_URL`
- [ ] Deployment successful in Vercel Dashboard
- [ ] Frontend loads: `https://devlens.vercel.app`
- [ ] API communication works (analyze a repo)

### Post-Deployment Testing

- [ ] Health check: `curl https://devlens-api.fly.dev/health`
- [ ] CORS works: Open frontend, check Network tab in DevTools
- [ ] Analyze repo: Enter GitHub URL in frontend
- [ ] No CORS errors in browser console
- [ ] Backend logs show requests: `flyctl logs`

---

## 9️⃣ Monitoring & Logs

### View Backend Logs

```bash
# Stream live logs
flyctl logs --app devlens-api

# View last 100 lines
flyctl logs --app devlens-api -n 100
```

### View Vercel Logs

1. Go to [Vercel Dashboard](https://vercel.com/dashboard)
2. Select your project
3. Click "Deployments" tab
4. Click a deployment → "Logs"

### Monitor Performance

**Fly Metrics**:
```bash
flyctl metrics
flyctl status --detailed
```

**Vercel Analytics**:
- Dashboard → Analytics tab
- Shows response times, bandwidth, etc.

---

## 🔟 Troubleshooting

### Issue: CORS Error in Browser

**Symptoms**: 
```
Access to XMLHttpRequest blocked by CORS policy
```

**Solution**:
1. Check backend logs: `flyctl logs`
2. Verify `ALLOWED_ORIGINS` includes your frontend domain
3. Update secrets: `flyctl secrets set ALLOWED_ORIGINS="..."`
4. Redeploy: `flyctl deploy`

### Issue: 404 Not Found

**Symptoms**: 
```
Backend returns 404 on /health or /api/v1/analyze
```

**Solution**:
1. Check if backend deployed: `flyctl status`
2. Check routes in backend code
3. Verify backend is running: `flyctl logs`

### Issue: Backend Not Starting

**Symptoms**: 
```
App won't start on Fly
```

**Solution**:
1. Check logs: `flyctl logs`
2. Verify Go version compatibility: `go version`
3. Test locally: `go run cmd/server/main.go`
4. Rebuild: `flyctl deploy --rebuild`

### Issue: Database Connection Error

**Symptoms**: 
```
pq: password authentication failed
```

**Solution**:
1. Verify DATABASE_URL is correct
2. Check Neon credentials
3. Update secret: `flyctl secrets set DATABASE_URL="..."`

---

## 1️⃣1️⃣ Cost Summary (Free Tier)

| Service | Cost | Limits |
|---------|------|--------|
| Fly.io | **Free** | 3 shared CPUs, 3GB storage |
| Vercel | **Free** | Unlimited deployments, 6000 build minutes/month |
| Neon (PostgreSQL) | **Free** | 1 project, 3GB storage |
| **Total** | **$0/month** | ✓ Production-ready |

---

## 1️⃣2️⃣ Next Steps

1. **Deploy Backend** → Follow Step 2 above
2. **Deploy Frontend** → Follow Step 4 above
3. **Test End-to-End** → Use deployment checklist
4. **Monitor** → Check logs regularly
5. **Scale** (if needed) → Upgrade paid tiers on Fly/Vercel

---

## 📚 Additional Resources

- [Fly.io Docs](https://fly.io/docs/)
- [Vercel Docs](https://vercel.com/docs)
- [Neon Docs](https://neon.tech/docs)
- [Go Deployment Best Practices](https://golang.org/doc/effective_go)
- [CORS Specification](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)

---

**Deployment Guide Complete! 🚀**

Once deployed, your app will be:
- ✅ Fast (Fly.io has no cold starts)
- ✅ Scalable (free tier, upgrade as needed)
- ✅ Secure (CORS properly configured)
- ✅ Environment-driven (same code, different configs)
- ✅ Zero-cost (on free tier)

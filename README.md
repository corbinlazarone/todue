# Todue

<p align="center">
  <img src="web-client/public/todue-logo.png" alt="Todue Logo" width="80" height="80">
  <h1 align="center"><b>Todue</b></h1>
</p>
<p align="center">AI syllabus to calendar automation tool for students</p>
<!-- <p align="center"> -->
<!--   <a href="https://trytodue.com"><b>Visit Todue →</b></a> -->
<!-- </p> -->
<p align="center">
  <a href="https://go.dev"><img alt="Go" src="https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go" /></a>
  <a href="https://react.dev"><img alt="React" src="https://img.shields.io/badge/React-18-61DAFB?style=flat-square&logo=react" /></a>
</p>

## Features

- **Google Sign-In** — Authentication with your Google account
- **PDF Syllabus Upload** — Upload your course syllabus and let AI extract all assignments and deadlines
- **Smart Extraction** — Currently using OpenCode zen 'big-pickle' free model for fun
- **Review & Edit** — Verify and customize extracted assignments before syncing
- **Calendar Sync** — Push assignments directly to your Google Calendar
- **History** — View all your previously uploaded courses and assignments

## Tech Stack

**Backend**: Go 1.24, httprouter, PostgreSQL 16, goose  
**Frontend**: Next.js 16, React 19, TypeScript, Tailwind CSS 4, shadcn/ui  
**APIs**: OpenCode AI, Google Calendar API, Google OAuth2  

## Architecture

```
┌─────────────┐      ┌─────────────┐      ┌─────────────┐
│   Browser   │ ───▶ │  Go Server  │ ───▶ │ PostgreSQL  │
│ (Next.js)   │      │   (Port 4000)     │  (Port 5432)│
│ (Port 3000) │ ◀────│             │      │             │
└─────────────┘      └─────────────┘      └─────────────┘
```

## Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.24+ (for local backend development)
- Node.js 18+ (for local frontend development)

### Environment Setup

Copy `.env.example` to `.env` and configure:

```bash
# Database
DATABASE_URL=postgres://user:pass@localhost:5432/todue?sslmode=disable

# Google OAuth
GOOGLE_CLIENT_ID=your_client_id
GOOGLE_CLIENT_SECRET=your_client_secret

# JWT
JWT_SECRET=your_jwt_secret

# OpenCode AI
OPENCODE_API_KEY=your_api_key
```

### Running with Docker (Recommended)

```bash
# Start all services
make dev

# Rebuild after code changes
make dev-build

# Stop services
make dev-stop
```

### Running Locally

**Database:**
```bash
make db-up
make migrate-up
```

**Backend:**
```bash
make run
```

**Frontend:**
```bash
cd web-client && npm run dev
```

### Access

- Frontend: http://localhost:3000
- API: http://localhost:4000

## Video

[video-placeholder]

## License

MIT

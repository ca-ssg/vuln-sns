# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a deliberately vulnerable SNS (Social Networking Service) application similar to Twitter/X, created for security learning purposes. It contains intentional security vulnerabilities and should NOT be used in production environments.

## Tech Stack

- **Frontend**: Vue 3 + TypeScript + Quasar UI + Tailwind CSS
- **Backend**: Go with Gin framework
- **Database**: MySQL 8.0
- **Infrastructure**: Docker Compose

## Common Commands

### Frontend Development (in frontend/ directory)
```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Type checking
npm run type-check

# Linting
npm run lint

# Format code
npm run format
```

### Backend Development
The backend is Go-based and runs in Docker. There are no specific Go commands to run locally.

### Docker & Infrastructure
```bash
# Start all services
docker-compose up -d
# OR
make start

# Stop all services
docker-compose down
# OR
make stop

# Rebuild and restart services
make restart

# Connect to MySQL database
docker-compose exec db mysql -uroot -ppassword -Dvuln_app --default-character-set=utf8mb4
# OR
make db

# Reset everything (remove volumes and restart)
make reset

# Clear all containers and volumes
make clear
```

### Service URLs
- Frontend: http://localhost:5173
- Backend API: http://localhost:9090
- API endpoints prefix: /api

### Test Accounts
- alice/alice
- bob/bob
- charlie/charlie

## Architecture Overview

### Frontend Structure
- **src/views/**: Main application pages (Home, Login, Profile)
- **src/components/**: Reusable Vue components (PostCard, PostDialog, BaseDialog)
- **src/stores/**: Pinia state management (auth, posts)
- **src/router/**: Vue Router configuration
- Uses Quasar UI framework for UI components
- Axios for API calls with base URL from VITE_API_URL environment variable

### Backend Structure
- **main.go**: Entry point, sets up Gin router, CORS, and database connection
- **internal/handlers/**: HTTP request handlers
  - auth.go: Login functionality
  - post.go: Post CRUD operations
  - profile.go: User profile management
  - search.go: Hashtag search
- **internal/middleware/**: Authentication middleware
- **internal/models/**: Data models (User, Post)
- **internal/database/**: Database connection management
- Uses raw SQL queries (intentionally vulnerable)

### API Endpoints
Public endpoints:
- GET /api/posts - Get all posts
- POST /api/login - User login
- GET /api/search - Search posts by hashtag
- GET /api/health - Health check

Protected endpoints (require authentication):
- POST /api/posts - Create post
- PUT /api/posts/:id - Update post
- DELETE /api/posts/:id - Delete post
- POST /api/posts/:id/like - Like a post
- DELETE /api/posts/:id/like - Unlike a post
- GET /api/profile - Get user profile
- PUT /api/profile - Update profile
- POST /api/profile/avatar - Upload avatar

### Security Context
This codebase contains intentional vulnerabilities for educational purposes. When working with this code:
- DO NOT fix vulnerabilities unless explicitly asked
- DO NOT add security measures unless requested
- Vulnerabilities are documented in docs/vulnerabilities.md
- This is a learning tool for security education

### Database Schema
The database uses MySQL 8.0 with utf8mb4 character set. Schema is initialized via init.sql during container startup.

## Important Notes
- Frontend uses Vite for development with HMR enabled
- Backend logs all requests in debug mode
- CORS is configured to allow specific origins
- Database connection includes retry logic (30 attempts)
- All services run in Docker containers for isolation
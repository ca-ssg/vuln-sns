# ca-ssg/vuln-sns Repository Knowledge

## Repository Overview
- **Repository**: ca-ssg/vuln-sns (GitHub)
- **Purpose**: Educational vulnerable SNS application for security learning
- **Type**: X (Twitter)-like social networking service with intentional vulnerabilities

## Technology Stack
- **Frontend**: Vue3 + Quasar + TypeScript + Vite
- **Backend**: Go (Gin framework)
- **Database**: MySQL 8.0
- **Orchestration**: Docker Compose
- **Ports**: Frontend (5173), Backend API (9090), Database (3306)

## Application Features
- Post viewing (public access)
- User authentication (alice/alice, bob/bob, charlie/charlie)
- Post creation/editing/deletion (authenticated users)
- Profile editing (nickname changes)
- Avatar upload functionality
- Like/unlike functionality
- Hashtag search functionality

## Project Structure
```
vuln-sns/
├── frontend/          # Vue3 frontend with Quasar
├── backend/           # Go backend with Gin
├── docs/             # Vulnerability documentation
├── docker-compose.yml # Container orchestration
├── init.sql          # Database initialization
└── README.md         # Setup instructions
```

## Database Schema
- **users**: id, password, nickname, avatar_data
- **posts**: id, user_id, content, created_at, updated_at
- **likes**: post_id, user_id (composite primary key)

## Implemented Vulnerabilities

### Authentication Vulnerabilities
1. **SQL Injection in Login** (`docs/auth/sqli-login.md`)
   - Location: `backend/internal/handlers/auth.go`
   - Vulnerable code: Direct string concatenation in SQL query
   - Attack payload: `' OR '1'='1' --` or `alice' --`

2. **Authentication Token Exposure** (`docs/auth/auth-token-exposure.md`)
   - Simple token generation: userID + "_token"
   - Predictable token format

3. **Password Logging** (`docs/auth/password-logging.md`)
   - Passwords may be logged in application logs

### Post Vulnerabilities
1. **SQL Injection in Posts** (`docs/post/sqli-post.md`)
   - Location: `backend/internal/handlers/post.go`
   - Affects: CreatePost, UpdatePost, DeletePost functions
   - Vulnerable code: fmt.Sprintf with user input

2. **XSS in Post Display** (`docs/post/xss-post.md`)
   - Location: `frontend/src/components/PostCard.vue`
   - Vulnerable code: `v-html="post.content"`
   - Attack payload: `<img src=x onerror="alert(1)">`

### Profile Vulnerabilities
1. **XSS in Profile** (`docs/profile/xss-profile.md`)
   - Profile display vulnerabilities

2. **OS Command Injection in Avatar Upload** (`docs/profile/cmdi-avatar.md`)
   - Location: `backend/internal/handlers/profile.go`
   - Vulnerable function: `scanFile()`
   - Attack payload: `avatar.jpg && cat /etc/passwd`

## Setup Commands
```bash
git clone https://github.com/ca-ssg/vuln-sns.git
cd vuln-sns
docker-compose up -d
```

## Access URLs
- Frontend: http://localhost:5173
- Backend API: http://localhost:9090
- Database: `docker-compose exec db mysql -uroot -ppassword -Dvuln_app`

## Documentation Standards
- File naming: `{脆弱性種類}-{機能名}.md`
- Language: Japanese descriptions with English code examples
- Location: `docs/` directory with subdirectories by feature
- Content: Vulnerability description, attack methods, countermeasures with code examples

## Key Files for Analysis
- `backend/internal/handlers/auth.go` - Authentication logic with SQL injection
- `backend/internal/handlers/post.go` - Post operations with SQL injection
- `backend/internal/handlers/profile.go` - Profile operations with command injection
- `frontend/src/components/PostCard.vue` - XSS vulnerability in post display
- `docs/vulnerabilities.md` - Main vulnerability overview
- `init.sql` - Database schema and initial data

## Security Learning Focus
This repository demonstrates common web application vulnerabilities:
- SQL Injection (authentication and CRUD operations)
- Cross-Site Scripting (XSS)
- OS Command Injection
- Insecure authentication token handling
- Information disclosure through logging

## Initial Test Accounts
- alice/alice
- bob/bob  
- charlie/charlie

## Container Environment
- Uses Docker Compose for orchestration
- MySQL with persistent data via init.sql
- Go backend with hot reload capabilities
- Vue frontend with Vite development server

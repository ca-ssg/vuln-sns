# Code Structure

## Frontend Structure (frontend/)
- **src/views/**: Main application pages
  - Home.vue: Main feed page
  - Login.vue: Authentication page
  - Profile.vue: User profile page
- **src/components/**: Reusable Vue components
  - PostCard, PostDialog, BaseDialog, etc.
- **src/stores/**: Pinia state management
  - auth.ts: Authentication state
  - posts.ts: Posts state management
- **src/router/**: Vue Router configuration
- **src/types/**: TypeScript type definitions
- **src/assets/**: Static assets

## Backend Structure (backend/)
- **main.go**: Entry point, router setup, CORS, DB connection
- **internal/handlers/**: HTTP request handlers
  - auth.go: Login functionality
  - post.go: Post CRUD operations
  - profile.go: User profile management
  - search.go: Hashtag search
  - handler.go: Base handler structure
- **internal/middleware/**: 
  - auth.go: Authentication middleware
- **internal/models/**: Data models
  - user.go: User model
  - post.go: Post model
- **internal/database/**: 
  - mysql.go: Database connection management

## Root Structure
- **docker-compose.yml**: Main orchestration file
- **docker-compose.prd.yml**: Production config
- **Makefile**: Convenience commands
- **init.sql**: Database schema initialization
- **my.cnf**: MySQL configuration
- **docs/**: Documentation including vulnerabilities
- **nuclei/**: Security testing templates
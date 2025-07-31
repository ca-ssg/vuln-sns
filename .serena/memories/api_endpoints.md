# API Endpoints

Base URL: http://localhost:9090/api

## Public Endpoints (No Authentication Required)
- **GET /api/posts** - Get all posts
- **POST /api/login** - User login
  - Body: `{"username": "string", "password": "string"}`
  - Returns: `{"token": "string", "user": {...}}`
- **GET /api/search** - Search posts by hashtag
  - Query param: `?tag=<hashtag>`
- **GET /api/health** - Health check endpoint

## Protected Endpoints (Require Authentication Token)
All protected endpoints require the `Authorization` header with the token from login.

### Posts
- **POST /api/posts** - Create a new post
  - Body: `{"content": "string"}`
- **PUT /api/posts/:id** - Update a post
  - Body: `{"content": "string"}`
- **DELETE /api/posts/:id** - Delete a post

### Likes
- **POST /api/posts/:id/like** - Like a post
- **DELETE /api/posts/:id/like** - Unlike a post

### Profile
- **GET /api/profile** - Get current user's profile
- **PUT /api/profile** - Update profile
  - Body: `{"nickname": "string"}`
- **POST /api/profile/avatar** - Upload avatar image
  - Form data with file upload

## Authentication
- Token is returned from login endpoint
- Include in requests as: `Authorization: <token>`
- Token is stored in frontend's Pinia auth store
- No token expiration implemented (intentionally insecure)
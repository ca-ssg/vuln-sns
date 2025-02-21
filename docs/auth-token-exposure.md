# Authentication Token Exposure

This application intentionally includes security vulnerabilities for learning purposes:

## Overview
- Authentication tokens are logged in server logs
- Simple token format (userID_token)
- Tokens stored in plain text in localStorage

## Learning Objectives
1. Understanding security considerations in logging
2. Learning secure token handling
3. Understanding the importance of secure token design

## Implementation Location
- Backend: `backend/internal/middleware/auth.go`
- Frontend: `frontend/src/stores/auth.ts`

## Mitigation Methods
For production environments, the following measures are necessary:

### 1. Disable Token Logging
```go
// Before
log.Printf("Auth header: %s", authHeader)
log.Printf("Invalid token format: %s", authHeader)

// After
log.Printf("Authentication attempt for request: %s %s", c.Request.Method, c.Request.URL.Path)
if err != nil {
    log.Printf("Authentication failed: %v", err)
}
```

### 2. JWT Implementation
```go
// Before
token := userID + "_token"

// After
claims := jwt.MapClaims{
    "user_id": userID,
    "exp":     time.Now().Add(time.Hour * 24).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
```

### 3. Secure Token Storage
```typescript
// Before (frontend/src/stores/auth.ts)
localStorage.setItem('token', token)

// After
// Encrypt token before storage
const encryptToken = (token: string): string => {
    const key = await deriveKey(process.env.VITE_APP_KEY)
    return await encrypt(token, key)
}
const encryptedToken = await encryptToken(token)
sessionStorage.setItem('token', encryptedToken)
```

### 4. Secure Headers
```go
// Before
c.Header("Access-Control-Allow-Headers", "*")

// After
c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
c.Header("X-Content-Type-Options", "nosniff")
c.Header("X-Frame-Options", "DENY")
```

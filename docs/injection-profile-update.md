# SQL Injection (Profile Update)

This application intentionally includes security vulnerabilities for learning purposes:

## Overview
SQL injection vulnerability exists in the profile update function:
```go
query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
```

## Learning Objectives
1. Understanding SQL injection mechanisms
2. Understanding the importance of parameterized queries
3. Learning proper input sanitization methods

## Implementation Location
- `backend/internal/handlers/auth.go` UpdateProfile method

## Mitigation Methods
For production environments, the following measures are necessary:

### 1. Use Parameterized Queries
```go
// Before
query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
_, err := h.db.Exec(query)

// After
query := "UPDATE users SET nickname = ? WHERE id = ?"
_, err := h.db.Exec(query, profile.Nickname, userID)
```

### 2. Input Validation
```go
// Before
if err := c.BindJSON(&profile); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    return
}

// After
if err := c.BindJSON(&profile); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    return
}

// Validate nickname length and characters
if len(profile.Nickname) > 50 || !validateNickname(profile.Nickname) {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nickname format"})
    return
}

func validateNickname(nickname string) bool {
    // Define allowed characters with regex
    pattern := "^[a-zA-Z0-9_-]{1,50}$"
    match, _ := regexp.MatchString(pattern, nickname)
    return match
}
```

### 3. Input Sanitization
```go
// Before
// No sanitization

// After
import (
    "html"
    "strings"
)

func sanitizeInput(input string) string {
    // HTML escape
    escaped := html.EscapeString(input)
    // Remove special characters
    escaped = strings.ReplaceAll(escaped, "'", "")
    escaped = strings.ReplaceAll(escaped, "\"", "")
    return escaped
}

// Usage
profile.Nickname = sanitizeInput(profile.Nickname)
```

### 4. Improved Error Handling
```go
// Before
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
    return
}

// After
if err != nil {
    if strings.Contains(err.Error(), "Duplicate entry") {
        c.JSON(http.StatusConflict, gin.H{"error": "Nickname already exists"})
        return
    }
    log.Printf("Error updating profile: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
    return
}
```

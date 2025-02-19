package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
    return &AuthHandler{
        db: db,
    }
}

func (h *AuthHandler) Login(c *gin.Context) {
    var credentials struct {
        ID       string `json:"id"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Intentionally vulnerable login implementation with SQL injection
    query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND password = SHA2('%s', 256)", 
        credentials.ID, credentials.Password)
    fmt.Printf("Debug: Executing query: %s\n", query)
    
    if h.db == nil {
        fmt.Printf("Error: Database connection is nil\n")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
        return
    }

    // Execute the vulnerable query
    rows, err := h.db.Query(query)
    if err != nil {
        fmt.Printf("Error executing query: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    if rows.Next() {
        var id, nickname string
        if err := rows.Scan(&id, &nickname); err != nil {
            fmt.Printf("Error scanning results: %v\n", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        fmt.Printf("Login successful for user: %s\n", id)
        c.JSON(http.StatusOK, gin.H{
            "token": "dummy-token",  // Intentionally weak token for learning
            "user": gin.H{
                "id": id,
                "nickname": nickname,
            },
        })
        return
    }
    fmt.Printf("Login failed for user: %s\n", credentials.ID)

    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

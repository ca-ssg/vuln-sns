package handlers

import (
    "database/sql"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
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
    log.Printf("Login attempt")
    var credentials struct {
        UserID string `json:"user_id"`
    }

    if err := c.BindJSON(&credentials); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // For testing, accept any user_id and return a simple token
    log.Printf("Login successful for user: %s", credentials.UserID)
    c.JSON(http.StatusOK, gin.H{
        "token": credentials.UserID + "_token",
        "user": models.User{
            ID:       credentials.UserID,
            Nickname: credentials.UserID,
        },
    })
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var profile struct {
        Nickname string `json:"nickname"`
    }

    if err := c.BindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Intentionally vulnerable SQL query
    query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
    log.Printf("Executing query: %s", query)

    _, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

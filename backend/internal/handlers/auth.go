package handlers

import (
    "database/sql"
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
        UserID   string `json:"user_id"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&credentials); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // データベースでユーザーの存在確認とパスワード検証
    var user models.User
    query := "SELECT id, nickname, password FROM users WHERE id = ?"
    err := h.db.QueryRow(query, credentials.UserID).Scan(&user.ID, &user.Nickname, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("User not found: %s", credentials.UserID)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        log.Printf("Database error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // パスワード検証
    // aliceのパスワードはハッシュ化されているため、特別な処理が必要
    if user.ID == "alice" {
        hashedPassword := models.HashPassword(credentials.Password)
        if user.Password != hashedPassword {
            log.Printf("Invalid password for user: %s", credentials.UserID)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
    } else if user.Password != credentials.Password {
        log.Printf("Invalid password for user: %s", credentials.UserID)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // 認証成功、トークン生成
    log.Printf("Login successful for user: %s", credentials.UserID)
    c.JSON(http.StatusOK, gin.H{
        "token": user.ID + "_token",
        "user": models.User{
            ID:       user.ID,
            Nickname: user.Nickname,
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

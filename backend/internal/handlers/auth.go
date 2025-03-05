package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
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
	// 脆弱なSQLクエリ（SQLインジェクションの可能性あり）
	query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND password = SHA2('%s', 256)",
		credentials.UserID, credentials.Password)
	log.Printf("Executing query: %s", query)

	err := h.db.QueryRow(query).Scan(&user.ID, &user.Nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found or invalid password: %s", credentials.UserID)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
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

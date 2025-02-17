package handlers

import (
	"fmt"
	"net/http"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/middleware"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 意図的な脆弱性:
	// 1. SQLインジェクション - 文字列連結によるSQLクエリの構築
	//    攻撃例: ID: alice'; DROP TABLE users; --
	// 2. パスワードの平文保存
	// 3. エラーメッセージの詳細な情報開示
	//
	// Vulnerabilities:
	// 1. SQL Injection through string concatenation
	// 2. Plain text password storage
	// 3. Detailed error message disclosure
	query := fmt.Sprintf("SELECT id, password, nickname FROM users WHERE id = '%s' AND password = '%s'", req.ID, req.Password)
	var user models.User
	err := database.DB.QueryRow(query).Scan(&user.ID, &user.Password, &user.Nickname)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"user":     user,
		"message":  "Login successful",
	})
}

func UpdateNickname(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.UpdateNicknameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 意図的な脆弱性:
	// 1. XSS - ニックネームのエスケープ処理なし
	// 2. SQLインジェクション - 文字列連結によるクエリ構築
	// 3. CSRF - トークンの検証なし
	// 4. アクセス制御の不備 - ユーザー所有権の検証が不十分
	//
	// Vulnerabilities:
	// 1. XSS through unescaped nickname
	// 2. SQL Injection through string concatenation
	// 3. CSRF - no token validation
	// 4. Broken Access Control - insufficient ownership verification
	query := fmt.Sprintf("UPDATE users SET nickname = '%s' WHERE id = '%s'", req.Nickname, userID)
	_, err := database.DB.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nickname"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Nickname updated successfully",
		"nickname": req.Nickname,
	})
}

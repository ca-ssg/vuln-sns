package main

import (
	"log"
	"os"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/handlers"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 意図的な脆弱性を含むアプリケーションであることを警告
	log.Println("⚠️ WARNING: This application contains intentional vulnerabilities for security learning purposes.")
	
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	r := gin.Default()

	// CORS設定 - 意図的に緩い設定
	// Vulnerability: 不適切なCORS設定
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ルーティング設定
	setupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 認証関連
	r.POST("/auth/login", handlers.Login)
	
	// 投稿関連（未認証でも閲覧可能）
	r.GET("/posts", handlers.GetPosts)

	// 認証が必要な機能
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// プロフィール
		auth.PUT("/profile/nickname", handlers.UpdateNickname)

		// 投稿の作成・編集・削除
		auth.POST("/posts", handlers.CreatePost)
		auth.PUT("/posts/:id", handlers.UpdatePost)
		auth.DELETE("/posts/:id", handlers.DeletePost)

		// いいね機能
		auth.POST("/posts/:id/like", handlers.LikePost)
	}
}

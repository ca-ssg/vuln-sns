package main

import (
    "log"
    "os"
    "strings"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/database"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/handlers"
)

func main() {
    r := gin.Default()

    // Get allowed origins from environment variable
    allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
    if allowedOrigins == "" {
        allowedOrigins = "http://localhost:5173"
    }

    // CORS configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     strings.Split(allowedOrigins, ","),
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
        AllowCredentials: true,
        ExposeHeaders:    []string{"Content-Length"},
        MaxAge:           12 * time.Hour,
    }))

    // Initialize database
    err := database.InitDB("root:password@tcp(db:3306)/vuln_app?parseTime=true&multiStatements=true")
    if err != nil {
        log.Fatal(err)
    }

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(database.DB)
    postHandler := handlers.NewPostHandler(database.DB)

    // Routes
    r.POST("/api/login", authHandler.Login)
    r.GET("/api/posts", postHandler.GetPosts)
    r.POST("/api/posts", postHandler.CreatePost)
    r.PUT("/api/posts/:id", postHandler.UpdatePost)
    r.DELETE("/api/posts/:id", postHandler.DeletePost)
    r.POST("/api/posts/:id/like", postHandler.LikePost)
    r.GET("/api/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    log.Fatal(r.Run(":8080"))
}

package main

import (
    "database/sql"
    "log"
    "os"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    _ "github.com/go-sql-driver/mysql"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/database"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/handlers"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/middleware"
)

func main() {
    // Initialize Gin router with debug mode
    gin.SetMode(gin.DebugMode)
    r := gin.Default()
    log.Printf("Initializing Gin router...")

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Get allowed origins from environment variable
    allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
    if allowedOrigins == "" {
        allowedOrigins = "https://twitter-clone-app-56dhic2m.devinapps.com"
    }

    // CORS configuration - Apply before route registration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{allowedOrigins},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
        AllowCredentials: true,
        ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
        MaxAge:           12 * time.Hour,
    }))

    // Initialize database with retries
    var db *sql.DB
    var err error
    maxRetries := 30
    dsn := os.Getenv("MYSQL_DSN")
    if dsn == "" {
        dsn = "root:password@tcp(db:3306)/vuln_app?parseTime=true&multiStatements=true&charset=utf8mb4"
    }
    log.Printf("Connecting to database with DSN: %s", dsn)
    for i := 0; i < maxRetries; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil {
            err = db.Ping()
            if err == nil {
                log.Printf("Successfully connected to database")
                break
            }
        }
        log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
        time.Sleep(time.Second)
    }
    if err != nil {
        log.Fatalf("Failed to connect to database after %d retries: %v", maxRetries, err)
    }
    database.DB = db

    // Seed database with sample data
    if err := database.SeedDatabase(db); err != nil {
        log.Printf("Warning: Failed to seed database: %v", err)
    }

    // Initialize handlers with logging
    log.Printf("Initializing handlers...")
    authHandler := handlers.NewAuthHandler(db)
    postHandler := handlers.NewPostHandler(db)
    searchHandler := handlers.NewSearchHandler(db)
    if searchHandler == nil {
        log.Fatal("Failed to initialize search handler")
    }
    log.Printf("Handlers initialized successfully")

    // Register routes
    log.Printf("Registering routes...")

    // Public routes
    r.GET("/api/search", func(c *gin.Context) {
        log.Printf("Search endpoint hit with URL: %s", c.Request.URL.String())
        searchHandler.SearchByHashtag(c)
    })
    r.POST("/api/login", authHandler.Login)
    r.GET("/api/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Protected routes
    protected := r.Group("/api")
    protected.Use(middleware.Auth())
    {
        protected.GET("/posts", postHandler.GetPosts)
        protected.POST("/posts", postHandler.CreatePost)
        protected.PUT("/posts/:id", postHandler.UpdatePost)
        protected.DELETE("/posts/:id", postHandler.DeletePost)
        protected.POST("/posts/:id/like", postHandler.LikePost)
        protected.DELETE("/posts/:id/like", postHandler.UnlikePost)
        protected.PUT("/profile", authHandler.UpdateProfile)
    }
    log.Printf("Routes registered successfully")
    log.Printf("Starting server on port %s", port)
    log.Fatal(r.Run(":" + port))
}

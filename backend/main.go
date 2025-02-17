package main

import (
	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/handlers"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Initialize database with retry
	var db *database.DB
	var err error
	maxRetries := 30
	for i := 0; i < maxRetries; i++ {
		dsn := os.Getenv("MYSQL_DSN")
		if dsn == "" {
			dsn = "root:password@tcp(db:3306)/vuln_app?parseTime=true&multiStatements=true"
		}

		db, err = database.NewDB(dsn)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Initialize database schema and seed data
	if err := db.InitDB(); err != nil {
		log.Printf("Warning: Failed to initialize database: %v", err)
	}

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)
	postHandler := handlers.NewPostHandler(db)

	// Initialize router
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Public routes
	r.POST("/api/login", authHandler.Login)
	r.POST("/api/register", authHandler.Register)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.Auth())
	{
		// Posts
		protected.GET("/posts", postHandler.GetPosts)
		protected.POST("/posts", postHandler.CreatePost)
		protected.PUT("/posts/:id", postHandler.UpdatePost)
		protected.DELETE("/posts/:id", postHandler.DeletePost)
		protected.POST("/posts/:id/like", postHandler.ToggleLike)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

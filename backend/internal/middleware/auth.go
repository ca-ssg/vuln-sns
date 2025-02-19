package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Printf("Missing Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}
		log.Printf("Auth header: %s", authHeader)

		// Intentionally simple token validation (userID_token format)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("Invalid token format: %s", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		// For testing, accept both "Bearer alice_token" and "Bearer alice"
		userID := parts[1]
		if strings.Contains(parts[1], "_token") {
			tokenParts := strings.Split(parts[1], "_")
			if len(tokenParts) != 2 || tokenParts[1] != "token" {
				log.Printf("Invalid token format: %s", parts[1])
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
				return
			}
			userID = tokenParts[0]
		}

		log.Printf("Setting user_id in context: %s", userID)
		c.Set("user_id", userID)
		c.Next()
	}
}

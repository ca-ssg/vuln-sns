package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	db *database.DB
}

func NewAuthHandler(db *database.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginReq struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Intentionally vulnerable to SQL injection
	user, err := h.db.GetUser(loginReq.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Hash password for comparison
	hash := sha256.Sum256([]byte(loginReq.Password))
	hashedPassword := hex.EncodeToString(hash[:])

	if user.Password != hashedPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token (intentionally simple and insecure)
	token := user.ID + "_token"

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": models.User{
			ID: user.ID,
		},
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerReq struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash password
	hash := sha256.Sum256([]byte(registerReq.Password))
	hashedPassword := hex.EncodeToString(hash[:])

	user := &models.User{
		ID:       registerReq.ID,
		Password: hashedPassword,
	}

	// Intentionally vulnerable to SQL injection
	if err := h.db.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

package handlers

import (
	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type PostHandler struct {
	db *database.DB
}

func NewPostHandler(db *database.DB) *PostHandler {
	return &PostHandler{db: db}
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	posts, err := h.db.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	// Get current user ID from token
	userID := strings.Split(c.GetHeader("Authorization"), "_")[0]

	// Check if current user has liked each post
	for i := range posts {
		var exists bool
		err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", posts[i].ID, userID).Scan(&exists)
		if err == nil && exists {
			posts[i].IsLiked = true
		}
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var postReq struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Get user ID from token
	userID := strings.Split(c.GetHeader("Authorization"), "_")[0]

	post := &models.Post{
		UserID:  userID,
		Content: postReq.Content,
	}

	if err := h.db.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var postReq struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Intentionally vulnerable to SQL injection
	if err := h.db.UpdatePost(id, postReq.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := h.db.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHandler) ToggleLike(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Get user ID from token
	userID := strings.Split(c.GetHeader("Authorization"), "_")[0]

	// Check if like exists
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", id, userID).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check like status"})
		return
	}

	if exists {
		// Remove like
		_, err = h.db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", id, userID)
	} else {
		// Add like
		_, err = h.db.Exec("INSERT INTO likes (post_id, user_id) VALUES (?, ?)", id, userID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to toggle like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like toggled successfully"})
}

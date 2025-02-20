package handlers

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

type PostHandler struct {
    db *sql.DB
}

func NewPostHandler(db *sql.DB) *PostHandler {
    return &PostHandler{
        db: db,
    }
}

func (h *PostHandler) GetPosts(c *gin.Context) {
    userID := c.GetString("user_id")
    rows, err := h.db.Query(`
        SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at,
               (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes,
               EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as is_liked
        FROM posts p
        ORDER BY p.created_at DESC
    `, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.IsLiked)
        if err != nil {
            continue
        }
        posts = append(posts, post)
    }

    c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var post models.Post
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Intentionally vulnerable SQL query
    query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')", userID, post.Content)
    log.Printf("Executing query: %s", query)

    result, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
        return
    }

    id, _ := result.LastInsertId()
    post.ID = id
    post.UserID = userID

    c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    postID := c.Param("id")
    var post models.Post
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Check if post exists and belongs to user
    var postExists bool
    err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ? AND user_id = ?)", postID, userID).Scan(&postExists)
    if err != nil || !postExists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or unauthorized"})
        return
    }

    // Intentionally vulnerable SQL query
    query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %s AND user_id = '%s'", post.Content, postID, userID)
    log.Printf("Executing query: %s", query)

    _, err = h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    postID := c.Param("id")

    // Check if post exists and belongs to user
    var postExists bool
    err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ? AND user_id = ?)", postID, userID).Scan(&postExists)
    if err != nil || !postExists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found or unauthorized"})
        return
    }

    // Intentionally vulnerable SQL query
    query := fmt.Sprintf("DELETE FROM posts WHERE id = %s AND user_id = '%s'", postID, userID)
    log.Printf("Executing query: %s", query)

    _, err = h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHandler) LikePost(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    postID := c.Param("id")

    // Check if post exists
    var postExists bool
    err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM posts WHERE id = ?)", postID).Scan(&postExists)
    if err != nil || !postExists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    // Intentionally vulnerable SQL query
    query := fmt.Sprintf("INSERT INTO likes (user_id, post_id) VALUES ('%s', %s)", userID, postID)
    log.Printf("Executing query: %s", query)

    _, err = h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

func (h *PostHandler) UnlikePost(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    postID := c.Param("id")

    // Check if like exists
    var likeExists bool
    err := h.db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", postID, userID).Scan(&likeExists)
    if err != nil || !likeExists {
        c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
        return
    }

    // Delete like
    _, err = h.db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", postID, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}

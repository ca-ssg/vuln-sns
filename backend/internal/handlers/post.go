package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "strconv"
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
    // Intentionally vulnerable SQL query for learning purposes
    query := `
        SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at,
               (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes_count
        FROM posts p
        ORDER BY p.created_at DESC
    `
    rows, err := h.db.Query(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes)
        if err != nil {
            continue
        }
        posts = append(posts, post)
    }

    c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
    userID, _ := c.Get("user_id")
    var req models.CreatePostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Intentionally vulnerable SQL query for learning purposes
    query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')", userID, req.Content)
    result, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
        return
    }

    id, _ := result.LastInsertId()
    c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Post created successfully"})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
    userID, _ := c.Get("user_id")
    postID := c.Param("id")
    var req models.UpdatePostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Intentionally vulnerable SQL query for learning purposes
    query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %s AND user_id = '%s'", req.Content, postID, userID)
    _, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
    userID, _ := c.Get("user_id")
    postID := c.Param("id")

    // Intentionally vulnerable SQL query for learning purposes
    query := fmt.Sprintf("DELETE FROM posts WHERE id = %s AND user_id = '%s'", postID, userID)
    _, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHandler) LikePost(c *gin.Context) {
    userID, _ := c.Get("user_id")
    postID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    // Intentionally vulnerable SQL query for learning purposes
    var exists bool
    checkQuery := fmt.Sprintf("SELECT 1 FROM likes WHERE post_id = %d AND user_id = '%s'", postID, userID)
    err = h.db.QueryRow(checkQuery).Scan(&exists)
    
    if err != nil {
        // いいねが存在しない場合は追加
        insertQuery := fmt.Sprintf("INSERT INTO likes (post_id, user_id) VALUES (%d, '%s')", postID, userID)
        _, err = h.db.Exec(insertQuery)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
    } else {
        // いいねが存在する場合は削除
        deleteQuery := fmt.Sprintf("DELETE FROM likes WHERE post_id = %d AND user_id = '%s'", postID, userID)
        _, err = h.db.Exec(deleteQuery)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike post"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
    }
}

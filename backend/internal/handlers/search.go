package handlers

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

type SearchHandler struct {
    db *sql.DB
}

func NewSearchHandler(db *sql.DB) *SearchHandler {
    return &SearchHandler{
        db: db,
    }
}

func (h *SearchHandler) SearchByHashtag(c *gin.Context) {
    log.Printf("SearchByHashtag called with context: %+v", c.Request.URL)
    
    hashtag := c.Query("tag")
    if hashtag == "" {
        log.Printf("No hashtag provided in request")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Hashtag parameter is required"})
        return
    }
    log.Printf("Searching for hashtag: %s", hashtag)

    // Intentionally vulnerable SQL query for learning purposes
    query := fmt.Sprintf(`
        SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at,
               (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes
        FROM posts p
        WHERE p.content LIKE '%%%s%%'
        ORDER BY p.created_at DESC
    `, hashtag)

    log.Printf("Executing search query: %s", query)
    rows, err := h.db.Query(query)
    if err != nil {
        log.Printf("Error searching posts: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to search posts: %v", err)})
        return
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes)
        if err != nil {
            log.Printf("Error scanning post: %v", err)
            continue
        }
        posts = append(posts, post)
    }

    c.JSON(http.StatusOK, posts)
}

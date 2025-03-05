package handlers

import (
	"log"
	"net/http"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SearchByHashtag(c *gin.Context) {
	log.Printf("SearchByHashtag called with context: %+v", c.Request.URL)

	hashtag := c.Query("tag")
	if hashtag == "" {
		log.Printf("No hashtag provided in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hashtag parameter is required"})
		return
	}
	log.Printf("Searching for hashtag: %s", hashtag)

	// SQL
	userID := c.GetString("user_id")
	query := `
        SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at,
               (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes,
               EXISTS(SELECT 1 FROM likes WHERE post_id = p.id AND user_id = ?) as is_liked
        FROM posts p
        WHERE p.content LIKE ?
        ORDER BY p.created_at DESC
    `
	hashtagParam := "%" + hashtag + "%"
	stmt, err := h.db.Prepare(query)
	if err != nil {
		log.Printf("プリペアードステートメントの準備エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "検索処理に失敗しました"})
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, hashtagParam)
	if err != nil {
		log.Printf("クエリ実行エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "検索処理に失敗しました"})
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.IsLiked)
		if err != nil {
			log.Printf("Error scanning post: %v", err)
			continue
		}
		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

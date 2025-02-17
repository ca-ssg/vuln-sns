package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/database"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"github.com/gin-gonic/gin"
)

// GetPosts - 投稿一覧を取得
func GetPosts(c *gin.Context) {
	// 意図的な脆弱性:
	// 1. SQLインジェクション - ORDER BYの条件が操作可能
	// 2. XSS - コンテンツのエスケープなし
	// 3. 情報漏洩 - エラー時の詳細な情報開示
	//
	// Vulnerabilities:
	// 1. SQL Injection in ORDER BY clause
	// 2. XSS through unescaped content
	// 3. Information leakage in error messages
	query := `
		SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at,
		       (SELECT COUNT(*) FROM likes WHERE post_id = p.id) as likes_count
		FROM posts p
		ORDER BY p.created_at DESC
	`
	rows, err := database.DB.Query(query)
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

// CreatePost - 新規投稿を作成
func CreatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 意図的な脆弱性: XSSとSQLインジェクション
	// Vulnerability: XSS and SQL Injection
	query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')", userID, req.Content)
	result, err := database.DB.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Post created successfully"})
}

// UpdatePost - 投稿を更新
func UpdatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Param("id")
	var req models.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 意図的な脆弱性: アクセス制御の不備とSQLインジェクション
	// Vulnerability: Broken Access Control and SQL Injection
	query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %s AND user_id = '%s'", req.Content, postID, userID)
	_, err := database.DB.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

// DeletePost - 投稿を削除
func DeletePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Param("id")

	// 意図的な脆弱性: アクセス制御の不備とSQLインジェクション
	// Vulnerability: Broken Access Control and SQL Injection
	query := fmt.Sprintf("DELETE FROM posts WHERE id = %s AND user_id = '%s'", postID, userID)
	_, err := database.DB.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

// LikePost - 投稿にいいねを付ける/解除する
func LikePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 意図的な脆弱性: CSRFとSQLインジェクション
	// Vulnerability: CSRF and SQL Injection
	var exists bool
	checkQuery := fmt.Sprintf("SELECT 1 FROM likes WHERE post_id = %d AND user_id = '%s'", postID, userID)
	err = database.DB.QueryRow(checkQuery).Scan(&exists)
	
	if err != nil {
		// いいねが存在しない場合は追加
		insertQuery := fmt.Sprintf("INSERT INTO likes (post_id, user_id) VALUES (%d, '%s')", postID, userID)
		_, err = database.DB.Exec(insertQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
	} else {
		// いいねが存在する場合は削除
		deleteQuery := fmt.Sprintf("DELETE FROM likes WHERE post_id = %d AND user_id = '%s'", postID, userID)
		_, err = database.DB.Exec(deleteQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike post"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
	}
}

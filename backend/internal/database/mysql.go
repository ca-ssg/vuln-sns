package database

import (
    "database/sql"
    "fmt"
    "time"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
    _ "github.com/go-sql-driver/mysql"
)

// DB is the global database connection
var DB *sql.DB

func InitDB(dsn string) error {
    var err error
    maxRetries := 30
    retryInterval := time.Second

    for i := 0; i < maxRetries; i++ {
        DB, err = sql.Open("mysql", dsn)
        if err != nil {
            time.Sleep(retryInterval)
            continue
        }

        if err = DB.Ping(); err != nil {
            time.Sleep(retryInterval)
            continue
        }

        // Call SeedData after successful connection
        return SeedData()
    }

    return fmt.Errorf("failed to connect to database after %d retries: %v", maxRetries, err)
}

func GetDB() *sql.DB {
    return DB
}

func GetPosts() ([]models.Post, error) {
    // Intentionally vulnerable SQL query for learning purposes
    rows, err := DB.Query("SELECT id, user_id, content, created_at, updated_at, likes FROM posts ORDER BY created_at DESC")
    if err != nil {
        return nil, fmt.Errorf("failed to query posts: %v", err)
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes)
        if err != nil {
            return nil, fmt.Errorf("failed to scan post: %v", err)
        }
        posts = append(posts, post)
    }
    return posts, nil
}

func CreatePost(post *models.Post) error {
    // Intentionally vulnerable SQL query for learning purposes
    query := fmt.Sprintf("INSERT INTO posts (user_id, content, created_at, updated_at, likes) VALUES ('%s', '%s', NOW(), NOW(), 0)",
        post.UserID, post.Content)
    
    result, err := DB.Exec(query)
    if err != nil {
        return fmt.Errorf("failed to create post: %v", err)
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to get last insert id: %v", err)
    }
    
    // Fetch the created post to get all fields
    err = DB.QueryRow("SELECT id, user_id, content, created_at, updated_at, likes FROM posts WHERE id = ?", id).
        Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Likes)
    if err != nil {
        return fmt.Errorf("failed to fetch created post: %v", err)
    }
    
    return nil
}

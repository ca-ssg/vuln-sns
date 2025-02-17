package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
	"log"
	"time"
)

type DB struct {
	*sql.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Check connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) CreateUser(user *models.User) error {
	// Intentionally vulnerable to SQL injection
	query := fmt.Sprintf("INSERT INTO users (id, password) VALUES ('%s', '%s')", user.ID, user.Password)
	_, err := db.Exec(query)
	return err
}

func (db *DB) GetUser(id string) (*models.User, error) {
	// Intentionally vulnerable to SQL injection
	query := fmt.Sprintf("SELECT id, password FROM users WHERE id = '%s'", id)
	row := db.QueryRow(query)

	var user models.User
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) CreatePost(post *models.Post) error {
	// Intentionally vulnerable to SQL injection
	query := fmt.Sprintf(
		"INSERT INTO posts (user_id, content, created_at) VALUES ('%s', '%s', NOW())",
		post.UserID, post.Content,
	)
	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.ID = int(id)

	return nil
}

func (db *DB) GetPosts() ([]models.Post, error) {
	rows, err := db.Query("SELECT id, user_id, content, created_at, (SELECT COUNT(*) FROM likes WHERE post_id = posts.id) as likes FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.Likes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db *DB) UpdatePost(id int, content string) error {
	// Intentionally vulnerable to SQL injection
	query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %d", content, id)
	_, err := db.Exec(query)
	return err
}

func (db *DB) DeletePost(id int) error {
	// No SQL injection here since we're using a parameterized query
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
	return err
}

func (db *DB) ToggleLike(postID int, userID string) error {
	// Check if like exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = ? AND user_id = ?)", postID, userID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// Remove like
		_, err = db.Exec("DELETE FROM likes WHERE post_id = ? AND user_id = ?", postID, userID)
	} else {
		// Add like
		_, err = db.Exec("INSERT INTO likes (post_id, user_id) VALUES (?, ?)", postID, userID)
	}

	return err
}

func (db *DB) InitDB() error {
	// Create tables
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			password VARCHAR(255) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS posts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`,
		`CREATE TABLE IF NOT EXISTS likes (
			post_id INT NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			PRIMARY KEY (post_id, user_id),
			FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	// Seed initial data
	if err := SeedData(db); err != nil {
		log.Printf("Error seeding data: %v", err)
		return err
	}

	return nil
}

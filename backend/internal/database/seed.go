package database

import (
	"fmt"
	"log"
	"time"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

func SeedInitialData() error {
	// Create tables if they don't exist
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			password VARCHAR(255) NOT NULL,
			nickname VARCHAR(255)
		)
	`
	if _, err := DB.Exec(createUsersTable); err != nil {
		log.Printf("Error creating users table: %v", err)
		return err
	}

	createPostsTable := `
		CREATE TABLE IF NOT EXISTS posts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id VARCHAR(255),
			content TEXT,
			created_at DATETIME,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	if _, err := DB.Exec(createPostsTable); err != nil {
		log.Printf("Error creating posts table: %v", err)
		return err
	}

	// Create likes table
	createLikesTable := `
		CREATE TABLE IF NOT EXISTS likes (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id VARCHAR(255),
			post_id INT,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (post_id) REFERENCES posts(id),
			UNIQUE KEY unique_like (user_id, post_id)
		)
	`
	if _, err := DB.Exec(createLikesTable); err != nil {
		log.Printf("Error creating likes table: %v", err)
		return err
	}

	// 初期ユーザーの作成（パスワードはSHA256でハッシュ化）
	hashedPassword := models.HashPassword("alice")
	userStmt, err := DB.Prepare("INSERT INTO users (id, password, nickname) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE password = ?")
	if err != nil {
		log.Printf("Error preparing user statement: %v", err)
		return err
	}
	defer userStmt.Close()

	if _, err := userStmt.Exec("alice", hashedPassword, "Alice", hashedPassword); err != nil {
		log.Printf("Error seeding user: %v", err)
		return err
	}

	// サンプル投稿のコンテンツ
	posts := []string{
		"Hello World! 初めての投稿です。",
		"今日はいい天気ですね！<script>alert('XSS!')</script>",
		"セキュリティの学習って楽しいですよね。",
		"' OR '1'='1' --",
		"明日は晴れるかな？",
		"新しい技術の学習を始めました。",
		"美味しいランチを食べました！",
		"週末の予定を立てています。",
		"プログラミングの練習中です。",
		"今日も一日頑張りましょう！",
	}

	// 各投稿を3回繰り返して、異なるタイムスタンプで作成
	postStmt, err := DB.Prepare("INSERT INTO posts (user_id, content, created_at) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("Error preparing post statement: %v", err)
		return err
	}
	defer postStmt.Close()

	for i := 0; i < 3; i++ {
		for _, content := range posts {
			if _, err := postStmt.Exec(
				"alice",
				content,
				time.Now().Add(-time.Duration(i*24)*time.Hour).Format("2006-01-02 15:04:05"),
			); err != nil {
				log.Printf("Error seeding post: %v", err)
				continue
			}
		}
	}

	return nil
}

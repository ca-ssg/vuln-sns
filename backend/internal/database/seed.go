package database

import (
	"fmt"
	"log"
	"time"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

func SeedInitialData() error {
	// 初期ユーザーの作成（パスワードはSHA256でハッシュ化）
	hashedPassword := models.HashPassword("alice")
	userQuery := fmt.Sprintf("INSERT INTO users (id, password, nickname) VALUES ('alice', '%s', 'Alice') ON DUPLICATE KEY UPDATE password = '%s'", hashedPassword, hashedPassword)
	if _, err := DB.Exec(userQuery); err != nil {
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
	for i := 0; i < 3; i++ {
		for _, content := range posts {
			// 意図的な脆弱性: SQLインジェクション
			// Vulnerability: SQL Injection
			query := fmt.Sprintf(
				"INSERT INTO posts (user_id, content, created_at) VALUES ('alice', '%s', '%s')",
				content,
				time.Now().Add(-time.Duration(i*24)*time.Hour).Format("2006-01-02 15:04:05"),
			)
			
			if _, err := DB.Exec(query); err != nil {
				log.Printf("Error seeding post: %v", err)
				continue
			}
		}
	}

	return nil
}

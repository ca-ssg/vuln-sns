package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	// 意図的な脆弱性: 環境変数から直接文字列連結でDSNを構築
	// Vulnerability: SQLインジェクションの可能性
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// データベースのセットアップ
	if err := setupDatabase(); err != nil {
		return fmt.Errorf("error setting up database: %v", err)
	}

	return nil
}

func setupDatabase() error {
	// ユーザーテーブルの作成
	// 意図的な脆弱性: パスワードをプレーンテキストで保存
	// Vulnerability: 不適切なパスワード保存
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		password VARCHAR(255) NOT NULL,
		nickname VARCHAR(255) NOT NULL
	)`

	if _, err := DB.Exec(createUserTable); err != nil {
		return err
	}

	// 投稿テーブルの作成
	createPostTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT AUTO_INCREMENT PRIMARY KEY,
		user_id VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	if _, err := DB.Exec(createPostTable); err != nil {
		return err
	}

	// いいねテーブルの作成
	createLikeTable := `
	CREATE TABLE IF NOT EXISTS likes (
		post_id INT,
		user_id VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (post_id, user_id),
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`

	if _, err := DB.Exec(createLikeTable); err != nil {
		return err
	}

	// 初期ユーザーの作成
	// 意図的な脆弱性: SQLインジェクションの可能性
	// Vulnerability: 文字列連結によるSQLクエリ
	checkUser := "SELECT id FROM users WHERE id = 'alice'"
	var existingID string
	err := DB.QueryRow(checkUser).Scan(&existingID)
	if err == sql.ErrNoRows {
		insertUser := "INSERT INTO users (id, password, nickname) VALUES ('alice', 'alice', 'Alice')"
		if _, err := DB.Exec(insertUser); err != nil {
			return err
		}
		log.Println("Initial user 'alice' created")
	}

	// 初期データの作成
	if err := SeedInitialData(); err != nil {
		log.Printf("Warning: Failed to seed initial data: %v", err)
	}

	return nil
}

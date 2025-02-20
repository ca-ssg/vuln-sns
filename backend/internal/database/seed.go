package database

import (
    "database/sql"
    "log"
    "time"
)

func SeedDatabase(db *sql.DB) error {
    // Sample users with passwords
    users := []struct {
        ID       string
        Nickname string
        Password string
    }{
        {"alice", "Alice", "alice"},
        {"bob", "Bob", "bob"},
        {"charlie", "Charlie", "charlie"},
    }

    // Sample posts with hashtags
    posts := []struct {
        UserID  string
        Content string
    }{
        {"alice", "Webアプリケーションの #セキュリティ について学んでいます。"},
        {"bob", "新しい #脆弱性 が見つかりました。"},
        {"charlie", "#セキュリティ 対策は重要ですね。"},
        {"alice", "SQLインジェクションは怖い #脆弱性 です。"},
        {"bob", "今日は #セキュリティ の勉強会に参加しました！"},
    }

    // Insert users
    for _, user := range users {
        _, err := db.Exec("INSERT INTO users (id, nickname, password) VALUES (?, ?, ?)", 
            user.ID, user.Nickname, user.Password)
        if err != nil {
            log.Printf("Error inserting user %s: %v", user.ID, err)
            continue
        }
    }

    // Insert posts
    for _, post := range posts {
        _, err := db.Exec(
            "INSERT INTO posts (user_id, content, created_at, updated_at) VALUES (?, ?, ?, ?)",
            post.UserID,
            post.Content,
            time.Now(),
            time.Now(),
        )
        if err != nil {
            log.Printf("Error inserting post for user %s: %v", post.UserID, err)
            continue
        }
    }

    return nil
}

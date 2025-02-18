package database

import (
    "log"
    "time"
)

func SeedData() error {
    if db == nil {
        return nil
    }

    // Create initial user
    _, err := db.Exec("INSERT IGNORE INTO users (id, password, nickname) VALUES (?, SHA2(?, 256), ?)", "alice", "alice", "Alice")
    if err != nil {
        log.Printf("Error seeding user: %v", err)
        return err
    }

    // Create sample posts
    posts := []struct {
        userID  string
        content string
    }{
        {"alice", "セキュリティについて考えています #セキュリティ"},
        {"alice", "脆弱性の学習は大切ですね #脆弱性"},
        {"alice", "今日もコードレビューを頑張ります！"},
    }

    for _, p := range posts {
        now := time.Now()
        _, err := db.Exec(
            "INSERT IGNORE INTO posts (user_id, content, created_at, updated_at, likes) VALUES (?, ?, ?, ?, ?)",
            p.userID, p.content, now, now, 0,
        )
        if err != nil {
            log.Printf("Error seeding post: %v", err)
            return err
        }
    }

    return nil
}

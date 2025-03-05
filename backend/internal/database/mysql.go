package database

import (
    "database/sql"
    "log"
)

var DB *sql.DB

func InitDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    // Create tables
    if err = createTables(db); err != nil {
        return nil, err
    }

    // Seed initial data
    if err = SeedDatabase(db); err != nil {
        log.Printf("Warning: Failed to seed database: %v", err)
    }

    DB = db
    return db, nil
}

func createTables(db *sql.DB) error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS users (
            id VARCHAR(255) PRIMARY KEY,
            nickname VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL,
            avatar_data LONGTEXT DEFAULT NULL
        )`,
        `CREATE TABLE IF NOT EXISTS posts (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            user_id VARCHAR(255) NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
        )`,
        `CREATE TABLE IF NOT EXISTS likes (
            id BIGINT AUTO_INCREMENT PRIMARY KEY,
            user_id VARCHAR(255) NOT NULL,
            post_id BIGINT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id),
            FOREIGN KEY (post_id) REFERENCES posts(id),
            UNIQUE KEY unique_like (user_id, post_id)
        )`,
    }

    for _, query := range queries {
        if _, err := db.Exec(query); err != nil {
            return err
        }
    }

    return nil
}

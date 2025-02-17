package database

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

func SeedData(db *DB) error {
	// Create initial user
	password := "alice"
	hash := sha256.Sum256([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	user := &models.User{
		ID:       "alice",
		Password: hashedPassword,
	}

	if err := db.CreateUser(user); err != nil {
		return err
	}

	// Create sample posts
	posts := []models.Post{
		{
			UserID:  "alice",
			Content: "初めての投稿です！",
		},
		{
			UserID:  "alice",
			Content: "セキュリティの学習頑張ります！",
		},
	}

	for _, post := range posts {
		if err := db.CreatePost(&post); err != nil {
			return err
		}
	}

	return nil
}

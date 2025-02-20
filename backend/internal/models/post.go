package models

import "time"

type Post struct {
    ID        int64     `json:"id" db:"id"`
    UserID    string    `json:"userId" db:"user_id"`
    Content   string    `json:"content" db:"content"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
    Likes     int       `json:"likes" db:"likes"`
    IsLiked   bool      `json:"isLiked" db:"is_liked"`
}

type CreatePostRequest struct {
    Content string `json:"content"`
}

type UpdatePostRequest struct {
    Content string `json:"content"`
}

package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Likes     int       `json:"likes"`
	IsLiked   bool      `json:"isLiked,omitempty"`
}

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

package models

import "time"

type Post struct {
    ID        int64     `json:"id"`
    UserID    string    `json:"userId"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    Likes     int       `json:"likes"`
}

package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"

    "github.com/ca-ssg/devin-vuln-app/internal/database"
    "github.com/ca-ssg/devin-vuln-app/internal/models"
    "github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    post.CreatedAt = time.Now()
    post.UpdatedAt = time.Now()

    result, err := database.DB.Exec(
        "INSERT INTO posts (user_id, content, created_at, updated_at, likes) VALUES (?, ?, ?, ?, ?)",
        post.UserID, post.Content, post.CreatedAt, post.UpdatedAt, 0,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()
    post.ID = id

    json.NewEncoder(w).Encode(post)
}

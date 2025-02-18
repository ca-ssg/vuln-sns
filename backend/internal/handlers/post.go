package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/database"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

func GetPosts(c *gin.Context) {
    posts, err := database.GetPosts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.CreatePost(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, post)
}

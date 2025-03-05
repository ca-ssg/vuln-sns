package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var profile struct {
		Nickname string `json:"nickname"`
	}

	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// プリペアードステートメントを使用してSQLインジェクションを防止
	stmt, err := h.db.Prepare("UPDATE users SET nickname = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベース準備エラー"})
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(profile.Nickname, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

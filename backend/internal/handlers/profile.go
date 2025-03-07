package handlers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ca-ssg/devin-vuln-app/backend/internal/models"
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

func (h *Handler) UploadAvatar(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req models.UploadAvatarRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Base64デコード
	imageData, err := base64.StdEncoding.DecodeString(req.ImageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image data"})
		return
	}

	// 一時ディレクトリの作成（ウィルススキャン用）
	uploadDir := "/tmp/avatars"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary directory"})
		return
	}

	// 一時ファイルの保存（ウィルススキャン用）
	tempFileName := fmt.Sprintf("%d.png", time.Now().Unix())
	filePath := filepath.Join(uploadDir, tempFileName)
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save temporary file"})
		return
	}
	defer os.Remove(filePath) // 処理完了後に一時ファイルを削除

	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（FileID）を適切にエスケープせずにコマンドに渡している
	scanResult, err := scanFile(filePath, req.FileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Virus detected or scan failed", "scan_result": scanResult})
		return
	}

	// データベースに画像データを直接保存
	stmt, err := h.db.Prepare("UPDATE users SET avatar_data = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベース準備エラー"})
		return
	}
	defer stmt.Close()

	// Base64エンコードしたデータをそのまま保存
	_, err = stmt.Exec(req.ImageData, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"malware_scan": scanResult, "avatar_data": req.ImageData})
}

// GetProfile - ユーザープロフィール情報を取得
func (h *Handler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// ユーザー情報を取得
	query := "SELECT id, nickname, avatar_data FROM users WHERE id = ?"
	row := h.db.QueryRow(query, userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Nickname, &user.AvatarData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ウィルススキャン関数（脆弱性あり）
func scanFile(filePath string, fileID string) (string, error) {
	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（fileID）を適切にエスケープせずにコマンドに渡している
	cmd := exec.Command("sh", "-c", "echo 'Scanning file: ' && echo "+fileID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("ウィルスが検出されました: fileID=%s, filePath=%s", fileID, filePath)
	}
	return string(output), nil
}

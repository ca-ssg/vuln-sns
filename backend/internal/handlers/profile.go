package handlers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

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

	// アップロードディレクトリの作成
	uploadDir := "/tmp/avatars"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// ファイルの保存
	filePath := filepath.Join(uploadDir, req.FileID)
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（FileID）を適切にエスケープせずにコマンドに渡している
	if err := scanFile(filePath); err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Virus detected or scan failed"})
		return
	}

	// データベースの更新
	stmt, err := h.db.Prepare("UPDATE users SET avatar_path = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベース準備エラー"})
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(filePath, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully", "avatar_path": filePath})
}

// ウィルススキャン関数（脆弱性あり）
func scanFile(filePath string) error {
	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（filePath）を適切にエスケープせずにコマンドに渡している
	cmd := exec.Command("sh", "-c", "echo 'Scanning file: "+filePath+"' && grep -q 'virus_signature' "+filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			// grepコマンドがパターンを見つけられなかった場合（終了コード1）は正常とみなす
			return nil
		}
		log.Printf("ウィルススキャンエラー: %v, 出力: %s", err, string(output))
		return err
	}
	return fmt.Errorf("ウィルスが検出されました")
}

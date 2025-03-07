package handlers

import (
	"database/sql"
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

	// 一時ディレクトリの作成（ウィルススキャン用）
	uploadDir := "/tmp/avatars"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary directory"})
		return
	}

	// 一時ファイルの保存（ウィルススキャン用）
	filePath := filepath.Join(uploadDir, req.FileID)
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save temporary file"})
		return
	}
	defer os.Remove(filePath) // 処理完了後に一時ファイルを削除

	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（FileID）を適切にエスケープせずにコマンドに渡している
	scanResult, err := scanFile(filePath)
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

	c.JSON(http.StatusOK, gin.H{"scan_result": scanResult, "avatar_data": req.ImageData})
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
	var avatarData sql.NullString
	err := row.Scan(&user.ID, &user.Nickname, &avatarData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	if avatarData.Valid {
		user.AvatarData = avatarData.String
	}

	c.JSON(http.StatusOK, user)
}

// ウィルススキャン関数（脆弱性あり）
func scanFile(filePath string) (string, error) {
	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（filePath）を適切にエスケープせずにコマンドに渡している
	cmd := exec.Command("sh", "-c", "echo 'Scanning file: "+filePath+"' && grep -q 'virus_signature' "+filePath)
	output, err := cmd.CombinedOutput()
	scanResult := string(output)
	
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			// grepコマンドがパターンを見つけられなかった場合（終了コード1）は正常とみなす
			return "OK", nil
		}
		log.Printf("ウィルススキャンエラー: %v, 出力: %s", err, scanResult)
		return scanResult, err
	}
	return scanResult, fmt.Errorf("ウィルスが検出されました")
}

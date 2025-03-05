package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID         string `json:"id"`
	Password   string `json:"password,omitempty"`
	Nickname   string `json:"nickname"`
	AvatarPath string `json:"avatar_path"`
}

type LoginRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateNicknameRequest struct {
	Nickname string `json:"nickname" binding:"required"`
}

type UploadAvatarRequest struct {
	FileID    string `json:"file_id" binding:"required"`
	ImageData string `json:"image_data" binding:"required"`
}

// HashPassword - パスワードをSHA256でハッシュ化
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

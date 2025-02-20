package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID       string `json:"id"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateNicknameRequest struct {
	Nickname string `json:"nickname" binding:"required"`
}

// HashPassword - パスワードをSHA256でハッシュ化
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

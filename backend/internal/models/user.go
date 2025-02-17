package models

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

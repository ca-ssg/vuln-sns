package handlers

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

type AuthHandler struct {
    db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
    return &AuthHandler{
        db: db,
    }
}

func (h *AuthHandler) Login(c *gin.Context) {
    log.Printf("Login attempt")
    var credentials struct {
        UserID   string `json:"user_id"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&credentials); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // データベースでユーザーの存在確認とパスワード検証
    var user models.User
    // 脆弱なSQLクエリ（SQLインジェクションの可能性あり）
    // 意図的に脆弱なクエリ構造を使用（教育目的）
    // 単純な文字列連結によるSQLインジェクションの脆弱性
    
    // 脆弱なSQLクエリ構造（SQLインジェクションに対して脆弱）
    userID := credentials.UserID
    
    // MySQLのコメント構文のために -- の後に空白を追加
    if len(userID) >= 2 && userID[len(userID)-2:] == "--" {
        userID = userID[:len(userID)-2] + "-- "
    }
    
    // 単純な文字列連結によるSQLインジェクションの脆弱性
    // 括弧を使わないことで、OR条件やUNION SELECTによる認証バイパスが可能になる
    // UNION SELECTを使った攻撃を可能にするために、WHERE句を単純化
    query := "SELECT id, nickname FROM users WHERE id = '" + userID + "'"
    log.Printf("Executing query: %s", query)
    
    err := h.db.QueryRow(query).Scan(&user.ID, &user.Nickname)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("User not found or invalid password: %s", credentials.UserID)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        log.Printf("Database error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // 認証成功、トークン生成
    log.Printf("Login successful for user: %s", credentials.UserID)
    c.JSON(http.StatusOK, gin.H{
        "token": user.ID + "_token",
        "user": models.User{
            ID:       user.ID,
            Nickname: user.Nickname,
        },
    })
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
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

    // Intentionally vulnerable SQL query
    query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
    log.Printf("Executing query: %s", query)

    _, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

# アクセス制御の不備

## 脆弱性の説明
アクセス制御の不備により、ユーザーが本来アクセスできないはずのリソースや
機能にアクセスできてしまう脆弱性です。

## 実装箇所
1. 投稿の編集・削除 (`backend/internal/handlers/post.go`)
```go
func (h *PostHandler) UpdatePost(c *gin.Context) {
    // ユーザー所有権の検証が不十分
    postID := c.Param("id")
    // ...
}
```

2. プロフィール更新 (`backend/internal/handlers/auth.go`)
```go
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
    // ユーザー認証の検証が不十分
    userID := c.GetString("user_id")
    // ...
}
```

## 確認手順

### 投稿の不正編集
1. ユーザーAでログイン
2. ユーザーBの投稿IDを取得
3. 以下のようなリクエストを送信:
```bash
curl -X PUT -H "Content-Type: application/json" \
     -H "Authorization: <ユーザーAのトークン>" \
     -d '{"content":"Hacked!"}' \
     http://localhost:9090/api/posts/<ユーザーBの投稿ID>
```
4. ユーザーBの投稿が更新されることを確認

### プロフィールの不正更新
1. ユーザーAでログイン
2. 以下のようなリクエストを送信:
```bash
curl -X PUT -H "Content-Type: application/json" \
     -H "Authorization: <ユーザーAのトークン>" \
     -d '{"id":"user_b","nickname":"Hacked!"}' \
     http://localhost:9090/api/profile
```
3. ユーザーBのプロフィールが更新されることを確認

## 対策方法

### 1. 適切な認可チェックの実装
```go
// ミドルウェアでの認可チェック
func authorizationMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        if userID == "" {
            c.AbortWithStatus(401)
            return
        }

        // リソースへのアクセス権限チェック
        resourceID := c.Param("id")
        if !hasPermission(userID, resourceID) {
            c.AbortWithStatus(403)
            return
        }
        c.Next()
    }
}

func hasPermission(userID, resourceID string) bool {
    // データベースでリソースの所有者を確認
    var ownerID string
    err := db.QueryRow("SELECT user_id FROM posts WHERE id = ?", resourceID).Scan(&ownerID)
    if err != nil {
        return false
    }
    return userID == ownerID
}
```

### 2. リソース所有権の検証
```go
// 投稿の更新処理
func (h *PostHandler) UpdatePost(c *gin.Context) {
    userID := c.GetString("user_id")
    postID := c.Param("id")

    // 投稿の所有者確認
    var post Post
    err := h.db.QueryRow("SELECT id, user_id FROM posts WHERE id = ?", postID).Scan(&post.ID, &post.UserID)
    if err != nil {
        c.JSON(404, gin.H{"error": "Post not found"})
        return
    }

    if post.UserID != userID {
        c.JSON(403, gin.H{"error": "Not authorized to update this post"})
        return
    }

    // 更新処理を実行
    // ...
}
```

### 3. ロールベースのアクセス制御
```go
// ロールの定義
type Role string

const (
    RoleUser  Role = "user"
    RoleAdmin Role = "admin"
)

type UserRole struct {
    UserID string `json:"user_id"`
    Role   Role   `json:"role"`
}

// ロールチェックミドルウェア
func roleCheckMiddleware(requiredRole Role) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.GetString("user_id")
        
        var role Role
        err := db.QueryRow("SELECT role FROM user_roles WHERE user_id = ?", userID).Scan(&role)
        if err != nil || role != requiredRole {
            c.AbortWithStatus(403)
            return
        }
        c.Next()
    }
}

// ルーティングでの使用例
r.PUT("/api/admin/users", roleCheckMiddleware(RoleAdmin), adminHandler.UpdateUsers)
```

### 4. セッション管理の強化
```go
// セッショントークンの検証強化
func validateSession(c *gin.Context) (*Session, error) {
    token := c.GetHeader("Authorization")
    if token == "" {
        return nil, errors.New("no token provided")
    }

    // セッションの検証
    session, err := h.db.GetSession(token)
    if err != nil {
        return nil, err
    }

    // セッションの有効期限チェック
    if session.ExpiresAt.Before(time.Now()) {
        h.db.InvalidateSession(token)
        return nil, errors.New("session expired")
    }

    // IPアドレスの検証
    if session.IP != c.ClientIP() {
        h.db.InvalidateSession(token)
        return nil, errors.New("invalid session")
    }

    return session, nil
}
```

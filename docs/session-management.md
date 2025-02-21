# セッション管理の不備

## 脆弱性の説明
セッション管理の不備により、認証済みセッションの盗用や
不正な再利用が可能になる脆弱性です。

## 実装箇所
1. JWT実装 (`backend/internal/middleware/auth.go`)
```go
// トークンの有効期限が長期間
// トークンの再利用が可能
func (m *AuthMiddleware) ValidateToken(token string) (*Claims, error) {
    // ...
}
```

## 確認手順

### トークンの再利用
1. ユーザーAでログイン
2. ネットワークタブでトークンを確認
3. ログアウト
4. 保存したトークンを使用して以下のリクエストを送信:
```bash
curl -H "Authorization: <保存したトークン>" \
     http://localhost:9090/api/posts
```
5. APIにアクセスできることを確認

### 長期間有効なトークン
1. ユーザーAでログイン
2. トークンを保存
3. 数日後に保存したトークンでアクセス
4. トークンが依然として有効であることを確認

## 対策方法

### 1. トークンの有効期限の適切な設定
```go
// JWTの設定
type Claims struct {
    UserID    string `json:"user_id"`
    ExpiresAt int64  `json:"exp"`
    IssuedAt  int64  `json:"iat"`
}

func generateToken(userID string) (string, error) {
    now := time.Now()
    claims := Claims{
        UserID:    userID,
        ExpiresAt: now.Add(15 * time.Minute).Unix(), // アクセストークンは15分で期限切れ
        IssuedAt:  now.Unix(),
    }
    return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}
```

### 2. リフレッシュトークンの実装
```go
type RefreshToken struct {
    Token     string    `json:"token"`
    UserID    string    `json:"user_id"`
    ExpiresAt time.Time `json:"expires_at"`
}

func generateRefreshToken(userID string) (*RefreshToken, error) {
    token := &RefreshToken{
        Token:     uuid.New().String(),
        UserID:    userID,
        ExpiresAt: time.Now().Add(24 * time.Hour), // リフレッシュトークンは24時間有効
    }
    // データベースに保存
    return token, nil
}

func refreshAccessToken(refreshToken string) (string, error) {
    // リフレッシュトークンの検証
    token, err := validateRefreshToken(refreshToken)
    if err != nil {
        return "", err
    }
    // 新しいアクセストークンの生成
    return generateToken(token.UserID)
}
```

### 3. トークンの無効化機能の実装
```go
// ブラックリストの実装
type TokenBlacklist struct {
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
}

func invalidateToken(token string) error {
    blacklist := &TokenBlacklist{
        Token:     token,
        ExpiresAt: time.Now().Add(24 * time.Hour),
    }
    // データベースに保存
    return nil
}

func isTokenBlacklisted(token string) bool {
    // データベースでトークンの検証
    return false
}
```

### 4. セッションの定期的な更新
```go
// ミドルウェアでのトークン更新チェック
func tokenRefreshMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        claims, err := validateToken(token)
        if err != nil {
            c.AbortWithStatus(401)
            return
        }

        // トークンの有効期限が近い場合は更新
        if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 5*time.Minute {
            newToken, err := generateToken(claims.UserID)
            if err == nil {
                c.Header("X-New-Token", newToken)
            }
        }
        c.Next()
    }
}
```

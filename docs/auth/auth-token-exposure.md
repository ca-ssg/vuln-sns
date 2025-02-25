# 認証トークンの露出

このアプリケーションは学習目的で意図的にセキュリティ脆弱性を含んでいます：

## 概要
- 認証トークンがサーバーログに出力される
- 単純なトークン形式（userID_token）
- トークンがlocalStorageに平文で保存される

## 学習目的
1. ログ出力におけるセキュリティ上の考慮事項の理解
2. 安全なトークン処理の学習
3. 安全なトークン設計の重要性の理解

## 実装箇所
- バックエンド：`backend/internal/middleware/auth.go`
- フロントエンド：`frontend/src/stores/auth.ts`

## 対策方法
本番環境では以下の対策が必要：

### 1. トークンのログ出力の無効化
```go
// 修正前
log.Printf("Auth header: %s", authHeader)
log.Printf("Invalid token format: %s", authHeader)

// 修正後
log.Printf("Authentication attempt for request: %s %s", c.Request.Method, c.Request.URL.Path)
if err != nil {
    log.Printf("Authentication failed: %v", err)
}
```

### 2. JWTの採用
```go
// 修正前
token := userID + "_token"

// 修正後
claims := jwt.MapClaims{
    "user_id": userID,
    "exp":     time.Now().Add(time.Hour * 24).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
```

### 3. トークンの安全な保存
```typescript
// 修正前 (frontend/src/stores/auth.ts)
localStorage.setItem('token', token)

// 修正後
// トークンを暗号化して保存
const encryptToken = (token: string): string => {
    const key = await deriveKey(process.env.VITE_APP_KEY)
    return await encrypt(token, key)
}
const encryptedToken = await encryptToken(token)
sessionStorage.setItem('token', encryptedToken)
```

### 4. セキュアなヘッダーの設定
```go
// 修正前
c.Header("Access-Control-Allow-Headers", "*")

// 修正後
c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
c.Header("X-Content-Type-Options", "nosniff")
c.Header("X-Frame-Options", "DENY")
```

# クロスサイトリクエストフォージェリ（CSRF）

## 脆弱性の説明
CSRFは、認証済みのユーザーに意図しない操作を実行させる脆弱性です。
攻撃者は、ユーザーのブラウザに保存された認証情報を利用して、
ユーザーになりすまして操作を実行できます。

## 実装箇所
1. いいね機能 (`backend/internal/handlers/post.go`)
```go
r.POST("/api/posts/:id/like", postHandler.LikePost)
```

2. プロフィール更新 (`backend/internal/handlers/auth.go`)
```go
r.PUT("/api/profile", authHandler.UpdateProfile)
```

## 確認手順

### いいね機能でのCSRF
1. 以下のようなHTMLを用意:
```html
<form action="http://localhost:9090/api/posts/1/like" method="POST">
  <input type="submit" value="Click me!">
</form>
```
2. 別のWebサイトとしてこのHTMLを開く
3. ログイン済みの状態で、フォームを送信
4. 投稿に「いいね」が追加されることを確認

### プロフィール更新でのCSRF
1. 以下のようなHTMLを用意:
```html
<form action="http://localhost:9090/api/profile" method="PUT">
  <input type="hidden" name="nickname" value="Hacked!">
  <input type="submit" value="Click me!">
</form>
```
2. 別のWebサイトとしてこのHTMLを開く
3. ログイン済みの状態で、フォームを送信
4. プロフィールが更新されることを確認

## 対策方法

### 1. CSRFトークンの実装
```go
// バックエンド（Go）
func generateCSRFToken(c *gin.Context) string {
    token := uuid.New().String()
    c.SetCookie("csrf_token", token, 3600, "/", "", true, true)
    return token
}

func validateCSRFToken(c *gin.Context) bool {
    cookie, err := c.Cookie("csrf_token")
    if err != nil {
        return false
    }
    token := c.GetHeader("X-CSRF-Token")
    return cookie == token
}

// ミドルウェアの実装
func CSRFMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != "GET" {
            if !validateCSRFToken(c) {
                c.AbortWithStatus(403)
                return
            }
        }
        c.Next()
    }
}
```

```javascript
// フロントエンド（Vue）
// APIリクエスト時にCSRFトークンを設定
axios.interceptors.request.use(config => {
  const token = getCookie('csrf_token');
  if (token) {
    config.headers['X-CSRF-Token'] = token;
  }
  return config;
});
```

### 2. SameSite属性の設定
```go
// セッションCookieの設定
c.SetCookie("session", sessionID, 3600, "/", "", true, true, http.SameSiteStrictMode)
```

### 3. Refererチェックの実装
```go
func checkReferer(c *gin.Context) bool {
    referer := c.Request.Referer()
    if referer == "" {
        return false
    }
    allowedDomain := "localhost:5173"
    refererURL, err := url.Parse(referer)
    if err != nil {
        return false
    }
    return refererURL.Host == allowedDomain
}
```

### 4. カスタムヘッダーの要求
```go
// バックエンド（Go）
func customHeaderMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Request.Method != "GET" {
            if c.GetHeader("X-Requested-With") != "XMLHttpRequest" {
                c.AbortWithStatus(403)
                return
            }
        }
        c.Next()
    }
}
```

```javascript
// フロントエンド（Vue）
// Axiosの設定
axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
```

### 5. 環境変数を使用したAPIエンドポイント設定
```go
// .env
API_HOST=localhost
API_PORT=9090
FRONTEND_URL=http://localhost:5173
CORS_ALLOWED_ORIGINS=http://localhost:5173

// config/api.go
type APIConfig struct {
    Host            string
    Port            string
    FrontendURL     string
    AllowedOrigins  []string
}

func getAPIConfig() APIConfig {
    return APIConfig{
        Host:           os.Getenv("API_HOST"),
        Port:           os.Getenv("API_PORT"),
        FrontendURL:    os.Getenv("FRONTEND_URL"),
        AllowedOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
    }
}

// main.go
func main() {
    config := getAPIConfig()
    r := gin.Default()
    
    // CORSの設定
    r.Use(cors.New(cors.Config{
        AllowOrigins: config.AllowedOrigins,
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
    }))
    
    r.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
```

```javascript
// フロントエンド（Vue）
// .env
VITE_API_URL=http://localhost:9090

// api/config.js
export const API_URL = import.meta.env.VITE_API_URL;

// api/client.js
import axios from 'axios';
import { API_URL } from './config';

const apiClient = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});
```

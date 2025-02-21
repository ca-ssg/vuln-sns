# プロフィール更新機能のCSRF

## 脆弱性の説明
プロフィール更新処理でCSRF対策が実装されていないため、クロスサイトリクエストフォージェリ（CSRF）攻撃が可能です。

## 影響範囲
- 不正なプロフィール更新
- ユーザー情報の改ざん
- 意図しない操作の実行

## 確認手順
1. 別のWebサイトから攻撃用のフォームを作成
2. プロフィール更新のリクエストを送信
3. プロフィールが更新されることを確認

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

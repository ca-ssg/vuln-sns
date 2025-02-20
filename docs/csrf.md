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
1. CSRFトークンの実装
2. SameSite属性の設定
3. Refererチェックの実装
4. カスタムヘッダーの要求

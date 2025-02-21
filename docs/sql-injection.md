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

### 1. プリペアドステートメントの使用
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND password = SHA2('%s', 256)", 
    credentials.ID, credentials.Password)

// 修正後（安全なコード）
query := "SELECT id, nickname FROM users WHERE id = ? AND password = SHA2(?, 256)"
err := db.QueryRow(query, credentials.ID, credentials.Password).Scan(&user.ID, &user.Nickname)
```

### 2. パラメータのバインド
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')",
    post.UserID, post.Content)

// 修正後（安全なコード）
stmt, err := db.Prepare("INSERT INTO posts (user_id, content) VALUES (?, ?)")
if err != nil {
    return err
}
defer stmt.Close()
_, err = stmt.Exec(post.UserID, post.Content)
```

### 3. 入力値のバリデーション
```go
func validateInput(input string) error {
    // 特殊文字や制御文字のチェック
    if strings.ContainsAny(input, "'\"\\;") {
        return errors.New("invalid characters in input")
    }
    return nil
}
```

### 4. 最小権限原則の適用
```sql
-- アプリケーション用のDBユーザーを作成し、必要最小限の権限を付与
CREATE USER 'app_user'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT, INSERT, UPDATE ON vuln_app.* TO 'app_user'@'localhost';
-- DELETEやDROPなどの危険な操作は許可しない
```

### 5. 環境変数を使用したデータベース接続設定
```go
// .env
DB_USER=app_user
DB_PASSWORD=secure_password
DB_NAME=vuln_app
DB_HOST=localhost
DB_PORT=3306

// config/database.go
func getDSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"))
}

// database/mysql.go
func InitDB() (*sql.DB, error) {
    dsn := getDSN()
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    // ...
    return db, nil
}
```

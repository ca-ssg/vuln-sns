# ログイン機能のSQLインジェクション

## 脆弱性の説明
ログイン処理でユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- ユーザー認証のバイパス
- データベース情報の漏洩
- 不正なアカウントアクセス

## 確認手順
1. ログインフォームにSQLインジェクションペイロードを入力
2. 認証をバイパスしてログインできることを確認

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

### 2. 環境変数を使用したデータベース接続設定
```go
// .env
DB_USER=app_user
DB_PASSWORD=password
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
```

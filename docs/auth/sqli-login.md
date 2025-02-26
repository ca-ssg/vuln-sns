# ログイン機能のSQLインジェクション

## 脆弱性の説明
ログイン処理でユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- ユーザー認証のバイパス
- データベース情報の漏洩
- 不正なアカウントアクセス

## 攻撃方法と手順

### 1. 認証バイパス攻撃

1. アプリケーションのログイン画面にアクセスします
2. ユーザーIDフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   alice' OR '1'='1
   ```
   または
   ```
   admin' --
   ```
3. パスワードフィールドには任意の値（例：`password`）を入力します
4. ログインボタンをクリックします

このペイロードは、WHERE句の条件を常に真にするため、データベース内の最初のユーザー（多くの場合は管理者）としてログインできます。

### 2. ユーザー情報の抽出攻撃

1. アプリケーションのログイン画面にアクセスします
2. ユーザーIDフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   ' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1
   ```
3. パスワードフィールドには任意の値を入力します
4. ログインボタンをクリックします

このペイロードは、UNIONを使用して偽のユーザー情報を結果セットに追加し、認証をバイパスします。

### 3. データベース情報の取得攻撃

1. アプリケーションのログイン画面にアクセスします
2. ユーザーIDフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   ' UNION SELECT table_name, column_name FROM information_schema.columns WHERE table_schema = 'vuln_app' --
   ```
3. パスワードフィールドには任意の値を入力します
4. ログインボタンをクリックします

このペイロードは、データベースのスキーマ情報を取得しようとします。エラーメッセージが表示される場合は、そのエラーメッセージからデータベース構造に関する情報が漏洩している可能性があります。

## 攻撃成功の確認手順

### 1. 認証バイパスの確認

1. 上記のペイロードを使用してログインを試みます
2. ログイン後、ダッシュボードやプロフィールページなど、認証が必要なページにアクセスできることを確認します
3. ページ上にユーザー名やプロフィール情報が表示されていることを確認します

### 2. エラーメッセージからの情報漏洩の確認

1. 意図的に不正なSQLインジェクションペイロードを入力します（例：`' OR 1=1;`）
2. エラーメッセージが表示される場合、そのメッセージにデータベース情報（テーブル名、カラム名、SQLクエリの構造など）が含まれていないか確認します

### 3. ブラウザの開発者ツールでのレスポンス確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Network」タブを選択します
3. SQLインジェクションペイロードを使用してログインを試みます
4. ログインリクエストのレスポンスを確認し、認証トークンや成功メッセージが含まれているか確認します

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

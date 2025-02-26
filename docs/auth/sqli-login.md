# ログイン機能のセキュリティ

## 実装の説明
現在のログイン処理は、SQLクエリを使用せずに実装されています。代わりに、任意のユーザーIDを受け入れ、それに基づいてトークンを生成します。

## 影響範囲
- 認証のバイパス（任意のユーザーIDでログイン可能）
- セッション管理の問題

## 攻撃方法と手順

### 1. 認証バイパス攻撃

1. アプリケーションのログイン画面にアクセスします
2. ユーザーIDフィールドに任意のユーザーID（例：`admin`）を入力します
3. パスワードフィールドには任意の値（例：`password`）を入力します
4. ログインボタンをクリックします

現在の実装では、入力されたユーザーIDに基づいてトークンが生成されるため、任意のユーザーとしてログインできます。

## 攻撃成功の確認手順

### 1. 認証バイパスの確認

1. 任意のユーザーID（例：`admin`、`alice`、`bob`など）でログインを試みます
2. ログイン後、ダッシュボードやプロフィールページなど、認証が必要なページにアクセスできることを確認します
3. ページ上にユーザー名やプロフィール情報が表示されていることを確認します

### 2. ブラウザの開発者ツールでのレスポンス確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Network」タブを選択します
3. 任意のユーザーIDでログインを試みます
4. ログインリクエストのレスポンスを確認し、認証トークンや成功メッセージが含まれているか確認します

## 対策方法
### 1. 適切なユーザー認証の実装
```go
// 修正前（現在のコード）
// For testing, accept any user_id and return a simple token
log.Printf("Login successful for user: %s", credentials.UserID)
c.JSON(http.StatusOK, gin.H{
    "token": credentials.UserID + "_token",
    "user": models.User{
        ID:       credentials.UserID,
        Nickname: credentials.UserID,
    },
})

// 修正後（安全なコード）
query := "SELECT id, nickname FROM users WHERE id = ? AND password = SHA2(?, 256)"
var user models.User
err := h.db.QueryRow(query, credentials.UserID, credentials.Password).Scan(&user.ID, &user.Nickname)
if err != nil {
    if err == sql.ErrNoRows {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    log.Printf("Database error: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
    return
}

// トークン生成（JWTなどを使用）
token := generateSecureToken(user.ID)
c.JSON(http.StatusOK, gin.H{
    "token": token,
    "user": user,
})
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

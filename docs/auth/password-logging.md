# ログイン処理におけるパスワードのログ出力

## 脆弱性の説明
ログイン処理において、ユーザーが入力したパスワードがそのままログに出力されています。これにより、ログファイルにアクセスできる人物（システム管理者など）が、ユーザーのパスワードを閲覧できる状態になっています。

## 影響範囲
- ユーザーパスワードの漏洩
- プライバシーの侵害
- 不正アクセスのリスク増加

## 攻撃方法と手順

### 1. ログファイルへのアクセス

1. サーバーのログファイルにアクセスします
2. ログイン処理のログを検索します
3. 以下のようなログエントリを見つけます：
   ```
   Executing query: SELECT id, nickname FROM users WHERE id = 'alice' AND (password = 'alice' OR password = SHA2('alice', 256))
   ```
4. ログから直接ユーザーIDとパスワードを取得できます

## 攻撃成功の確認手順

### 1. ログファイルの確認

1. アプリケーションのログファイルを確認します
2. ログイン処理のログエントリを探します
3. ユーザーIDとパスワードが平文で記録されていることを確認します

### 2. 複数ユーザーのパスワード収集

1. 複数のユーザーがログインするのを待ちます
2. ログファイルを定期的に確認します
3. 複数のユーザーのパスワードを収集できることを確認します

## 対策方法

### 1. センシティブ情報のログ出力を避ける
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND (password = '%s' OR password = SHA2('%s', 256))", 
    credentials.UserID, credentials.Password, credentials.Password)
log.Printf("Executing query: %s", query)

// 修正後（安全なコード）
query := fmt.Sprintf("SELECT id, nickname FROM users WHERE id = '%s' AND (password = '%s' OR password = SHA2('%s', 256))", 
    credentials.UserID, credentials.Password, credentials.Password)
log.Printf("Executing login query for user: %s", credentials.UserID)
```

### 2. ログ出力時のマスキング処理
```go
// 修正前（脆弱なコード）
log.Printf("Executing query: %s", query)

// 修正後（安全なコード）
// クエリからパスワード部分をマスクする関数
func maskPasswordInQuery(query string) string {
    // 正規表現を使用してパスワード部分をマスク
    re := regexp.MustCompile(`password = '([^']*)'`)
    return re.ReplaceAllString(query, "password = '********'")
}
log.Printf("Executing query: %s", maskPasswordInQuery(query))
```

### 3. 構造化ログの使用
```go
// 修正前（脆弱なコード）
log.Printf("Executing query: %s", query)

// 修正後（安全なコード）
logger.Info("Executing login query",
    zap.String("user_id", credentials.UserID),
    zap.String("query_type", "login"),
    zap.String("client_ip", c.ClientIP()),
)
```

## セキュリティのベストプラクティス
1. パスワードなどの機密情報は絶対にログに出力しない
2. デバッグ目的でもパスワードをログに出力しない
3. ログレベルを適切に設定し、本番環境では詳細なクエリログを無効にする
4. 構造化ログを使用し、センシティブ情報をフィールドから除外する
5. ログローテーションとログの保持期間を適切に設定する

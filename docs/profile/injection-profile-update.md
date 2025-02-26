# SQLインジェクション（プロフィール更新）

このアプリケーションは学習目的で意図的にセキュリティ脆弱性を含んでいます：

## 概要
プロフィール更新機能にSQLインジェクションの脆弱性が存在します：
```go
query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
```

## 学習目的
1. SQLインジェクションのメカニズムの理解
2. パラメータ化クエリの重要性の理解
3. 適切な入力サニタイズ方法の学習

## 実装箇所
- `backend/internal/handlers/auth.go` UpdateProfileメソッド

## 攻撃方法と手順

### 1. プロフィール更新時のSQLインジェクション

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   新しいニックネーム', bio='ハッキングされました' WHERE id='admin' --
   ```
4. 更新ボタンをクリックします

このペイロードは、SQLクエリの構造を操作して、他のユーザー（この場合は管理者）のプロフィール情報も更新しようとします。

### 2. 複数ユーザーのプロフィールを一度に更新するSQLインジェクション

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   ハッキングされました' WHERE id LIKE '%' --
   ```
4. 更新ボタンをクリックします

このペイロードは、WHERE句の条件を操作して、すべてのユーザーのニックネームを一度に更新しようとします。

### 3. データベーススキーマ情報を取得するSQLインジェクション

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   test', (SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE() LIMIT 1)) -- 
   ```
4. 更新ボタンをクリックします

このペイロードは、データベースのスキーマ情報を取得しようとします。エラーメッセージが表示される場合は、そのエラーメッセージからデータベース構造に関する情報が漏洩している可能性があります。

## 攻撃成功の確認手順

### 1. 他ユーザーのプロフィール更新の確認

1. 上記のペイロードを使用してプロフィールを更新します
2. 別のユーザーアカウント（例：admin）でログインします
3. プロフィールページにアクセスし、情報が変更されているか確認します

### 2. エラーメッセージからの情報漏洩の確認

1. 意図的に不正なSQLインジェクションペイロードを入力します（例：`', (SELECT * FROM non_existent_table) --`）
2. エラーメッセージが表示される場合、そのメッセージにデータベース情報（テーブル名、カラム名、SQLクエリの構造など）が含まれていないか確認します

### 3. ブラウザの開発者ツールでのレスポンス確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Network」タブを選択します
3. SQLインジェクションペイロードを使用してプロフィールを更新します
4. リクエストのレスポンスを確認し、エラーメッセージやデータベース情報が含まれているか確認します

## 対策方法
本番環境では以下の対策が必要：

### 1. パラメータ化クエリの使用
```go
// 修正前
query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
_, err := h.db.Exec(query)

// 修正後
query := "UPDATE users SET nickname = ? WHERE id = ?"
_, err := h.db.Exec(query, profile.Nickname, userID)
```

### 2. 入力値のバリデーション
```go
// 修正前
if err := c.BindJSON(&profile); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    return
}

// 修正後
if err := c.BindJSON(&profile); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
    return
}

// ニックネームの長さと文字種のバリデーション
if len(profile.Nickname) > 50 || !validateNickname(profile.Nickname) {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nickname format"})
    return
}

func validateNickname(nickname string) bool {
    // 許可する文字をregexで定義
    pattern := "^[a-zA-Z0-9_-]{1,50}$"
    match, _ := regexp.MatchString(pattern, nickname)
    return match
}
```

### 3. 入力値のサニタイズ
```go
// 修正前
// サニタイズなし

// 修正後
import (
    "html"
    "strings"
)

func sanitizeInput(input string) string {
    // HTMLエスケープ
    escaped := html.EscapeString(input)
    // 特殊文字の除去
    escaped = strings.ReplaceAll(escaped, "'", "")
    escaped = strings.ReplaceAll(escaped, "\"", "")
    return escaped
}

// 使用例
profile.Nickname = sanitizeInput(profile.Nickname)
```

### 4. エラーハンドリングの改善
```go
// 修正前
if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
    return
}

// 修正後
if err != nil {
    if strings.Contains(err.Error(), "Duplicate entry") {
        c.JSON(http.StatusConflict, gin.H{"error": "ニックネームが既に使用されています"})
        return
    }
    log.Printf("Error updating profile: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "内部エラーが発生しました"})
    return
}
```

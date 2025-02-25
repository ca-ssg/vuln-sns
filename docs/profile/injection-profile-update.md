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

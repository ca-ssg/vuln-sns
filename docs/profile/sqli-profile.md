# プロフィール機能のSQLインジェクション

## 脆弱性の説明
プロフィール更新処理で、ユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- 不正なプロフィール更新
- データベース情報の漏洩
- 他ユーザーの情報改ざん

## 確認手順
1. プロフィール更新フォームにSQLインジェクションペイロードを入力
2. 不正な更新が行われることを確認

## 対策方法
### 1. プリペアドステートメントの使用
```go
// 修正前（脆弱なコード）
query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"

// 修正後（安全なコード）
query := "UPDATE users SET nickname = ? WHERE id = ?"
_, err = db.Exec(query, profile.Nickname, userID)
```

### 2. 入力値のバリデーション
```go
func validateInput(input string) error {
    // 特殊文字や制御文字のチェック
    if strings.ContainsAny(input, "'\"\\;") {
        return errors.New("invalid characters in input")
    }
    return nil
}
```

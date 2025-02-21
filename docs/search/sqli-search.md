# 検索機能のSQLインジェクション

## 脆弱性の説明
検索処理で、ユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- データベース情報の漏洩
- 不正な検索結果の取得
- システム情報の漏洩

## 確認手順
1. 検索フォームにSQLインジェクションペイロードを入力
2. 不正な検索結果が表示されることを確認

## 対策方法
### 1. プリペアドステートメントの使用
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("SELECT * FROM posts WHERE content LIKE '%%%s%%'", searchTerm)

// 修正後（安全なコード）
query := "SELECT * FROM posts WHERE content LIKE ?"
searchPattern := "%" + searchTerm + "%"
rows, err := db.Query(query, searchPattern)
```

### 2. 入力値のバリデーション
```go
func validateSearchTerm(term string) error {
    // 特殊文字や制御文字のチェック
    if strings.ContainsAny(term, "'\"\\;") {
        return errors.New("invalid characters in search term")
    }
    return nil
}
```

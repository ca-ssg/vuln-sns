# 投稿機能のSQLインジェクション

## 脆弱性の説明
投稿の作成・更新・削除処理で、ユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- 不正な投稿の作成・更新・削除
- データベース情報の漏洩
- 他ユーザーの投稿の改ざん

## 確認手順
1. 投稿作成フォームにSQLインジェクションペイロードを入力
2. 不正な投稿が作成されることを確認

## 対策方法
### 1. プリペアドステートメントの使用
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')", userID, post.Content)

// 修正後（安全なコード）
stmt, err := db.Prepare("INSERT INTO posts (user_id, content) VALUES (?, ?)")
if err != nil {
    return err
}
defer stmt.Close()
_, err = stmt.Exec(userID, post.Content)
```

### 2. パラメータのバインド
```go
// 修正前（脆弱なコード）
query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %s AND user_id = '%s'", 
    post.Content, postID, userID)

// 修正後（安全なコード）
query := "UPDATE posts SET content = ? WHERE id = ? AND user_id = ?"
_, err = db.Exec(query, post.Content, postID, userID)
```

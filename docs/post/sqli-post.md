# 投稿機能のSQLインジェクション

## 脆弱性の説明
投稿の作成・更新・削除処理で、ユーザー入力を直接SQLクエリに結合しているため、SQLインジェクションが可能です。

## 影響範囲
- 不正な投稿の作成・更新・削除
- データベース情報の漏洩
- 他ユーザーの投稿の改ざん

## 攻撃方法と手順

### 1. 投稿作成時のSQLインジェクション

1. アプリケーションにログインします
2. 投稿作成フォームにアクセスします
3. 投稿内容フィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   通常の投稿'), ('hacker', 'ハッキングされました
   ```
4. 投稿ボタンをクリックします
5. 画面を更新してみてください

このペイロードは、SQLクエリの構造を操作して、追加の投稿を作成したり、機密情報を抽出しようとします。

### 2. 投稿更新時のSQLインジェクション

1. アプリケーションにログインします
2. 自分の投稿を編集するページにアクセスします
3. 投稿内容フィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   新しい内容'), ('hacker', (SELECT CONCAT('ユーザー情報: ', id, ':', password) FROM users LIMIT 1)) -- 
   ```
4. 更新ボタンをクリックします
5. 画面を更新してみてください

このペイロードは、ユーザーテーブルからユーザーIDとパスワードハッシュを抽出しようとします。

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

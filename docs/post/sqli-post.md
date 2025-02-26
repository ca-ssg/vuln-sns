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
   通常の投稿内容', (SELECT password FROM users WHERE id = 'admin')) -- 
   ```
   または
   ```
   '); INSERT INTO posts (user_id, content) VALUES ('hacker', 'ハッキングされました'); --
   ```
4. 投稿ボタンをクリックします

このペイロードは、SQLクエリの構造を操作して、追加の投稿を作成したり、機密情報を抽出しようとします。

### 2. 投稿更新時のSQLインジェクション

1. アプリケーションにログインします
2. 自分の投稿を編集するページにアクセスします
3. 投稿内容フィールドに以下のSQLインジェクションペイロードを入力します：
   ```
   新しい内容'), (SELECT CONCAT(id, ':', password) FROM users LIMIT 1)) -- 
   ```
4. 更新ボタンをクリックします

このペイロードは、ユーザーテーブルからユーザーIDとパスワードハッシュを抽出しようとします。

### 3. 他ユーザーの投稿を改ざんするSQLインジェクション

1. アプリケーションにログインします
2. 自分の投稿を編集するページにアクセスします
3. 投稿IDパラメータを操作します（URLまたはフォームの隠しフィールド）：
   ```
   1 OR id IN (SELECT id FROM posts WHERE user_id = 'admin')
   ```
4. 投稿内容を変更して更新ボタンをクリックします

このペイロードは、WHERE句の条件を操作して、他のユーザー（この場合は管理者）の投稿を更新しようとします。

## 攻撃成功の確認手順

### 1. 不正な投稿作成の確認

1. 上記のペイロードを使用して投稿を作成します
2. 投稿一覧ページに移動します
3. 意図した通りに複数の投稿が作成されているか、または投稿内容に機密情報が含まれているかを確認します

### 2. エラーメッセージからの情報漏洩の確認

1. 意図的に不正なSQLインジェクションペイロードを入力します（例：`'); SELECT * FROM non_existent_table; --`）
2. エラーメッセージが表示される場合、そのメッセージにデータベース情報（テーブル名、カラム名、SQLクエリの構造など）が含まれていないか確認します

### 3. ブラウザの開発者ツールでのレスポンス確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Network」タブを選択します
3. SQLインジェクションペイロードを使用して投稿を作成または更新します
4. リクエストのレスポンスを確認し、エラーメッセージやデータベース情報が含まれているか確認します

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

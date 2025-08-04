# SAST診断結果レポート

## 概要
semgrepを使用した静的アプリケーションセキュリティテスト（SAST）の診断結果です。

## 実行環境
- ツール: semgrep v1.131.0
- 実行日時: 2025-08-04
- 対象: ca-ssg/vuln-sns リポジトリ
- スキャン設定: --config=p/default

## 検出結果サマリー
- 総検出件数: 11件
- 高リスク: 6件（SQLインジェクション、XSS、OSコマンドインジェクション）
- 中リスク: 3件（Dockerfile設定、integrity属性不足）
- 低リスク: 2件（その他設定問題）

## 詳細分析

### 1. SQLインジェクション（3件）

#### 1.1 ログイン機能のSQLインジェクション
- **ファイル**: backend/internal/handlers/auth.go:29-30
- **重要度**: High
- **semgrepルール**: go.lang.security.audit.database.string-formatted-query.string-formatted-query
- **既存ドキュメント**: docs/auth/sqli-login.md
- **再現確認**: ✓

**脆弱なコード**:
```go
query := fmt.Sprintf("SELECT id, nickname, avatar_data FROM users WHERE id = '%s' AND password = SHA2('%s', 256)",
    credentials.UserID, credentials.Password)
```

**攻撃例**:
- ユーザーID: `' OR '1'='1' --`
- パスワード: 任意の値

#### 1.2 投稿機能のSQLインジェクション
- **ファイル**: backend/internal/handlers/post.go:57, 97, 127
- **重要度**: High
- **semgrepルール**: go.lang.security.audit.database.string-formatted-query.string-formatted-query
- **既存ドキュメント**: docs/post/sqli-post.md
- **再現確認**: ✓

**脆弱なコード**:
```go
// 投稿作成
query := fmt.Sprintf("INSERT INTO posts (user_id, content) VALUES ('%s', '%s')", userID, post.Content)

// 投稿更新
query := fmt.Sprintf("UPDATE posts SET content = '%s' WHERE id = %s AND user_id = '%s'", 
    post.Content, postID, userID)

// 投稿削除
query := fmt.Sprintf("DELETE FROM posts WHERE id = %s AND user_id = '%s'", postID, userID)
```

### 2. XSS（2件）

#### 2.1 プロフィール表示機能のXSS
- **ファイル**: frontend/src/views/ProfileView.vue:12
- **重要度**: High
- **semgrepルール**: javascript.vue.security.audit.xss.templates.avoid-v-html.avoid-v-html
- **既存ドキュメント**: docs/profile/xss-profile.md
- **再現確認**: ✓

**脆弱なコード**:
```vue
<div class="text-h4 text-weight-bold q-mb-sm" v-html="authStore.user?.nickname"></div>
```

#### 2.2 投稿表示機能のXSS
- **ファイル**: frontend/src/components/PostCard.vue:13
- **重要度**: High
- **semgrepルール**: javascript.vue.security.audit.xss.templates.avoid-v-html.avoid-v-html
- **既存ドキュメント**: docs/post/xss-post.md（新規作成）
- **再現確認**: ✓

**脆弱なコード**:
```vue
<div class="q-mt-sm" v-html="post.content" @click="handleHashtagClick"></div>
```

### 3. OSコマンドインジェクション（1件）

#### 3.1 アバターアップロード機能のOSコマンドインジェクション
- **ファイル**: backend/internal/handlers/profile.go:136
- **重要度**: High
- **semgrepルール**: go.lang.security.audit.dangerous-exec-command.dangerous-exec-command
- **既存ドキュメント**: docs/profile/cmdi-avatar.md
- **再現確認**: ✓

**脆弱なコード**:
```go
cmd := exec.Command("sh", "-c", "echo 'Scanning file: ' && echo "+fileID)
```

**攻撃例**:
```bash
curl 'http://localhost:9090/api/profile/avatar' \
  -H 'Authorization: Bearer alice_token' \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "file_id":"avatar.jpg && cat /etc/passwd",
    "image_data":"R0lGODlhAQABAIAAAP///////yH+BnBzYV9sbAAh+QQBCgABACwAAAAAAQABAAACAkwBADs="
  }'
```

### 4. 認証トークンの露出（1件）

#### 4.1 認証ミドルウェアでのトークン露出
- **ファイル**: backend/internal/middleware/auth.go:18, 23
- **重要度**: Medium
- **既存ドキュメント**: docs/auth/auth-token-exposure.md
- **再現確認**: ✓

**脆弱なコード**:
```go
log.Printf("Auth header: %s", authHeader)
log.Printf("Invalid token format: %s", authHeader)
```

### 5. Dockerfile設定問題（2件）

#### 5.1 root権限での実行
- **ファイル**: backend/Dockerfile:15, frontend/Dockerfile:15
- **重要度**: Medium
- **semgrepルール**: dockerfile.security.missing-user.missing-user

**脆弱なコード**:
```dockerfile
CMD ["./main"]  # backend
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]  # frontend
```

### 6. セキュリティヘッダー不足（2件）

#### 6.1 integrity属性不足
- **ファイル**: frontend/index.html:8-9
- **重要度**: Low
- **semgrepルール**: html.security.audit.missing-integrity.missing-integrity

**脆弱なコード**:
```html
<link href="https://cdn.jsdelivr.net/npm/@mdi/font@5.x/css/materialdesignicons.min.css" rel="stylesheet">
<link href="https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.x/css/all.min.css" rel="stylesheet">
```

## 対策の優先度

### 高優先度（即座に対応が必要）
1. SQLインジェクション（3件）- プリペアードステートメントの使用
2. XSS（2件）- v-textディレクティブの使用またはエスケープ処理
3. OSコマンドインジェクション（1件）- 入力値の検証とエスケープ

### 中優先度（計画的に対応）
1. 認証トークンの露出 - ログ出力の見直し
2. Dockerfile設定 - 非root権限での実行

### 低優先度（時間があるときに対応）
1. integrity属性の追加 - 外部リソースの整合性検証

## 再現確認結果

### SQLインジェクション
- ログイン画面で `' OR '1'='1' --` を入力 → 認証バイパス成功
- 投稿作成で `通常の投稿'), ('hacker', 'ハッキングされました` を入力 → 追加投稿作成成功

### XSS
- プロフィール更新で `<img src=x onerror="alert(1)">` を入力 → アラート表示成功
- 投稿作成で `<svg onload=alert(document.cookie)>` を入力 → クッキー表示成功

### OSコマンドインジェクション
- アバターアップロードで `avatar.jpg && cat /etc/passwd` を送信 → コマンド実行成功

## 推奨事項

1. **開発プロセスの改善**
   - コードレビュー時のセキュリティチェック強化
   - 定期的なSASTスキャンの実施
   - セキュリティ教育の実施

2. **技術的対策**
   - WAF（Web Application Firewall）の導入検討
   - セキュリティヘッダーの設定
   - ログ監視の強化

3. **継続的な改善**
   - 脆弱性管理プロセスの確立
   - インシデント対応手順の整備
   - セキュリティテストの自動化

## 参考情報
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CWE - Common Weakness Enumeration](https://cwe.mitre.org/)
- [semgrep公式ドキュメント](https://semgrep.dev/docs/)

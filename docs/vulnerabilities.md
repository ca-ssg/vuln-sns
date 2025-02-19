# 脆弱性の詳細

このドキュメントでは、アプリケーションに意図的に実装された脆弱性の詳細と確認手順について説明します。

## 実装されている脆弱性

1. SQLインジェクション
   - ログイン処理 (`backend/internal/handlers/auth.go`)
     - ユーザーIDとパスワードの検証時に文字列連結を使用
   - 投稿の操作 (`backend/internal/handlers/post.go`)
     - 投稿の作成・編集・削除時にクエリを文字列連結で構築

2. XSS（クロスサイトスクリプティング）
   - 投稿内容の表示 (`frontend/src/components/PostCard.vue`)
     - `v-html`ディレクティブによる生のHTML表示
   - プロフィール表示 (`frontend/src/views/ProfileView.vue`)
     - ニックネームのエスケープ処理なし

3. CSRF（クロスサイトリクエストフォージェリ）
   - いいね機能 (`backend/internal/handlers/post.go`)
     - CSRFトークンの検証なし
   - プロフィール更新 (`backend/internal/handlers/auth.go`)
     - CSRFトークンの検証なし

4. セッション管理の不備
   - JWT実装 (`backend/internal/middleware/auth.go`)
     - トークンの有効期限が長期間
     - トークンの再利用が可能
   - セッション管理
     - セッションの無効化機能なし

5. アクセス制御の不備
   - 投稿の編集・削除 (`backend/internal/handlers/post.go`)
     - ユーザー所有権の検証が不十分
   - プロフィール更新 (`backend/internal/handlers/auth.go`)
     - ユーザー認証の検証が不十分

## 確認手順

1. SQLインジェクション
   ```sql
   -- ログイン時のSQLインジェクション
   ID: alice'; DROP TABLE users; --
   Password: anything
   ```

2. XSS
   ```html
   <!-- 投稿内容でのXSS -->
   <script>alert('XSS!');</script>
   ```

3. CSRF
   - いいねボタンやプロフィール更新時のCSRFトークンなし
   - 別オリジンからのリクエストが可能

4. アクセス制御
   - 他ユーザーの投稿を編集・削除可能
   - URLパラメータの改ざんで他ユーザーの情報にアクセス可能

**注意**: このアプリケーションは学習目的で作成されています。本番環境での使用や、実際のサービスへの攻撃に使用することは絶対に避けてください。

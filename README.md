# Vulnerable SNS Application

**⚠️ 警告: このアプリケーションには意図的に脆弱性が含まれています。学習目的以外での使用は避けてください。**

TwitterライクなSNSアプリケーションで、セキュリティ学習用に意図的に脆弱性を組み込んでいます。

## 技術スタック

- フロントエンド: Vue3
- バックエンドAPI: Go
- データベース: MySQL
- オーケストレーション: Docker Compose

## 機能

- 投稿の閲覧（ログイン不要）
- ユーザーログイン
- 投稿の作成・編集・削除（ログインユーザーのみ）
- プロフィール編集（ニックネーム変更）
- いいね機能

## 初期アカウント

```
ID: alice
パスワード: alice
```

## セットアップ

```bash
# リポジトリのクローン
git clone https://github.com/ca-ssg/devin-vuln-app.git
cd devin-vuln-app

# アプリケーションの起動
docker-compose up -d

# フロントエンド: http://localhost:5173
# バックエンドAPI: http://localhost:8080
```

## プロジェクト構成

```
.
├── frontend/          # Vue3フロントエンド
├── backend/           # Goバックエンド
├── docker-compose.yml # コンテナオーケストレーション
└── README.md
```

## ⚠️ 脆弱性について

このアプリケーションには以下のような脆弱性が意図的に実装されています：

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

**注意**: このアプリケーションは学習目的で作成されています。本番環境での使用や、実際のサービスへの攻撃に使用することは絶対に避けてください。

## 脆弱性の学習方法

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

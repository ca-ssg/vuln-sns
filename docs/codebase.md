# コードベース詳細ドキュメント

## プロジェクト概要

本プロジェクトは、X（旧Twitter）ライクなSNSアプリケーションで、**セキュリティ学習用に意図的に脆弱性を組み込んだ教育目的のアプリケーション**です。実際のサービスとして使用することは推奨されません。

### 主な機能
- 投稿の閲覧（認証不要）
- ユーザーログイン
- 投稿の作成・編集・削除（認証必須）
- プロフィール編集（ニックネーム、アバター）
- いいね機能
- ハッシュタグ検索

## 技術スタック

- **Frontend**: Vue 3 + TypeScript + Quasar UI + Tailwind CSS
- **Backend**: Go with Gin framework
- **Database**: MySQL 8.0
- **Infrastructure**: Docker Compose

## アーキテクチャ

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   ブラウザ   │────▶│ フロントエンド │────▶│ バックエンド  │
│             │     │   (Vue 3)    │     │    (Go)     │
└─────────────┘     └─────────────┘     └──────┬──────┘
                         :5173                   │ :9090
                                                 ▼
                                          ┌─────────────┐
                                          │   MySQL     │
                                          │   (8.0)     │
                                          └─────────────┘
                                               :3306
```

## コマンド

### 開発環境の起動・停止
```bash
# Docker環境の起動
docker-compose up -d
# または
make start

# Docker環境の停止
docker-compose down
# または
make stop

# 再ビルドして再起動
make restart

# 全データをリセット（ボリュームも削除）
make reset

# 全コンテナとボリュームを削除
make clear
```

### フロントエンド開発（frontend/ディレクトリで実行）
```bash
# 依存関係のインストール
npm install

# 開発サーバーの起動
npm run dev

# プロダクションビルド
npm run build

# 型チェック
npm run type-check

# Linting
npm run lint

# コードフォーマット
npm run format
```

### データベース操作
```bash
# MySQLへの接続
docker-compose exec db mysql -uroot -ppassword -Dvuln_app --default-character-set=utf8mb4
# または
make db
```

### サービスURL
- フロントエンド: http://localhost:5173
- バックエンドAPI: http://localhost:9090
- APIエンドポイントプレフィックス: /api

### テストアカウント
- alice / alice
- bob / bob
- charlie / charlie

## ディレクトリ構成

### プロジェクトルート
```
.
├── frontend/              # Vue3フロントエンド
├── backend/               # Goバックエンド
├── docs/                  # ドキュメント
│   ├── vulnerabilities.md # 脆弱性の概要
│   ├── auth/             # 認証関連の脆弱性
│   ├── post/             # 投稿関連の脆弱性
│   └── profile/          # プロフィール関連の脆弱性
├── nuclei/               # Nucleiテンプレート（脆弱性検証用）
├── docker-compose.yml    # Docker Compose設定
├── docker-compose.prd.yml # 本番環境用Docker設定
├── init.sql              # データベース初期化SQL
├── my.cnf                # MySQL設定
├── Makefile              # Make コマンド定義
├── CLAUDE.md             # Claude AI用のプロジェクト説明
└── README.md             # プロジェクト説明
```

### フロントエンド構造
```
frontend/
├── src/
│   ├── views/           # メインアプリケーションページ
│   │   ├── HomeView.vue    # ホーム画面（投稿一覧）
│   │   ├── LoginView.vue   # ログイン画面
│   │   └── ProfileView.vue # プロフィール画面
│   ├── components/      # 再利用可能なVueコンポーネント
│   │   ├── PostCard.vue    # 投稿カード
│   │   ├── PostDialog.vue  # 投稿作成・編集ダイアログ
│   │   └── BaseDialog.vue  # 基本ダイアログ
│   ├── stores/          # Piniaステート管理
│   │   ├── auth.ts         # 認証ストア
│   │   └── posts.ts        # 投稿ストア
│   ├── router/          # Vue Router設定
│   │   └── index.ts        # ルーティング定義
│   ├── types/           # TypeScript型定義
│   ├── assets/          # 静的アセット
│   ├── App.vue          # ルートコンポーネント
│   └── main.ts          # エントリーポイント
├── public/              # 公開ディレクトリ
├── package.json         # npm依存関係
├── vite.config.ts       # Vite設定
├── tsconfig.json        # TypeScript設定
└── tailwind.config.js   # Tailwind CSS設定
```

### バックエンド構造
```
backend/
├── internal/
│   ├── handlers/        # HTTPリクエストハンドラー
│   │   ├── handler.go      # 共通ハンドラー
│   │   ├── auth.go         # 認証ハンドラー
│   │   ├── post.go         # 投稿ハンドラー
│   │   ├── profile.go      # プロフィールハンドラー
│   │   └── search.go       # 検索ハンドラー
│   ├── middleware/      # ミドルウェア
│   │   └── auth.go         # 認証ミドルウェア
│   ├── models/          # データモデル
│   │   ├── user.go         # ユーザーモデル
│   │   └── post.go         # 投稿モデル
│   └── database/        # データベース接続
│       └── db.go           # DB接続管理
├── main.go              # エントリーポイント
├── go.mod               # Go モジュール定義
└── go.sum               # Go 依存関係ロック
```

## APIエンドポイント

### 認証不要エンドポイント
- `GET /api/posts` - 全投稿の取得
- `POST /api/login` - ユーザーログイン
- `GET /api/search` - ハッシュタグ検索
- `GET /api/health` - ヘルスチェック

### 認証必須エンドポイント
- `POST /api/posts` - 投稿の作成
- `PUT /api/posts/:id` - 投稿の更新
- `DELETE /api/posts/:id` - 投稿の削除
- `POST /api/posts/:id/like` - 投稿にいいね
- `DELETE /api/posts/:id/like` - いいねを取り消し
- `GET /api/profile` - プロフィール取得
- `PUT /api/profile` - プロフィール更新
- `POST /api/profile/avatar` - アバターアップロード

## データ設計

### データベーススキーマ

#### usersテーブル
```sql
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,        -- ユーザーID
    password VARCHAR(255) NOT NULL,     -- パスワード（SHA256ハッシュ）
    nickname VARCHAR(255) NOT NULL,     -- ニックネーム
    avatar_data LONGTEXT DEFAULT NULL   -- アバター画像（Base64）
);
```

#### postsテーブル
```sql
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,  -- 投稿ID
    user_id VARCHAR(255) NOT NULL,      -- 投稿者ID
    content TEXT NOT NULL,              -- 投稿内容
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

#### likesテーブル
```sql
CREATE TABLE likes (
    post_id INT NOT NULL,               -- 投稿ID
    user_id VARCHAR(255) NOT NULL,      -- いいねしたユーザーID
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## アプリケーション設計

### フロントエンド設計

#### 状態管理（Pinia）
- **authStore**: ユーザー認証状態、ログイン/ログアウト処理
- **postsStore**: 投稿データの管理、CRUD操作

#### コンポーネント設計
- **PostCard**: 投稿の表示、いいね機能、編集・削除ボタン
- **PostDialog**: 投稿の作成・編集フォーム
- **BaseDialog**: 汎用的なダイアログコンポーネント

#### ルーティング
- `/`: ホーム画面（投稿一覧）
- `/login`: ログイン画面
- `/profile`: プロフィール画面（認証必須）

### バックエンド設計

#### ミドルウェア
- **CORS設定**: 特定のオリジンからのアクセスを許可
- **認証ミドルウェア**: JWTトークンによる認証

#### ハンドラー設計
- **エラーハンドリング**: 統一的なエラーレスポンス
- **ログ出力**: デバッグモードで全リクエストをログ出力

#### データベース接続
- 接続リトライ機能（最大30回）
- マルチステートメント対応
- UTF-8MB4文字セット対応

## セキュリティのコンテキスト

### 意図的な脆弱性

本アプリケーションには教育目的で以下の脆弱性が意図的に実装されています：

1. **認証機能の脆弱性**
   - ログイン機能のSQLインジェクション
   - 認証トークンの露出
   - パスワードのログ出力

2. **投稿機能の脆弱性**
   - 投稿機能のSQLインジェクション
   - 投稿表示機能のXSS

3. **プロフィール機能の脆弱性**
   - プロフィール表示機能のXSS
   - アバターアップロード機能のOSコマンドインジェクション

### セキュリティ検証

`nuclei/`ディレクトリに脆弱性検証用のNucleiテンプレートが含まれています。

## 重要なルール

1. **本番環境での使用禁止**
   - このアプリケーションは学習目的のみで使用すること
   - 実際のサービスとして公開しないこと

2. **脆弱性の取り扱い**
   - 脆弱性は意図的に実装されているため、修正しないこと
   - セキュリティ対策を追加する場合は、明示的に要求された場合のみ

3. **開発時の注意事項**
   - 新機能追加時も既存の脆弱性を維持すること
   - セキュリティ学習の観点を常に意識すること

4. **データベース操作**
   - 生のSQLクエリを使用（意図的に脆弱）
   - プリペアドステートメントは使用しない

5. **認証トークン**
   - JWTトークンは簡易的な実装
   - セキュリティよりも学習のしやすさを優先

6. **ログ出力**
   - デバッグ情報を詳細に出力
   - 本番環境では適切なログレベルの設定が必要
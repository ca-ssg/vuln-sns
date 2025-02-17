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

- SQLインジェクション
- XSS（クロスサイトスクリプティング）
- CSRF（クロスサイトリクエストフォージェリ）
- アクセス制御の不備
- セッション管理の不備

**注意**: このアプリケーションは学習目的で作成されています。本番環境での使用や、実際のサービスへの攻撃に使用することは絶対に避けてください。

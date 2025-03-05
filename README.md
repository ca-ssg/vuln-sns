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

```markdown
# ID/パスワード
- alice/alice
- bob/bob
- charlie/charlie
```

## セットアップ

```bash
# リポジトリのクローン
git clone https://github.com/ca-ssg/devin-vuln-app.git
cd devin-vuln-app

# アプリケーションの起動
docker-compose up -d

# フロントエンド: http://localhost:5173
# バックエンドAPI: http://localhost:9090

# データベースへの接続
# 認証情報は docker-compose.yml の環境変数で管理されています
docker-compose exec db mysql -u root -p
```

## プロジェクト構成

```
.
├── frontend/          # Vue3フロントエンド
├── backend/           # Goバックエンド
├── docker-compose.yml # コンテナオーケストレーション
├── docs/             # ドキュメント
└── README.md
```

## データベースへの接続

データベースへの接続には以下のコマンドを使用します：

```bash
docker-compose exec db mysql -uroot -ppassword -Dvuln_app  --default-character-set=utf8mb4
```

認証情報:
- ユーザー名: root
- パスワード: password
- データベース名: vuln_app

※ローカル開発環境のため、簡易的な認証情報を使用しています。本番環境では適切なセキュリティ設定を行ってください。

## ⚠️ 脆弱性について

このアプリケーションには学習目的で意図的に脆弱性が実装されています。
脆弱性の詳細と確認手順については、[脆弱性の概要](docs/vulnerabilities.md)を参照してください。

**注意**: このアプリケーションは学習目的で作成されています。本番環境での使用や、実際のサービスへの攻撃に使用することは絶対に避けてください。

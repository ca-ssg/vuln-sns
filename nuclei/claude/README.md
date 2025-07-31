# Nuclei Templates for Vuln SNS

このディレクトリには、Vuln SNSアプリケーションの脆弱性をテストするためのNucleiテンプレートが含まれています。

## 概要

これらのテンプレートは、意図的に脆弱性を含むSNSアプリケーションの各エンドポイントに対して、様々なセキュリティテストを実行します。

### テンプレートカテゴリ

- **avatar/**: アバターアップロード機能の脆弱性テスト
- **common/**: 共通的なセキュリティ設定のテスト
- **health/**: ヘルスチェックエンドポイントの情報漏洩テスト
- **login/**: ログイン機能の脆弱性テスト（SQLインジェクション、ブルートフォース等）
- **posts/**: 投稿機能の脆弱性テスト（XSS、SQLインジェクション、コマンドインジェクション等）
- **posts-delete/**: 投稿削除機能の認可バイパステスト
- **posts-like/**: いいね機能の認可バイパステスト
- **posts-public/**: 公開投稿エンドポイントの脆弱性テスト
- **posts-update/**: 投稿更新機能の認可バイパステスト
- **profile/**: プロファイル機能の脆弱性テスト
- **profile-update/**: プロファイル更新機能のSQLインジェクションテスト
- **search/**: 検索機能の脆弱性テスト

## 前提条件

1. Nucleiのインストール
```bash
# Go経由でインストール
go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest

# Homebrewでインストール（macOS）
brew install nuclei

# バージョン確認
nuclei -version
```

2. Vuln SNSアプリケーションの起動
```bash
# プロジェクトルートで実行
docker-compose up -d
# または
make start
```

## 使用方法

### 基本的な使用方法

```bash
# 単一のテンプレートを実行
nuclei -t nuclei/claude/login/sql-injection.yaml -u http://localhost:9090

# カテゴリ全体を実行
nuclei -t nuclei/claude/login/ -u http://localhost:9090

# すべてのテンプレートを実行
nuclei -t nuclei/claude/ -u http://localhost:9090
```

### 推奨される実行順序

1. **共通テスト**（セキュリティヘッダー、HTTPS設定など）
```bash
nuclei -t nuclei/claude/common/ -u http://localhost:9090
```

2. **認証関連のテスト**
```bash
nuclei -t nuclei/claude/login/ -u http://localhost:9090
```

3. **各機能のテスト**
```bash
# 投稿機能
nuclei -t nuclei/claude/posts/ -u http://localhost:9090

# プロファイル機能
nuclei -t nuclei/claude/profile/ -u http://localhost:9090

# 検索機能
nuclei -t nuclei/claude/search/ -u http://localhost:9090
```

### 詳細な出力オプション

```bash
# 詳細な出力を表示
nuclei -t nuclei/claude/ -u http://localhost:9090 -v

# JSON形式で出力
nuclei -t nuclei/claude/ -u http://localhost:9090 -json -o results.json

# Markdownレポートを生成
nuclei -t nuclei/claude/ -u http://localhost:9090 -markdown-export report/
```

### フィルタリングオプション

```bash
# 特定の深刻度のみテスト
nuclei -t nuclei/claude/ -u http://localhost:9090 -severity critical,high

# 特定のタグでフィルタ
nuclei -t nuclei/claude/ -u http://localhost:9090 -tags sqli
nuclei -t nuclei/claude/ -u http://localhost:9090 -tags xss

# 特定のタグを除外
nuclei -t nuclei/claude/ -u http://localhost:9090 -exclude-tags brute-force
```

## テスト結果の解釈

### 深刻度レベル

- **Critical**: 即座に対処が必要な重大な脆弱性（SQLインジェクション、コマンドインジェクション等）
- **High**: 重要な脆弱性（XSS、認可バイパス等）
- **Medium**: 中程度の脆弱性（情報漏洩、ブルートフォース保護の欠如等）
- **Low**: 軽微な脆弱性（セキュリティヘッダーの欠如等）

### 期待される結果

このアプリケーションは教育目的で意図的に脆弱性を含んでいるため、多くのテストで脆弱性が検出されることが期待されます。

主な検出される脆弱性：
- SQLインジェクション（ログイン、検索、投稿機能）
- XSS（投稿、プロファイル、検索機能）
- コマンドインジェクション（アバターアップロード）
- 認可バイパス（投稿の編集・削除）
- 情報漏洩（エラーメッセージ、デバッグ情報）

## トラブルシューティング

### テンプレートの検証

```bash
# テンプレートの構文を検証
nuclei -t nuclei/claude/login/sql-injection.yaml -validate

# すべてのテンプレートを検証
nuclei -t nuclei/claude/ -validate
```

### 一般的な問題

1. **接続エラー**: アプリケーションが起動していることを確認
```bash
curl http://localhost:9090/api/health
```

2. **認証エラー**: いくつかのテストは認証が必要です。テストアカウント（alice/alice）が利用可能であることを確認

3. **タイムアウト**: Time-basedのSQLインジェクションテストは意図的に遅延を発生させます

## カスタマイズ

### 環境変数の使用

```bash
# ベースURLを変更
nuclei -t nuclei/claude/ -u http://your-target:port -var BaseURL=http://your-target:port
```

### レート制限

```bash
# リクエストレートを制限（1秒あたり10リクエスト）
nuclei -t nuclei/claude/ -u http://localhost:9090 -rate-limit 10
```

## 注意事項

- これらのテンプレートは教育・学習目的でのみ使用してください
- 本番環境や許可のないシステムに対して実行しないでください
- いくつかのテスト（特にSQLインジェクション）はデータベースに影響を与える可能性があります
- テスト後は`make reset`でデータベースをリセットすることを推奨します

## 開発者向け情報

### 新しいテンプレートの追加

1. 適切なカテゴリディレクトリにYAMLファイルを作成
2. [Nucleiテンプレート仕様](https://nuclei.projectdiscovery.io/templating-guide/)に従って記述
3. `nuclei -t your-template.yaml -validate`で検証
4. 実際のターゲットでテスト

### ペイロードの追加

各テンプレートのpayloadsセクションに新しいペイロードを追加できます：

```yaml
payloads:
  injection:
    # 既存のペイロード
    - "' OR '1'='1"
    # 新しいペイロードを追加
    - "' OR '2'='2"
```

## 関連情報

- [Nuclei公式ドキュメント](https://nuclei.projectdiscovery.io/)
- [Vuln SNS脆弱性詳細](../../docs/vulnerabilities.md)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
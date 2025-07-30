# Nuclei Vulnerability Testing Templates for vuln-sns

このディレクトリには、vuln-sns アプリケーションの包括的な脆弱性テストを行うための Nuclei テンプレートが含まれています。

## 概要

全11のAPIエンドポイントに対して、提供された脆弱性診断項目表に基づいて包括的なテンプレートを作成しました。明らかに脆弱な箇所だけでなく、すべてのエンドポイントとパラメータに対して適用可能な脆弱性カテゴリをテストします。

## テンプレート構成

### Health Endpoint Templates
- `health/security-headers.yaml` - セキュリティヘッダーの欠落
- `health/error-disclosure.yaml` - 情報漏洩
- `health/https-enforcement.yaml` - HTTPS強制

### Search Endpoint Templates  
- `search/sql-injection.yaml` - tagパラメータ経由のSQLインジェクション
- `search/parameter-tampering.yaml` - パラメータ改ざん
- `search/xss-reflected.yaml` - tagパラメータ経由の反射型XSS
- `search/directory-traversal.yaml` - ディレクトリトラバーサル

### Posts Endpoint Templates
- `posts/sql-injection-get.yaml` - GET /api/posts でのSQLインジェクション
- `posts/sql-injection-post.yaml` - POST /api/posts でのSQLインジェクション
- `posts/sql-injection-put.yaml` - PUT /api/posts/:id でのSQLインジェクション
- `posts/xss-stored.yaml` - 投稿コンテンツでの格納型XSS
- `posts/authorization-bypass.yaml` - 認可制御の回避
- `posts/parameter-tampering.yaml` - IDパラメータ改ざん
- `posts/csrf.yaml` - CSRF保護テスト

### Login Endpoint Templates
- `login/sql-injection.yaml` - ログインでのSQLインジェクション
- `login/brute-force-protection.yaml` - ブルートフォース保護
- `login/weak-session-id.yaml` - 弱いトークン生成
- `login/credential-stuffing.yaml` - クレデンシャルスタッフィング攻撃

### Profile Endpoint Templates
- `profile/command-injection.yaml` - アバターアップロードでのOSコマンドインジェクション
- `profile/xss-reflected.yaml` - プロフィールニックネームでのXSS
- `profile/file-upload-bypass.yaml` - ファイルアップロード制限回避
- `profile/authorization-bypass.yaml` - プロフィールアクセス制御

### General Security Templates
- `general/session-management.yaml` - セッション処理の問題
- `general/cors-misconfiguration.yaml` - CORSポリシーテスト
- `general/http-methods.yaml` - HTTPメソッドテスト
- `general/rate-limiting.yaml` - レート制限チェック

## 使用方法

### 全テンプレートの検証
```bash
nuclei -validate -t ./nuclei/claude/
```

### 実行中のアプリケーションに対するテスト
```bash
# アプリケーションを起動
docker-compose up -d

# テンプレートを実行
nuclei -t ./nuclei/claude/ -u http://localhost:9090 -o results.txt
```

### 特定カテゴリのテスト
```bash
# SQLインジェクションテストのみ
nuclei -t ./nuclei/claude/*/sql-injection*.yaml -u http://localhost:9090

# 認可回避テストのみ  
nuclei -t ./nuclei/claude/*/authorization-bypass.yaml -u http://localhost:9090
```

## テスト結果

最新のテスト実行では42件の脆弱性マッチが検出されました：

### Critical (重大)
- SQLインジェクション (ログイン)
- OSコマンドインジェクション (プロフィール)

### High (高)
- SQLインジェクション (投稿作成・更新)
- 認可制御回避 (投稿・プロフィール)
- ファイルアップロード制限回避

### Medium (中)
- 弱いセッションID
- パラメータ改ざん
- レート制限なし
- セッション管理の問題

## カバレッジマトリックス

| エンドポイント | SQL注入 | XSS | コマンド注入 | 認可回避 | パラメータ改ざん | その他 |
|---------------|---------|-----|-------------|----------|----------------|--------|
| /api/health   | -       | -   | -           | -        | -              | ✓      |
| /api/search   | ✓       | ✓   | -           | -        | ✓              | ✓      |
| /api/posts    | ✓       | ✓   | -           | ✓        | ✓              | ✓      |
| /api/login    | ✓       | -   | -           | -        | -              | ✓      |
| /api/profile  | -       | ✓   | ✓           | ✓        | -              | ✓      |

## 注意事項

- これらのテンプレートは教育目的で作成されており、実際の脆弱性を含むアプリケーションに対してテストを行います
- 本番環境では使用しないでください
- テスト実行前にアプリケーションが正常に動作していることを確認してください

## 技術詳細

- Nuclei バージョン: 3.4.7
- テンプレート数: 25+
- 対象エンドポイント: 11
- 脆弱性カテゴリ: 20+

各テンプレートは実際の実装に基づいて作成されており、MySQL固有のSQLインジェクションペイロード、特定のscanFile関数実装を対象としたコマンドインジェクション、弱いトークン検証ロジックを悪用する認証回避などが含まれています。

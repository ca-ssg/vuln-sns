# 脆弱性の概要

このアプリケーションには、セキュリティ学習を目的として意図的に以下の脆弱性が実装されています。

## 認証機能の脆弱性
- [ログイン機能のSQLインジェクション](auth/sqli-login.md)
- [認証トークンの露出](auth/auth-token-exposure.md)

## 投稿機能の脆弱性
- [投稿機能のSQLインジェクション](post/sqli-post.md)
- [投稿表示機能のXSS](post/xss-post.md)

## プロフィール機能の脆弱性
- [プロフィール機能のSQLインジェクション](profile/sqli-profile.md)
- [プロフィール表示機能のXSS](profile/xss-profile.md)
- [プロフィール更新機能のCSRF](profile/csrf-profile.md)
- [プロフィール更新機能のSQLインジェクション](profile/injection-profile-update.md)

## 検索機能の脆弱性
- [検索機能のSQLインジェクション](search/sqli-search.md)

**注意**: このアプリケーションは学習目的で作成されています。本番環境での使用や、実際のサービスへの攻撃に使用することは絶対に避けてください。

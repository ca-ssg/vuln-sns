# Sub-Issue作成テンプレート

## 脆弱性の概要
- **脆弱性の種別**: [SQLインジェクション/XSS/OSコマンドインジェクション等]
- **重要度**: [High/Medium/Low]
- **影響範囲**: [具体的な影響内容]

## 検出内容
- **ファイル**: [ファイルパス:行番号]
- **semgrep検出ルール**: [ルール名]
- **脆弱なコード**: 
```[言語]
[コードスニペット]
```

## エンドポイント
- **URL**: [該当するAPIエンドポイントまたはページ]
- **HTTPメソッド**: [GET/POST/PUT/DELETE]
- **認証**: [必要/不要]

## 再現手順
1. [具体的な手順1]
2. [具体的な手順2]
3. [結果の確認方法]

### 攻撃例
```bash
# curlコマンドの例
curl -X POST 'http://localhost:9090/api/endpoint' \
  -H 'Authorization: Bearer token' \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "parameter": "malicious_payload"
  }'
```

## 影響
- **機密性**: [High/Medium/Low] - [具体的な影響]
- **完全性**: [High/Medium/Low] - [具体的な影響]
- **可用性**: [High/Medium/Low] - [具体的な影響]

### 攻撃シナリオ
1. [攻撃者の行動1]
2. [攻撃者の行動2]
3. [最終的な被害]

## 具体的な対策方法

### 修正前のコード
```[言語]
[脆弱なコード]
```

### 修正後のコード
```[言語]
[安全なコード]
```

### 実装における注意点
- [注意点1]
- [注意点2]
- [注意点3]

### 追加の対策
- [セキュリティヘッダーの設定]
- [入力値検証の強化]
- [ログ監視の実装]

## テスト方法
### 修正前の確認
```bash
# 脆弱性が存在することを確認するコマンド
```

### 修正後の確認
```bash
# 脆弱性が修正されたことを確認するコマンド
```

## リファレンス
- **親issue**: #40
- **CWE情報**: [CWE-XXX: 脆弱性名](https://cwe.mitre.org/data/definitions/XXX.html)
- **OWASP**: [該当するOWASP項目のURL]
- **検出日時**: [YYYY-MM-DD HH:MM:SS]
- **検出ツール**: semgrep v1.131.0

## ラベル
- `SAST`
- `devin`
- `security`
- `[脆弱性種別]` (例: `sql-injection`, `xss`, `command-injection`)
- `[重要度]` (例: `high-priority`, `medium-priority`, `low-priority`)

## チェックリスト
- [ ] 脆弱性の再現確認完了
- [ ] 対策方法の検証完了
- [ ] ドキュメントの更新完了
- [ ] テストケースの作成完了
- [ ] コードレビューの実施完了
- [ ] 修正の動作確認完了

## 備考
[その他の特記事項や補足情報]

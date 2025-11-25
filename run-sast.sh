#!/bin/bash

echo "=== SAST診断開始 ==="
echo "実行日時: $(date)"
echo "対象リポジトリ: ca-ssg/vuln-sns"
echo ""

echo "=== semgrepスキャン実行 ==="
semgrep scan --config=p/default --verbose --output=sast-results.json --json

echo ""
echo "=== スキャン結果サマリー ==="
semgrep scan --config=p/default --quiet

echo ""
echo "=== 診断完了 ==="
echo "詳細結果: sast-results.json"
echo "診断レポート: docs/sast-diagnosis.md"
echo "チェックリスト: docs/sast-checklist.md"
echo "sub-issue作成テンプレート: docs/sub-issue-template.md"
echo ""
echo "=== 次のステップ ==="
echo "1. docs/sast-diagnosis.mdで詳細な診断結果を確認"
echo "2. docs/sub-issue-template.mdを使用してsub-issueを作成"
echo "3. 各脆弱性の対策を実装"
echo "4. 修正後に再度SASTスキャンを実行して改善を確認"

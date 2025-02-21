# プロフィール表示機能のXSS

## 脆弱性の説明
プロフィールのニックネーム表示で、ユーザー入力を適切にエスケープせずに表示しているため、クロスサイトスクリプティング（XSS）が可能です。

## 影響範囲
- 悪意のあるスクリプトの実行
- ユーザーセッションの盗取
- 不正な操作の実行

## 確認手順
1. プロフィール更新フォームにJavaScriptコードを含むHTMLを入力
2. プロフィールページを表示し、スクリプトが実行されることを確認

## 対策方法
### 1. エスケープ処理の実施
```javascript
// ユーティリティ関数の作成
function escapeHtml(text) {
  return text
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#039;");
}

// コンポーネントでの使用
export default {
  computed: {
    safeNickname() {
      return escapeHtml(this.user.nickname);
    }
  }
}
```

### 2. v-textディレクティブの使用
```vue
<!-- 修正前（脆弱なコード） -->
<template>
  <div>{{ user.nickname }}</div>
</template>

<!-- 修正後（安全なコード） -->
<template>
  <div v-text="safeNickname"></div>
</template>
```

# クロスサイトスクリプティング（XSS）

## 脆弱性の説明
XSSは、Webアプリケーションにスクリプトを注入できる脆弱性です。
攻撃者は、この脆弱性を利用して、ユーザーのブラウザ上で悪意のあるスクリプトを実行できます。

## 実装箇所
1. 投稿内容の表示 (`frontend/src/components/PostCard.vue`)
```vue
<div v-html="post.content"></div>
```

2. プロフィール表示 (`frontend/src/views/ProfileView.vue`)
```vue
<div v-html="user.nickname"></div>
```

## 確認手順

### 投稿でのXSS
1. ログイン後、以下のようなスクリプトを含む投稿を作成:
```html
<script>alert('XSS!');</script>
```
2. 投稿一覧を表示し、アラートが表示されることを確認

### プロフィールでのXSS
1. プロフィール編集画面で、以下のようなニックネームを設定:
```html
<img src="x" onerror="alert('XSS!')">
```
2. プロフィールページを表示し、アラートが表示されることを確認

## 対策方法

### 1. v-htmlの代わりにv-textを使用
```vue
<!-- 修正前（脆弱なコード） -->
<template>
  <div v-html="post.content"></div>
</template>

<!-- 修正後（安全なコード） -->
<template>
  <div v-text="post.content"></div>
</template>
```

### 2. エスケープ処理の実施
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
    safeContent() {
      return escapeHtml(this.post.content);
    }
  }
}
```

### 3. Content Security Policyの設定
```javascript
// nuxt.config.js または vue.config.js
module.exports = {
  head: {
    meta: [
      {
        hid: 'Content-Security-Policy',
        httpEquiv: 'Content-Security-Policy',
        content: "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline';"
      }
    ]
  }
}
```

### 4. マークダウンやリッチテキストの安全な表示
```vue
<template>
  <div>
    <!-- マークダウンライブラリを使用した安全な表示 -->
    <markdown-it-vue :content="post.content" :options="markdownOptions"/>
  </div>
</template>

<script>
export default {
  data() {
    return {
      markdownOptions: {
        html: false, // HTMLタグの無効化
        linkify: true, // URLの自動リンク化
        breaks: true // 改行の有効化
      }
    }
  }
}
</script>
```

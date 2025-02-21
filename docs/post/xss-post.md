# 投稿表示機能のXSS

## 脆弱性の説明
投稿内容の表示にv-htmlディレクティブを使用しているため、クロスサイトスクリプティング（XSS）が可能です。

## 影響範囲
- 悪意のあるスクリプトの実行
- ユーザーセッションの盗取
- 不正な操作の実行

## 確認手順
1. 投稿作成フォームにJavaScriptコードを含むHTMLを入力
2. 投稿を表示し、スクリプトが実行されることを確認

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

### 2. マークダウンやリッチテキストの安全な表示
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

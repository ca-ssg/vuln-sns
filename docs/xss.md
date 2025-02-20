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
1. 入力値のサニタイズ
2. エスケープ処理の実施
3. Content Security Policyの設定
4. v-htmlの使用を避け、v-textやテンプレート構文を使用

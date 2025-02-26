# 認証トークンの露出

このアプリケーションは学習目的で意図的にセキュリティ脆弱性を含んでいます：

## 概要
- 認証トークンがサーバーログに出力される
- 単純なトークン形式（userID_token）
- トークンがlocalStorageに平文で保存される

## 学習目的
1. ログ出力におけるセキュリティ上の考慮事項の理解
2. 安全なトークン処理の学習
3. 安全なトークン設計の重要性の理解

## 実装箇所
- バックエンド：`backend/internal/middleware/auth.go`
- フロントエンド：`frontend/src/stores/auth.ts`

## 攻撃方法と手順

### 1. ブラウザの開発者ツールを使用したトークン取得

1. アプリケーションにログインします
2. ブラウザの開発者ツールを開きます（Chrome: F12キー、または右クリック→「検証」）
3. 「Application」タブ（Chrome）または「Storage」タブ（Firefox）を選択します
4. 左側のサイドバーから「Local Storage」を展開し、アプリケーションのドメインを選択します
5. 右側のペインに表示される「token」キーの値を確認します

この値は`userID_token`形式の認証トークンで、これを使用して他のユーザーになりすますことができます。

### 2. ネットワークログからのトークン取得

1. アプリケーションにログインします
2. ブラウザの開発者ツールを開きます
3. 「Network」タブを選択します
4. ページを更新するか、APIリクエストが発生する操作を行います
5. リクエストの一覧から任意のAPIリクエストを選択します
6. 「Headers」タブで「Authorization」ヘッダーの値を確認します

この値は`Bearer userID_token`形式になっており、トークンを簡単に抽出できます。

### 3. サーバーログからのトークン取得（サーバーアクセス権がある場合）

サーバーのログファイルには認証トークンが平文で出力されています：

```
Auth header: Bearer user1_token
```

## 攻撃成功の確認手順

### 1. 取得したトークンを使用した認証バイパス

1. 取得したトークン（例：`user1_token`）をメモします
2. ブラウザの開発者ツールを開きます
3. コンソールタブで以下のコマンドを実行します：

```javascript
localStorage.setItem('token', 'user1_token');
```

4. ページを更新します
5. 正常にログインされ、user1のアカウントとしてアプリケーションにアクセスできることを確認します

### 2. APIリクエストへの直接トークン挿入

1. Postmanなどのツールを使用して、APIエンドポイント（例：`/api/profile`）にGETリクエストを送信します
2. Authorizationヘッダーに`Bearer user1_token`を設定します
3. リクエストを送信し、正常にレスポンスが返ってくることを確認します

これにより、認証トークンが簡単に取得・再利用可能であることが確認できます。

## 対策方法
本番環境では以下の対策が必要：

### 1. トークンのログ出力の無効化
```go
// 修正前
log.Printf("Auth header: %s", authHeader)
log.Printf("Invalid token format: %s", authHeader)

// 修正後
log.Printf("Authentication attempt for request: %s %s", c.Request.Method, c.Request.URL.Path)
if err != nil {
    log.Printf("Authentication failed: %v", err)
}
```

### 2. JWTの採用
```go
// 修正前
token := userID + "_token"

// 修正後
claims := jwt.MapClaims{
    "user_id": userID,
    "exp":     time.Now().Add(time.Hour * 24).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
```

### 3. トークンの安全な保存
```typescript
// 修正前 (frontend/src/stores/auth.ts)
localStorage.setItem('token', token)

// 修正後
// セッションストレージを使用して保存（ブラウザを閉じると消去される）
sessionStorage.setItem('token', token)
```

### 4. セキュアなヘッダーの設定
```go
// 修正前
c.Header("Access-Control-Allow-Headers", "*")

// 修正後
c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
c.Header("X-Content-Type-Options", "nosniff")
c.Header("X-Frame-Options", "DENY")
```

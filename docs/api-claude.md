# API エンドポイント一覧

## 認証不要エンドポイント

### GET /api/health
**説明**: ヘルスチェック

**レスポンス**:
```json
{
  "status": "ok"
}
```

---

### POST /api/login
**説明**: ユーザーログイン

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| user_id | string | ✓ | ユーザーID |
| password | string | ✓ | パスワード |

**レスポンス**:
```json
{
  "token": "userID_token",
  "user": {
    "id": "string",
    "nickname": "string",
    "avatar_data": "string"
  }
}
```

---

### GET /api/posts
**説明**: 投稿一覧を取得

**リクエストパラメータ**: なし

**レスポンス**:
```json
[
  {
    "id": "number",
    "userId": "string",
    "content": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime",
    "likes": "number",
    "isLiked": "boolean"
  }
]
```

---

### GET /api/search
**説明**: ハッシュタグで投稿を検索

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| tag | string | ✓ | 検索するハッシュタグ（クエリパラメータ） |

**レスポンス**:
```json
[
  {
    "id": "number",
    "userId": "string",
    "content": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime",
    "likes": "number",
    "isLiked": "boolean"
  }
]
```

---

## 認証必要エンドポイント

すべての認証必要エンドポイントは、リクエストヘッダーに以下を含む必要があります：
| ヘッダー名 | 型 | 必須 | 説明 |
|------------|----|----|------|
| Authorization | string | ✓ | Bearer {token} 形式の認証トークン |

### POST /api/posts
**説明**: 新規投稿を作成

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| content | string | ✓ | 投稿内容 |

**レスポンス**:
```json
{
  "id": "number",
  "userId": "string",
  "content": "string",
  "createdAt": "datetime",
  "updatedAt": "datetime"
}
```

---

### PUT /api/posts/:id
**説明**: 投稿を更新

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| content | string | ✓ | 更新後の投稿内容 |

**レスポンス**:
```json
{
  "message": "Post updated successfully"
}
```

---

### DELETE /api/posts/:id
**説明**: 投稿を削除

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**レスポンス**:
```json
{
  "message": "Post deleted successfully"
}
```

---

### POST /api/posts/:id/like
**説明**: 投稿にいいねを付ける

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "message": "Post liked successfully"
}
```
または
```json
{
  "message": "Post already liked"
}
```

---

### DELETE /api/posts/:id/like
**説明**: 投稿のいいねを取り消す

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**レスポンス**:
```json
{
  "message": "Post unliked successfully"
}
```

---

### GET /api/profile
**説明**: ユーザープロフィール情報を取得

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "id": "string",
  "nickname": "string",
  "avatar_data": "string"
}
```

---

### PUT /api/profile
**説明**: プロフィール情報を更新

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| nickname | string | ✓ | 新しいニックネーム |

**レスポンス**:
```json
{
  "message": "Profile updated successfully"
}
```

---

### POST /api/profile/avatar
**説明**: アバター画像をアップロード

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| file_id | string | ✓ | ファイルID |
| image_data | string | ✓ | Base64エンコードされた画像データ |

**レスポンス**:
```json
{
  "malware_scan": "string",
  "avatar_data": "string"
}
```

---

## エラーレスポンス

すべてのエンドポイントは、エラー時に以下の形式でレスポンスを返します：

```json
{
  "error": "エラーメッセージ"
}
```

### 一般的なHTTPステータスコード
- 200: 成功
- 201: 作成成功
- 400: 不正なリクエスト
- 401: 認証エラー
- 404: リソースが見つからない
- 500: サーバー内部エラー
# API エンドポイント一覧

## 概要
ca-ssg/vuln-snsリポジトリのAPIエンドポイント仕様書です。
本APIは `/api` 配下で提供され、認証が必要なエンドポイントはBearerトークンによる認証を採用しています。

## 認証方法

### ログイン
ログインエンドポイント（`POST /api/login`）でユーザーIDとパスワードを送信し、認証トークンを取得します。

### 認証トークンの使用方法
認証が必要なエンドポイントでは、リクエストヘッダーに以下を含める必要があります：

| ヘッダー名 | 型 | 必須 | 説明 |
|-----------|----|----|------|
| Authorization | string | ✓ | Bearer {user_id}_token 形式の認証トークン |

例：
```
Authorization: Bearer alice_token
```

---

## パブリックエンドポイント（認証不要）

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

**エラーレスポンス**:
```json
{
  "error": "Invalid credentials"
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
    "id": 123,
    "userId": "string",
    "content": "string",
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "likes": 10,
    "isLiked": false
  }
]
```

---

### GET /api/search
**説明**: ハッシュタグで投稿を検索

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| tag | string | ✓ | 検索するハッシュタグ（URLクエリパラメータ） |

**レスポンス**:
```json
[
  {
    "id": 123,
    "userId": "string",
    "content": "string",
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z",
    "likes": 10,
    "isLiked": false
  }
]
```

**エラーレスポンス**:
```json
{
  "error": "Hashtag parameter is required"
}
```

---

### GET /api/health
**説明**: ヘルスチェック

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "status": "ok"
}
```

---

## 認証必要エンドポイント

すべての認証必要エンドポイントは、リクエストヘッダーに以下を含む必要があります：

| ヘッダー名 | 型 | 必須 | 説明 |
|-----------|----|----|------|
| Authorization | string | ✓ | Bearer {user_id}_token 形式の認証トークン |

### POST /api/posts
**説明**: 新規投稿を作成

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| content | string | ✓ | 投稿内容 |

**レスポンス**:
```json
{
  "id": 123,
  "userId": "string",
  "content": "string",
  "createdAt": "0001-01-01T00:00:00Z",
  "updatedAt": "0001-01-01T00:00:00Z",
  "likes": 0,
  "isLiked": false
}
```

**エラーレスポンス**:
```json
{
  "error": "Failed to create post"
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

**エラーレスポンス**:
```json
{
  "error": "Post not found or unauthorized"
}
```

---

### DELETE /api/posts/:id
**説明**: 投稿を削除

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "message": "Post deleted successfully"
}
```

**エラーレスポンス**:
```json
{
  "error": "Post not found or unauthorized"
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

**エラーレスポンス**:
```json
{
  "error": "Post not found"
}
```

---

### DELETE /api/posts/:id/like
**説明**: 投稿のいいねを取り消す

**URLパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| id | string | ✓ | 投稿ID |

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "message": "Post unliked successfully"
}
```

**エラーレスポンス**:
```json
{
  "error": "Like not found"
}
```

---

### GET /api/profile
**説明**: ユーザープロフィールを取得

**リクエストパラメータ**: なし

**レスポンス**:
```json
{
  "id": "string",
  "nickname": "string",
  "avatar_data": "string"
}
```

**エラーレスポンス**:
```json
{
  "error": "Failed to get profile"
}
```

---

### PUT /api/profile
**説明**: プロフィールを更新

**リクエストパラメータ**:
| パラメータ名 | 型 | 必須 | 説明 |
|-------------|----|----|------|
| nickname | string | ✓ | ニックネーム |

**レスポンス**:
```json
{
  "message": "Profile updated successfully"
}
```

**エラーレスポンス**:
```json
{
  "error": "Failed to update profile"
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

**エラーレスポンス**:
```json
{
  "error": "Virus detected or scan failed",
  "scan_result": "string"
}
```

---

## エラーレスポンス共通仕様

すべてのエンドポイントで共通のエラーレスポンス形式：

### 認証エラー (401)
```json
{
  "error": "Authorization header required"
}
```

### バリデーションエラー (400)
```json
{
  "error": "Invalid request format"
}
```

### サーバーエラー (500)
```json
{
  "error": "Internal server error"
}
```
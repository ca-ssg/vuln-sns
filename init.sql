CREATE DATABASE IF NOT EXISTS vuln_app CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE vuln_app;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    avatar_data LONGTEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS likes (
    post_id INT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT IGNORE INTO users (id, password, nickname) 
VALUES ('alice', SHA2('alice', 256), 'Alice'),
       ('bob', SHA2('bob', 256), 'Bob'),
       ('charlie', SHA2('charlie', 256), 'Charlie');

INSERT IGNORE INTO posts (user_id, content, created_at, updated_at) VALUES 
('alice', '初めての投稿です！ #初投稿', NOW(), NOW()),
('alice', 'セキュリティの基本について勉強中 #セキュリティ #勉強', NOW(), NOW()),
('bob', 'SQLインジェクションの対策方法 #セキュリティ #SQL', NOW(), NOW()),
('charlie', 'クロスサイトスクリプティングについて #XSS #セキュリティ', NOW(), NOW()),
('alice', 'パスワード管理の重要性 #セキュリティ #パスワード', NOW(), NOW()),
('bob', 'ファイアウォールの設定方法 #ネットワーク #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ監査のポイント #監査 #セキュリティ', NOW(), NOW()),
('alice', '多要素認証について #認証 #セキュリティ', NOW(), NOW()),
('bob', 'ゼロトラストセキュリティとは #ゼロトラスト #セキュリティ', NOW(), NOW()),
('charlie', 'ペネトレーションテストの方法 #ペネトレーション #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティインシデント対応計画 #インシデント #セキュリティ', NOW(), NOW()),
('bob', 'クラウドセキュリティの課題 #クラウド #セキュリティ', NOW(), NOW()),
('charlie', 'IoTデバイスのセキュリティリスク #IoT #セキュリティ', NOW(), NOW()),
('alice', 'ソーシャルエンジニアリングの手口 #ソーシャルエンジニアリング #セキュリティ', NOW(), NOW()),
('bob', 'ランサムウェア対策について #ランサムウェア #セキュリティ', NOW(), NOW()),
('charlie', 'セキュアコーディングの原則 #コーディング #セキュリティ', NOW(), NOW()),
('alice', 'VPNの選び方 #VPN #セキュリティ', NOW(), NOW()),
('bob', 'データ暗号化の重要性 #暗号化 #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ意識向上トレーニング #トレーニング #セキュリティ', NOW(), NOW()),
('alice', 'ネットワークセキュリティの基本 #ネットワーク #セキュリティ', NOW(), NOW()),
('bob', 'マルウェア分析の手法 #マルウェア #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティパッチ管理の重要性 #パッチ管理 #セキュリティ', NOW(), NOW()),
('alice', 'フィッシング詐欺の見分け方 #フィッシング #セキュリティ', NOW(), NOW()),
('bob', 'ブルートフォース攻撃対策 #ブルートフォース #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティフレームワークの比較 #フレームワーク #セキュリティ', NOW(), NOW()),
('alice', 'DDoS攻撃とその対策 #DDoS #セキュリティ', NOW(), NOW()),
('bob', 'エンドポイントセキュリティの重要性 #エンドポイント #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ監視ツールの選び方 #監視 #セキュリティ', NOW(), NOW()),
('alice', 'ゼロデイ脆弱性について #ゼロデイ #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ認証資格の種類 #認証資格 #セキュリティ', NOW(), NOW()),
('charlie', 'コンテナセキュリティの課題 #コンテナ #セキュリティ', NOW(), NOW()),
('alice', 'APIセキュリティのベストプラクティス #API #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ自動化ツール #自動化 #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドネイティブセキュリティ #クラウドネイティブ #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ設計の原則 #設計 #セキュリティ', NOW(), NOW()),
('bob', 'デジタルフォレンジック入門 #フォレンジック #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティコンプライアンスの重要性 #コンプライアンス #セキュリティ', NOW(), NOW()),
('alice', 'ハニーポットの設置方法 #ハニーポット #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ情報共有の重要性 #情報共有 #セキュリティ', NOW(), NOW()),
('charlie', 'DevSecOpsの実践方法 #DevSecOps #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティリスク評価の方法 #リスク評価 #セキュリティ', NOW(), NOW()),
('bob', 'ウェブアプリケーションファイアウォール #WAF #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ文化の醸成 #文化 #セキュリティ', NOW(), NOW()),
('alice', 'サイバー保険の選び方 #サイバー保険 #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティインシデントの分類 #インシデント #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドセキュリティアーキテクチャ #クラウド #アーキテクチャ #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ対策の費用対効果 #費用対効果 #セキュリティ', NOW(), NOW()),
('bob', 'サプライチェーンセキュリティ #サプライチェーン #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ意識向上キャンペーン #意識向上 #セキュリティ', NOW(), NOW()),
('alice', 'リモートワークのセキュリティ対策 #リモートワーク #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ監査の準備方法 #監査 #セキュリティ', NOW(), NOW()),
('charlie', 'ゼロトラストネットワークの実装 #ゼロトラスト #ネットワーク #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティポリシーの作成方法 #ポリシー #セキュリティ', NOW(), NOW()),
('bob', 'クラウドストレージのセキュリティ #クラウドストレージ #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ教育プログラムの設計 #教育 #セキュリティ', NOW(), NOW()),
('alice', 'モバイルデバイス管理のセキュリティ #MDM #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対応チームの構築 #CSIRT #セキュリティ', NOW(), NOW()),
('charlie', 'ブロックチェーンセキュリティの課題 #ブロックチェーン #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ脅威インテリジェンス #脅威インテリジェンス #セキュリティ', NOW(), NOW()),
('bob', '量子コンピューティングとセキュリティ #量子コンピューティング #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ自動化の利点 #自動化 #セキュリティ', NOW(), NOW()),
('alice', 'データプライバシーとセキュリティ #プライバシー #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対策のROI計算 #ROI #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドセキュリティベストプラクティス #クラウド #ベストプラクティス #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ認証の種類と比較 #認証 #セキュリティ', NOW(), NOW()),
('bob', 'サイバーセキュリティフレームワーク #フレームワーク #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ対策の自動化ツール #自動化 #ツール #セキュリティ', NOW(), NOW()),
('alice', 'ネットワークセグメンテーション #セグメンテーション #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対応プレイブック #プレイブック #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドネイティブアプリケーションのセキュリティ #クラウドネイティブ #アプリケーション #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ監視ダッシュボード #監視 #ダッシュボード #セキュリティ', NOW(), NOW()),
('bob', 'コンテナオーケストレーションのセキュリティ #コンテナ #オーケストレーション #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ対策の優先順位付け #優先順位 #セキュリティ', NOW(), NOW()),
('alice', 'クラウドセキュリティポスチャー管理 #CSPM #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対応の自動化 #自動化 #対応 #セキュリティ', NOW(), NOW()),
('charlie', 'マイクロサービスセキュリティ #マイクロサービス #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ対策の効果測定 #効果測定 #セキュリティ', NOW(), NOW()),
('bob', 'クラウドアクセスセキュリティブローカー #CASB #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ対応の効率化 #効率化 #セキュリティ', NOW(), NOW()),
('alice', 'サーバーレスセキュリティ #サーバーレス #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対策のコスト削減 #コスト削減 #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドワークロードプロテクション #ワークロード #プロテクション #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ対応の自動化ツール #自動化 #ツール #セキュリティ', NOW(), NOW()),
('bob', 'コンテナイメージスキャン #コンテナ #イメージスキャン #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ対策の可視化 #可視化 #セキュリティ', NOW(), NOW()),
('alice', 'クラウドセキュリティ監査 #クラウド #監査 #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対応の自動化事例 #自動化 #事例 #セキュリティ', NOW(), NOW()),
('charlie', 'マルチクラウドセキュリティ #マルチクラウド #セキュリティ', NOW(), NOW()),
('alice', 'セキュリティ対策のベンチマーク #ベンチマーク #セキュリティ', NOW(), NOW()),
('bob', 'クラウドネイティブセキュリティツール #クラウドネイティブ #ツール #セキュリティ', NOW(), NOW()),
('charlie', 'セキュリティ対応の効率化事例 #効率化 #事例 #セキュリティ', NOW(), NOW()),
('alice', 'コンテナセキュリティベストプラクティス #コンテナ #ベストプラクティス #セキュリティ', NOW(), NOW()),
('bob', 'セキュリティ対策の自動化戦略 #自動化 #戦略 #セキュリティ', NOW(), NOW()),
('charlie', 'クラウドセキュリティフレームワーク #クラウド #フレームワーク #セキュリティ', NOW(), NOW()),
('alice', '#セキュリティ の学習頑張ります！', NOW(), NOW()),
('bob', '新しい #脆弱性 が見つかりました', NOW(), NOW()),
('charlie', '#セキュリティ 大好き芸人です', NOW(), NOW());


-- ランダムにいいねデータを挿入
INSERT INTO likes (post_id, user_id)
SELECT p.id, u.id
FROM posts p
CROSS JOIN users u
WHERE (p.id + ASCII(SUBSTRING(u.id, 1, 1))) % 3 = 0
AND p.user_id != u.id
LIMIT 50;

CREATE DATABASE IF NOT EXISTS vuln_app;
USE vuln_app;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS likes (
    post_id INT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT IGNORE INTO users (id, password) VALUES ('alice', SHA2('alice', 256));

INSERT IGNORE INTO posts (user_id, content, created_at) VALUES 
('alice', '初めての投稿です！', NOW()),
('alice', 'セキュリティの学習頑張ります！', NOW());

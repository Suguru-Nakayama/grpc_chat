BEGIN;

CREATE TABLE IF NOT EXISTS grpc_chat.users(
    `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'サロゲートキー',
    `last_name` VARCHAR(25) NOT NULL COMMENT 'ユーザー姓',
    `first_name` VARCHAR(25) NOT NULL COMMENT 'ユーザー名',
    `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    `password` VARCHAR(255) NOT NULL COMMENT 'パスワード',
    `crated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

COMMIT;
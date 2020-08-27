BEGIN;

-- チャットルームテーブル
CREATE TABLE IF NOT EXISTS grpc_chat.chat_rooms (
    `chat_room_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'サロゲートキー',
    `crated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`chat_room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャットルームテーブル';

-- チャットメンバーテーブル
CREATE TABLE IF NOT EXISTS grpc_chat.chat_members (
    `chat_member_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'サロゲートキー',
    `chat_room_id` INT(11) UNSIGNED NOT NULL COMMENT 'チャットルームID',
    `user_id` INT(11) UNSIGNED NOT NULL COMMENT 'チャットルーム参加者のユーザーID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`chat_member_id`),
    CONSTRAINT `fk_chat_member_room_id` FOREIGN KEY (`chat_room_id`) REFERENCES `chat_rooms` (`chat_room_id`),
    CONSTRAINT `fk_chat_member_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
    UNIQUE uk_member_user_id_and_room_id (user_id, chat_room_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャットルームメンバーテーブル';

-- チャットメッセージテーブル
CREATE TABLE IF NOT EXISTS grpc_chat.chat_messages (
    `chat_message_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'サロゲートキー',
    `chat_room_id` INT(11) UNSIGNED NOT NULL COMMENT 'チャットルームID',
    `user_id` INT(11) UNSIGNED NOT NULL COMMENT 'メッセージを送信したユーザーID',
    `type` TINYINT(2) UNSIGNED NOT NULL COMMENT 'メッセージ種別（1: テキストメッセージ, 2: 画像, 3: ファイル）',
    `text` VARCHAR(255) DEFAULT NULL COMMENT 'メッセージ本文',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`chat_message_id`),
    CONSTRAINT `fk_chat_message_room_id` FOREIGN KEY (`chat_room_id`) REFERENCES `chat_rooms` (`chat_room_id`),
    CONSTRAINT `fk_chat_message_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャットメッセージテーブル';

-- チャットメッセージ未読テーブル
CREATE TABLE IF NOT EXISTS grpc_chat.chat_messages_unread (
    `chat_message_unread_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'サロゲートキー',
    `user_id` INT(11) UNSIGNED NOT NULL COMMENT 'ユーザーID',
    `chat_message_id` BIGINT(20) UNSIGNED NOT NULL COMMENT 'チャットメッセージID',
    `chat_room_id` INT(11) UNSIGNED NOT NULL COMMENT 'チャットルームID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    `updated_at` DATETIME DEFAULT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`chat_message_unread_id`),
    CONSTRAINT `fk_chat_message_unread_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
    CONSTRAINT `fk_chat_message_unread_message_id` FOREIGN KEY (`chat_message_id`) REFERENCES `chat_messages` (`chat_message_id`),
    CONSTRAINT `fk_chat_message_unread_room_id` FOREIGN KEY (`chat_room_id`) REFERENCES `chat_rooms` (`chat_room_id`),
    UNIQUE KEY `uk_unread_user_id_and_message_id` (`user_id`,`chat_message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='チャットメッセージ未読テーブル';

COMMIT;
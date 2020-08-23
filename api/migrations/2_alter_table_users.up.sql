BEGIN;

ALTER TABLE grpc_chat.users
    DROP COLUMN `email`,
    DROP COLUMN `password`,
    ADD COLUMN `firebase_user_id` VARCHAR(255) NOT NULL COMMENT 'FirebaseユーザーID' AFTER first_name;

COMMIT;
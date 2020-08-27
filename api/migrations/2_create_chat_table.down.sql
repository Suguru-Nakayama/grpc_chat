BEGIN;

DROP TABLE IF EXISTS grpc_chat.chat_messages_unread;
DROP TABLE IF EXISTS grpc_chat.chat_messages;
DROP TABLE IF EXISTS grpc_chat.chat_members;
DROP TABLE IF EXISTS grpc_chat.chat_rooms;

COMMIT;
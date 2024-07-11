-- +goose Up
-- +goose Up
CREATE TABLE IF NOT EXISTS `messages`
(
    `uuid`                     varchar(255) NOT NULL PRIMARY KEY,
    `content_type`             varchar(255) NOT NULL,
    `message_content_send_at`  TIMESTAMP,
    `message_content_provider` varchar(255),
    `message_content_consumer` varchar(255),
    `message_content_title`    varchar(255),
    `message_content_content`  TEXT
);

-- +goose Down
DROP TABLE `messages`;
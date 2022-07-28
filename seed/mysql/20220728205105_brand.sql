-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `brand`
(
    `id`         BIGINT unsigned NOT NULL,
    `name`       VARCHAR(50)     NOT NULL,
    `created_at` DATETIME(6)     NOT NULL DEFAULT current_timestamp(),
    `updated_at` DATETIME(6)     NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `is_deleted` BOOLEAN         NOT NULL DEFAULT 0,
    `deleted_at` DATETIME(6)     NULL     DEFAULT NULL,

    PRIMARY KEY (`id`),
    UNIQUE (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `brand`;
-- +goose StatementEnd

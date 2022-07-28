-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `product`
(
    `id`         BIGINT unsigned NOT NULL,
    `name`       TEXT            NOT NULL,
    `price`      INT unsigned    NOT NULL,
    `brand_id`   BIGINT unsigned NOT NULL,
    `created_at` DATETIME(6)     NOT NULL DEFAULT current_timestamp(),
    `updated_at` DATETIME(6)     NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `is_deleted` BOOLEAN         NOT NULL DEFAULT 0,
    `deleted_at` DATETIME(6)     NULL     DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `product`;
-- +goose StatementEnd

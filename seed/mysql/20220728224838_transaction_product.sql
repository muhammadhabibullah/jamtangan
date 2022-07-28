-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS `transaction_product`
(
    `transaction_id` BIGINT unsigned NOT NULL,
    `product_id`     BIGINT unsigned NOT NULL,
    `quantity`       INT unsigned    NOT NULL,
    `price`          INT unsigned    NOT NULL,
    `created_at`     DATETIME(6)     NOT NULL DEFAULT current_timestamp(),
    `updated_at`     DATETIME(6)     NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    `is_deleted`     BOOLEAN         NOT NULL DEFAULT 0,
    `deleted_at`     DATETIME(6)     NULL     DEFAULT NULL,

    PRIMARY KEY (`transaction_id`, `product_id`),
    FOREIGN KEY (`transaction_id`) REFERENCES `transaction` (`id`),
    FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transaction_product`;
-- +goose StatementEnd

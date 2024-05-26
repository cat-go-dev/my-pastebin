-- +goose Up
-- +goose StatementBegin
CREATE TABLE `pastes` (
	`id` VARCHAR(255),
	`pasta` TEXT(20),
	`created_at` TIMESTAMP(20),
	PRIMARY KEY (`id`)
);;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `pastes`';
-- +goose StatementEnd

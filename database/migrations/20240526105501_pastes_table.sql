-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pastes` (
	`hash` VARCHAR(255),
	`pasta` TEXT(20),
	`created_at` TIMESTAMP(20),
	PRIMARY KEY (`hash`)
);;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pastes`;
-- +goose StatementEnd

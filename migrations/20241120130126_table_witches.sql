-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS witches (
  uuid VARCHAR(36) unique, --TODO: добавить во все миграции
  id SERIAL PRIMARY KEY,
  name VARCHAR(64) );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS witches;
-- +goose StatementEnd



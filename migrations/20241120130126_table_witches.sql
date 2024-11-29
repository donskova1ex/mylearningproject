-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS witches (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  name VARCHAR(64) );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS witches;
-- +goose StatementEnd



-- +goose Up
-- +goose StatementBegin
--SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS recipes (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  name VARCHAR(64),
  brew_time_seconds INTEGER);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
-- +goose Up
-- +goose StatementBegin
--SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS ingredients (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY, 
  name VARCHAR(64));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';
DROP TABLE IF EXISTS ingredients;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
--SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS potions (
  uuid VARCHAR(36),
  id SERIAL PRIMARY KEY,
  witch_id SERIAL REFERENCES witches(id),
  recipe_id SERIAL REFERENCES recipes(id),
  status VARCHAR ,
  created_at TIMESTAMP,
  updated_at TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';
DROP TABLE IF EXISTS potions;
-- +goose StatementEnd

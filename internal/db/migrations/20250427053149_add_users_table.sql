-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id    BIGSERIAL     PRIMARY KEY,
  name  VARCHAR(100)  NOT NULL,
  role  VARCHAR(100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

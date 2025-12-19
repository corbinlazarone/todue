-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_google_id;
ALTER TABLE users DROP COLUMN IF EXISTS google_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN google_id VARCHAR(255) UNIQUE NOT NULL DEFAULT '';
CREATE INDEX idx_users_google_id ON users(google_id);
-- +goose StatementEnd

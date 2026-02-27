-- +goose Up
-- +goose StatementBegin
ALTER TABLE courses ADD COLUMN user_id UUID REFERENCES users(id) ON DELETE CASCADE;
CREATE INDEX idx_courses_user_id ON courses(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_courses_user_id;
ALTER TABLE courses DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd

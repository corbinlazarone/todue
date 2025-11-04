-- +goose Up
-- +goose StatementBegin
ALTER TABLE reviews DROP COLUMN image_url;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reviews ADD COLUMN image_url TEXT;
-- +goose StatementEnd

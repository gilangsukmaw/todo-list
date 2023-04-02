-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN password varchar(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd

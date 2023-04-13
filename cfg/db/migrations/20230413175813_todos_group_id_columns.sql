-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
    ADD COLUMN group_id varchar(38) DEFAULT '' NOT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd

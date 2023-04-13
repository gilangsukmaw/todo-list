-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo_groups (
                       id varchar(38) NOT NULL,
                       title varchar(25) NOT NULL,
                       color varchar(7),
                       user_id varchar(255) DEFAULT '' NOT NULL,
                       unique_name varchar(255) DEFAULT '' NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP,
                       PRIMARY KEY (id),
                       UNIQUE (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todo_groups;
-- +goose StatementEnd

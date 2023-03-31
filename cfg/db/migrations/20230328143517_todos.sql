-- +goose Up
-- +goose StatementBegin
CREATE TYPE status AS ENUM ('done', 'on-progress');


CREATE TABLE todos (
                       id varchar(38) NOT NULL,
                       user_id varchar(255) NOT NULL,
                       title TEXT NOT NULL,
                       color varchar(7) NOT NULL,
                       status status DEFAULT 'on-progress' NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP,
                       PRIMARY KEY (id),
                       UNIQUE (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd

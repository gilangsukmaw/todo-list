-- +goose Up
-- +goose StatementBegin
CREATE TYPE roles AS ENUM ('user', 'admin');


CREATE TABLE users (
  id varchar(38) NOT NULL,
  username varchar(20) NOT NULL,
  avatar varchar(255) NOT NULL,
  full_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  role roles DEFAULT 'user' NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

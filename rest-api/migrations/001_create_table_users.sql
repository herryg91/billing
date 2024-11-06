
-- +goose Up
CREATE TABLE users (
    id SERIAL NOT NULL,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    name text NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);

-- +goose Down
DROP TABLE IF EXISTS users;

-- +goose Up

    CREATE TABLE users (
        id UUID PRIMARY KEY,
        created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
        name text NOT NULL UNIQUE,
        display_name text NOT NULL,
        email text NOT NULL UNIQUE,
        password bytea NOT NULL,
        description text NOT NULL DEFAULT 'I am new on Kimb! :3',
        verified boolean NOT NULL DEFAULT false,
        organization boolean NOT NULL DEFAULT false
    );

-- +goose Down
DROP TABLE users;
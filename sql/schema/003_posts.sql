-- +goose Up

    CREATE TABLE posts (
        id UUID PRIMARY KEY,
        created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
        content text NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
    );

-- +goose Down
DROP TABLE posts;
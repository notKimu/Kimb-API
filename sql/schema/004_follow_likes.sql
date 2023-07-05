-- +goose Up

    CREATE TABLE follows (
        id UUID PRIMARY KEY,
        created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        followed_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        UNIQUE(user_id, followed_id)
    );

    CREATE TABLE liked_posts (
        id UUID PRIMARY KEY,
        created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
        UNIQUE(user_id, post_id)
    );

-- +goose Down
DROP TABLE follows;
DROP TABLE liked_posts;
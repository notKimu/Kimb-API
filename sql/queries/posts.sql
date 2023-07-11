-- name: CreatePost :one
INSERT INTO posts (id, content, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts ORDER BY created_at DESC;

-- name: GetPostsFromUser :many
SELECT * FROM posts WHERE user_id=$1 ORDER BY created_at DESC;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2;
-- name: CreatePost :one
INSERT INTO posts (id, content, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2;
-- name: LikePost :one
INSERT INTO liked_posts (id, user_id, post_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLikedPosts :many
SELECT * FROM liked_posts WHERE user_id=$1;

-- name: GetLikesOfPost :many
SELECT * FROM liked_posts WHERE post_id=$1;

-- name: UnlikePost :exec
DELETE FROM liked_posts WHERE post_id=$1 AND user_id=$2;
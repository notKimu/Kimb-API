-- name: FollowUser :one
INSERT INTO follows (id, user_id, followed_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetFollowing :many
SELECT * FROM follows WHERE user_id=$1;

-- name: GetFollowers :many
SELECT * FROM follows WHERE followed_id=$1;

-- name: Unfollow :exec
DELETE FROM follows WHERE user_id=$1 AND followed_id=$2;
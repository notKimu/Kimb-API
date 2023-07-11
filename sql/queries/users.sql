-- name: CreateUser :one
INSERT INTO users (id, name, display_name, email, password, api_key)
VALUES ($1, $2, $3, $4, $5, 
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserFromLogin :one
SELECT * FROM users WHERE name = $1 OR email = $1;

-- name: GetUser :one
SELECT * FROM users WHERE api_key = $1;

-- name: GetUserIDfromName :one
SELECT id FROM users WHERE name = $1;

-- name: GetUserInfo :one
SELECT id, created_at, name, display_name, description, verified, organization FROM users WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1 AND email=$2 AND password=$3;
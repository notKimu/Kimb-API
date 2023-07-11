// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, name, display_name, email, password, api_key)
VALUES ($1, $2, $3, $4, $5, 
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING id, created_at, updated_at, name, display_name, email, password, description, verified, organization, api_key
`

type CreateUserParams struct {
	ID          uuid.UUID
	Name        string
	DisplayName string
	Email       string
	Password    []byte
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.DisplayName,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
		&i.Verified,
		&i.Organization,
		&i.ApiKey,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1 AND email=$2 AND password=$3
`

type DeleteUserParams struct {
	ID       uuid.UUID
	Email    string
	Password []byte
}

func (q *Queries) DeleteUser(ctx context.Context, arg DeleteUserParams) error {
	_, err := q.db.ExecContext(ctx, deleteUser, arg.ID, arg.Email, arg.Password)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name, display_name, email, password, description, verified, organization, api_key FROM users WHERE api_key = $1
`

func (q *Queries) GetUser(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
		&i.Verified,
		&i.Organization,
		&i.ApiKey,
	)
	return i, err
}

const getUserFromLogin = `-- name: GetUserFromLogin :one
SELECT id, created_at, updated_at, name, display_name, email, password, description, verified, organization, api_key FROM users WHERE name = $1 OR email = $1
`

func (q *Queries) GetUserFromLogin(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromLogin, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
		&i.Verified,
		&i.Organization,
		&i.ApiKey,
	)
	return i, err
}

const getUserIDfromName = `-- name: GetUserIDfromName :one
SELECT id FROM users WHERE name = $1
`

func (q *Queries) GetUserIDfromName(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getUserIDfromName, name)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT id, created_at, name, display_name, description, verified, organization FROM users WHERE id = $1
`

type GetUserInfoRow struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	Name         string
	DisplayName  string
	Description  string
	Verified     bool
	Organization bool
}

func (q *Queries) GetUserInfo(ctx context.Context, id uuid.UUID) (GetUserInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, id)
	var i GetUserInfoRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.DisplayName,
		&i.Description,
		&i.Verified,
		&i.Organization,
	)
	return i, err
}

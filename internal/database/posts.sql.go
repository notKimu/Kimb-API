// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: posts.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, content, user_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, content, user_id
`

type CreatePostParams struct {
	ID      uuid.UUID
	Content string
	UserID  uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.ID, arg.Content, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Content,
		&i.UserID,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id=$1 AND user_id=$2
`

type DeletePostParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeletePost(ctx context.Context, arg DeletePostParams) error {
	_, err := q.db.ExecContext(ctx, deletePost, arg.ID, arg.UserID)
	return err
}
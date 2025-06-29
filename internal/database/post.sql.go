// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: post.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addPost = `-- name: AddPost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
)
`

type AddPostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	FeedID      uuid.UUID
}

func (q *Queries) AddPost(ctx context.Context, arg AddPostParams) error {
	_, err := q.db.ExecContext(ctx, addPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	return err
}

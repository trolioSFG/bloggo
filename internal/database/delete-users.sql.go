// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: delete-users.sql

package database

import (
	"context"
)

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM USERS
`

func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteUsers)
	return err
}

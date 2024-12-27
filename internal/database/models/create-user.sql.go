// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: create-user.sql

package models

import (
	"context"
)

const createUse = `-- name: CreateUse :one
SELECT id, name FROM users
`

func (q *Queries) CreateUse(ctx context.Context) (User, error) {
	row := q.db.QueryRow(ctx, createUse)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
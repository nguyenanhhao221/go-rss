// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: feeds_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedsFollow = `-- name: CreateFeedsFollow :one
INSERT INTO feeds_follow ( id, create_at, updated_at, user_id, feed_id) 
VALUES ( $1, $2, $3, $4, $5 )
RETURNING id, create_at, updated_at, user_id, feed_id
`

type CreateFeedsFollowParams struct {
	ID        uuid.UUID
	CreateAt  time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFeedsFollow(ctx context.Context, arg CreateFeedsFollowParams) (FeedsFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedsFollow,
		arg.ID,
		arg.CreateAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedsFollow
	err := row.Scan(
		&i.ID,
		&i.CreateAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}
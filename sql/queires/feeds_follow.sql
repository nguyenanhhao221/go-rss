-- name: CreateFeedsFollow :one
INSERT INTO feeds_follow ( id, create_at, updated_at, user_id, feed_id) 
VALUES ( $1, $2, $3, $4, $5 )
RETURNING *;

-- name: GetFeedToFollows :many
SELECT * FROM feeds_follow WHERE user_id=$1;

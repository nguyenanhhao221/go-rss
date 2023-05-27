-- +goose Up
CREATE TABLE feeds_follow (
    id UUID PRIMARY KEY,
    create_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds_follow;

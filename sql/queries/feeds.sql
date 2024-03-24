-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllFeed :many
SELECT * FROM feeds;

-- name: DeleteFeed :exec
DELETE FROM feeds where id = $1;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
    LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds SET updated_at = $2, last_fetched_at= $3 WHERE id = $1
RETURNING *;
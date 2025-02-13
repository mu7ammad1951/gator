-- name: GetFeeds :many
SELECT * FROM feeds;
--

-- name: AddFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;
--

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE feeds.url = $1;
--
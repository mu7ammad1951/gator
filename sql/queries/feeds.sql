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

-- name: MarkFeedFetched :one
UPDATE feeds 
    SET last_fetched_at = $1, 
        updated_at = $1 
WHERE id = $2 
RETURNING *;
--

-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST, id ASC
LIMIT 1;
--
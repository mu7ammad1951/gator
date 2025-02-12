-- name: GetFeeds: many
SELECT * FROM feeds;

-- name: AddFeed: one
INSERT INTO feeds(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

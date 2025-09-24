-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name 
FROM feeds, users
WHERE feeds.user_id=users.id;  

-- name: CreateFeedFollow :one
WITH inserted_ff AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_ff.*, 
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_ff
INNER JOIN feeds
ON feeds.id = inserted_ff.feed_id
INNER JOIN users
ON users.id = inserted_ff.user_id; 

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN users
ON users.id = feed_follows.user_id
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
WHERE users.name = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows USING users, feeds
WHERE feed_follows.user_id = users.id AND users.name = $2
AND feed_follows.feed_id = feeds.id AND feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET
    last_fetched_at = $1,
    updated_at = $1
WHERE 
    id = $2;

-- name: GetNextFeedToFetch :one
SELECT url FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT 1;

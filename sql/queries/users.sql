-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: Reset :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT name FROM users;

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

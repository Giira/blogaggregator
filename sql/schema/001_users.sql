-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID references users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID references users(id) ON DELETE CASCADE NOT NULL,
    feed_id UUID references feeds(id) ON DELETE CASCADE NOT NULL,
    UNIQUE (user_id, feed_id)
);

ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    feed_id UUID references feeds(id) ON DELETE CASCADE NOT NULL
);

-- +goose Down
DROP TABLE posts
DROP TABLE feed_follows;
DROP TABLE feeds;
DROP TABLE users;

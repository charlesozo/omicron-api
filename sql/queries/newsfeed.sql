-- name: CreateFeed :exec
INSERT INTO Feed (user_feed)
VALUES($1);
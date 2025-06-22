-- name: GetFeedByURL :one
SELECT * FROM FEEDS WHERE URL = $1;


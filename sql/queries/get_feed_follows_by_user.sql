-- name: GetFeedFollowsForUser :many
SELECT ff.*, feeds.name feed_name, users.name user_name
FROM feed_follows ff
INNER JOIN feeds
ON feeds.id = ff.feed_id
INNER JOIN users
ON users.id = ff.user_id
WHERE ff.user_id = $1;


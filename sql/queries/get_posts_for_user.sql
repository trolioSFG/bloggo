-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN feeds
ON posts.feed_id = feeds.id
INNER JOIN feed_follows
ON feeds.id = feed_follows.feed_id
INNER JOIN users
ON feed_follows.user_id = users.id
WHERE users.id = $1
ORDER BY posts.created_at DESC
LIMIT $2;


-- name: MarkFeedFetched :exec
UPDATE feeds
set last_fetched_at = $1, updated_at = $2
where id = $3;


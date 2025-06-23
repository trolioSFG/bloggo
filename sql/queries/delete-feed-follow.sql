-- name: DeleteFeedFollow :exec
DELETE FROM FEED_FOLLOWS
WHERE USER_ID = $1
AND FEED_ID = $2;



-- name: GetUserByID :one
SELECT * FROM USERS WHERE ID = $1;


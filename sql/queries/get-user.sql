-- name: GetUser :one
SELECT * FROM USERS WHERE NAME = $1;


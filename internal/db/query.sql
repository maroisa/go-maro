-- name: GetLink :one
SELECT * FROM links where alias = $1;

-- name: CreateLink :one
INSERT INTO links (source, alias) values ($1, $2) RETURNING *;

-- name: GetBake :one
SELECT * FROM bake
WHERE id = $1 LIMIT 1;

-- name: GetAllBake :many
SELECT * FROM bake;

-- name: CreateBake :one
INSERT INTO bake (
  name, bake_status 
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteBake :exec
DELETE FROM bake
WHERE id = $1;

-- name: UpdateBake :exec
UPDATE bake SET bake_status = $2
WHERE id = $1;
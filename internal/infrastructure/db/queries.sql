-- queries/user.sql

-- name: CreateUser :exec
INSERT INTO users (id, name, email, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users 
SET name = ?, email = ?, updated_at = ? 
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteUser :exec
UPDATE users 
SET deleted_at = ?, updated_at = ? 
WHERE id = ? AND deleted_at IS NULL;

-- name: HardDeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: GetUser :one
SELECT id, name, email, deleted_at, created_at, updated_at 
FROM users 
WHERE id = ? AND deleted_at IS NULL;

-- admin-only
-- name: GetDeletedUser :one
SELECT id, name, email, deleted_at, created_at, updated_at
FROM users
WHERE id = ?;
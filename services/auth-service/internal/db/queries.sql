-- name: CreateUser :one
INSERT INTO users (
  email,
  password_hash,
  first_name,
  last_name,
  full_name,
  role,
  plan
)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login_at = NOW()
WHERE id = $1;

-- name: UpdateUserProfile :one
UPDATE users
SET
  first_name = $2,
  last_name = $3,
  full_name = $4
WHERE id = $1
RETURNING *;

-- name: SoftDeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1;

-- name: CreateSession :one
INSERT INTO user_sessions (
  user_id,
  platform,
  refresh_token,
  access_token_jti,
  expires_at
)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetSessionByRefreshToken :one
SELECT *
FROM user_sessions
WHERE refresh_token = $1
AND is_revoked = FALSE
AND deleted_at IS NULL
LIMIT 1;

-- name: RevokeSession :exec
UPDATE user_sessions
SET is_revoked = TRUE
WHERE refresh_token = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM user_sessions
WHERE expires_at < NOW();

-- name: ListPermissions :many
SELECT *
FROM permissions
ORDER BY name;

-- name: GetPermissionByName :one
SELECT *
FROM permissions
WHERE name = $1
LIMIT 1;

-- name: GetPermissionsByRole :many
SELECT p.id, p.name, p.description
FROM permissions p
JOIN role_permissions rp
ON rp.permission_id = p.id
WHERE rp.role = $1;

-- name: RoleHasPermission :one
SELECT EXISTS (
  SELECT 1
  FROM role_permissions rp
  JOIN permissions p
  ON rp.permission_id = p.id
  WHERE rp.role = $1
  AND p.name = $2
);

-- name: AssignPermissionToRole :exec
INSERT INTO role_permissions (role, permission_id)
VALUES ($1,$2)
ON CONFLICT DO NOTHING;

-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions
WHERE role = $1
AND permission_id = $2;

-- name: GetUserPermissions :many
SELECT p.name
FROM permissions p
JOIN role_permissions rp ON rp.permission_id = p.id
JOIN users u ON u.role = rp.role
WHERE u.id = $1;

-- name: DeleteUserSessions :exec
DELETE FROM user_sessions
WHERE user_id = $1;

-- name: SetUserActiveStatus :exec
UPDATE users
SET is_active = $2
WHERE id = $1;

-- name: EmailExists :one
SELECT EXISTS (
  SELECT 1 FROM users
  WHERE email = $1
  AND deleted_at IS NULL
);

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $2
WHERE id = $1;

-- name: RevokeSessionByID :exec
UPDATE user_sessions
SET is_revoked = TRUE
WHERE id = $1;

-- name: ListUserSessions :many
SELECT *
FROM user_sessions
WHERE user_id = $1
AND is_revoked = FALSE
AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: RevokeAllUserSessions :exec
UPDATE user_sessions
SET is_revoked = TRUE
WHERE user_id = $1;
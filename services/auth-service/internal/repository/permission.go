package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/auth-service/internal/db/sqlc"
)

type PermissionRepository interface {
	ListPermissions(ctx context.Context) ([]db.Permission, error)
	GetPermissionByName(ctx context.Context, name string) (db.Permission, error)
	GetPermissionsByRole(ctx context.Context, role db.RoleType) ([]db.Permission, error)
	RoleHasPermission(ctx context.Context, role db.RoleHasPermissionParams) (bool, error)
	AssignPermissionToRole(ctx context.Context, role db.AssignPermissionToRoleParams) error
	RemovePermissionFromRole(ctx context.Context, permission db.RemovePermissionFromRoleParams) error
	GetUserPermissions(ctx context.Context, id pgtype.UUID) ([]string, error)
}

type permissionRepo struct {
	q *db.Queries
}

// AssignPermissionToRole implements [PermissionRepository].
func (p *permissionRepo) AssignPermissionToRole(ctx context.Context, role db.AssignPermissionToRoleParams) error {
	err := p.q.AssignPermissionToRole(ctx, role)
	if err != nil {
		return fmt.Errorf("assign permission to role: %w", err)
	}
	return nil
}

// GetPermissionByName implements [PermissionRepository].
func (p *permissionRepo) GetPermissionByName(ctx context.Context, name string) (db.Permission, error) {
	perm, err := p.q.GetPermissionByName(ctx, name)
	if err != nil {
		return db.Permission{}, fmt.Errorf("get permission by name: %w", err)
	}
	return perm, nil
}

// GetPermissionsByRole implements [PermissionRepository].
func (p *permissionRepo) GetPermissionsByRole(ctx context.Context, role db.RoleType) ([]db.Permission, error) {
	perms, err := p.q.GetPermissionsByRole(ctx, role)
	if err != nil {
		return nil, fmt.Errorf("get permissions by role: %w", err)
	}
	return perms, nil
}

// GetUserPermissions implements [PermissionRepository].
func (p *permissionRepo) GetUserPermissions(ctx context.Context, id pgtype.UUID) ([]string, error) {
	perms, err := p.q.GetUserPermissions(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user permissions: %w", err)
	}
	return perms, nil
}

// ListPermissions implements [PermissionRepository].
func (p *permissionRepo) ListPermissions(ctx context.Context) ([]db.Permission, error) {
	perms, err := p.q.ListPermissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("ist permissions: %w", err)
	}
	return perms, nil
}

// RemovePermissionFromRole implements [PermissionRepository].
func (p *permissionRepo) RemovePermissionFromRole(ctx context.Context, permission db.RemovePermissionFromRoleParams) error {
	err := p.q.RemovePermissionFromRole(ctx, permission)
	if err != nil {
		return fmt.Errorf("remove permission from role: %w", err)
	}
	return nil
}

// RoleHasPermission implements [PermissionRepository].
func (p *permissionRepo) RoleHasPermission(ctx context.Context, role db.RoleHasPermissionParams) (bool, error) {
	perm, err := p.q.RoleHasPermission(ctx, role)
	if err != nil {
		return false, fmt.Errorf("role has permission: %w", err)
	}
	return perm, nil
}

func NewPerissionRepository(q *db.Queries) PermissionRepository {
	return &permissionRepo{q: q}
}

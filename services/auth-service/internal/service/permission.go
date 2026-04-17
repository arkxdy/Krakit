package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/auth-service/internal/db/sqlc"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/repository"
)

type PermissionService interface {
	ListPermissions(ctx context.Context) ([]dto.PermissionResponse, error)
	AssignPermissionToRole(ctx context.Context, role string, permissionID string) error
	RemovePermissionFromRole(ctx context.Context, role string, permissionID string) error
}

type permService struct {
	permission repository.PermissionRepository
}

func parseRole(role string) (db.RoleType, error) {
	switch role {
	case string(db.RoleTypeSuperAdmin):
		return db.RoleTypeSuperAdmin, nil
	case string(db.RoleTypeAdmin):
		return db.RoleTypeAdmin, nil
	case string(db.RoleTypeExamCreator):
		return db.RoleTypeExamCreator, nil
	case string(db.RoleTypeReviewer):
		return db.RoleTypeReviewer, nil
	case string(db.RoleTypeProctor):
		return db.RoleTypeProctor, nil
	case string(db.RoleTypeSupport):
		return db.RoleTypeSupport, nil
	case string(db.RoleTypeCandidate):
		return db.RoleTypeCandidate, nil
	case string(db.RoleTypeViewer):
		return db.RoleTypeViewer, nil
	default:
		return "", errors.New("invalid role")
	}
}

// AssignPermissionToRole implements [PermissionService].
func (p *permService) AssignPermissionToRole(ctx context.Context, role string, permissionID string) error {
	roleType, err := parseRole(role)
	if err != nil {
		return err
	}

	permUUID, err := uuid.Parse(permissionID)
	if err != nil {
		return errors.New("invalid permission_id")
	}

	return p.permission.AssignPermissionToRole(ctx, db.AssignPermissionToRoleParams{
		Role:         roleType,
		PermissionID: pgtype.UUID{Bytes: permUUID, Valid: true},
	})
}

// ListPermissions implements [PermissionService].
func (p *permService) ListPermissions(ctx context.Context) ([]dto.PermissionResponse, error) {
	perms, err := p.permission.ListPermissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("list permissions: %w", err)
	}

	res := make([]dto.PermissionResponse, 0, len(perms))
	for _, perm := range perms {
		desc := ""
		if perm.Description.Valid {
			desc = perm.Description.String
		}
		res = append(res, dto.PermissionResponse{
			ID:          perm.ID.String(),
			Name:        perm.Name,
			Description: desc,
		})
	}

	return res, nil
}

// RemovePermissionFromRole implements [PermissionService].
func (p *permService) RemovePermissionFromRole(ctx context.Context, role string, permissionID string) error {
	roleType, err := parseRole(role)
	if err != nil {
		return err
	}

	permUUID, err := uuid.Parse(permissionID)
	if err != nil {
		return errors.New("invalid permission_id")
	}

	return p.permission.RemovePermissionFromRole(ctx, db.RemovePermissionFromRoleParams{
		Role:         roleType,
		PermissionID: pgtype.UUID{Bytes: permUUID, Valid: true},
	})
}

func NewPermissionService(r repository.PermissionRepository) PermissionService {
	return &permService{permission: r}
}

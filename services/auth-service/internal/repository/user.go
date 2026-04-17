package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/krakit/auth-service/internal/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error)
	UpdateUserLastLogin(ctx context.Context, id pgtype.UUID) error
	UpdateUserProfile(ctx context.Context, user db.UpdateUserProfileParams) (db.User, error)
	SoftDeleteUser(ctx context.Context, id pgtype.UUID) error
	SetUserActiveStatus(ctx context.Context, status db.SetUserActiveStatusParams) error
	EmailExists(ctx context.Context, email string) (bool, error)
	UpdateUserPassword(ctx context.Context, password db.UpdateUserPasswordParams) error
}

type userRepo struct {
	q *db.Queries
}

// EmailExists implements [UserRepository].
func (u *userRepo) EmailExists(ctx context.Context, email string) (bool, error) {
	exist, err := u.q.EmailExists(ctx, email)
	if err != nil {
		return false, fmt.Errorf("email exists: %w", err)
	}
	return exist, nil
}

// SetUserActiveStatus implements [UserRepository].
func (u *userRepo) SetUserActiveStatus(ctx context.Context, status db.SetUserActiveStatusParams) error {
	err := u.q.SetUserActiveStatus(ctx, status)
	if err != nil {
		return fmt.Errorf("set user active status: %w", err)
	}
	return nil
}

// UpdateUserPassword implements [UserRepository].
func (u *userRepo) UpdateUserPassword(ctx context.Context, password db.UpdateUserPasswordParams) error {
	err := u.q.UpdateUserPassword(ctx, password)
	if err != nil {
		return fmt.Errorf("update user password: %w", err)
	}
	return nil
}

// GetUserByEmail implements [UserRepository].
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := u.q.GetUserByEmail(ctx, email)
	if err != nil {
		return db.User{}, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}

// GetUserByID implements [UserRepository].
func (u *userRepo) GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error) {
	user, err := u.q.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, fmt.Errorf("get user by id: %w", err)
	}
	return user, nil
}

// SoftDeleteUser implements [UserRepository].
func (u *userRepo) SoftDeleteUser(ctx context.Context, id pgtype.UUID) error {
	err := u.q.SoftDeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("soft delete user: %w", err)
	}
	return nil
}

// UpdateUserLastLogin implements [UserRepository].
func (u *userRepo) UpdateUserLastLogin(ctx context.Context, id pgtype.UUID) error {
	err := u.q.UpdateUserLastLogin(ctx, id)
	if err != nil {
		return fmt.Errorf("update user last login: %w", err)
	}
	return nil
}

// UpdateUserProfile implements [UserRepository].
func (u *userRepo) UpdateUserProfile(ctx context.Context, user db.UpdateUserProfileParams) (db.User, error) {
	updatedUser, err := u.q.UpdateUserProfile(ctx, user)
	if err != nil {
		return db.User{}, fmt.Errorf("update user profile: %w", err)
	}
	return updatedUser, nil
}

// CreateUser implements [UserRepository].
func (u *userRepo) CreateUser(ctx context.Context, user db.CreateUserParams) (db.User, error) {
	newUser, err := u.q.CreateUser(ctx, user)
	if err != nil {
		return db.User{}, fmt.Errorf("create user: %w", err)
	}
	return newUser, nil
}

func NewUserRespository(q *db.Queries) UserRepository {
	return &userRepo{q: q}
}

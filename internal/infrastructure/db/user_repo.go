package db

import (
	"context"
	"database/sql"

	"github.com/hogartr/go-hexagonal-template/internal/domain"
)

type userRepo struct {
	q *Queries
}

var _ domain.UserRepository = (*userRepo)(nil)

func NewUserRepo(dbConn *sql.DB) domain.UserRepository {
	return &userRepo{
		q: New(dbConn),
	}
}

func (r *userRepo) Get(ctx context.Context, id domain.UserID) (*domain.User, error) {
	dbUser, err := r.q.GetUser(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return dbUser.ToDomain(), nil
}

func (r *userRepo) Create(ctx context.Context, user *domain.User) error {
	return r.q.CreateUser(ctx, CreateUserParamsFromDomain(user))
}

func (r *userRepo) Update(ctx context.Context, user *domain.User) error {
	return r.q.UpdateUser(ctx, UpdateUserParamsFromDomain(user))
}

func (r *userRepo) SoftDelete(ctx context.Context, id domain.UserID, now domain.Date) error {
	return r.q.SoftDeleteUser(ctx, SoftDeleteUserParams{
		ID:        id.String(),
		DeletedAt: sql.NullString{String: now.String(), Valid: true},
	})
}

// Converts db.User to domain.User
func (u *User) ToDomain() *domain.User {
	id, err := domain.ParseUserId(u.ID)
	if err != nil {
		return nil
	}
	createdAt, err := domain.ParseDate(u.CreatedAt)
	if err != nil {
		return nil
	}
	updatedAt, err := domain.ParseDate(u.UpdatedAt)
	if err != nil {
		return nil
	}
	domainUser := &domain.User{
		ID:        id,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if u.DeletedAt.String != "" {
		deletedAt, err := domain.ParseDate(u.DeletedAt.String)
		if err != nil {
			return nil
		}
		domainUser.DeletedAt = &deletedAt
	}
	return domainUser
}

func CreateUserParamsFromDomain(user *domain.User) CreateUserParams {
	return CreateUserParams{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

func UpdateUserParamsFromDomain(user *domain.User) UpdateUserParams {
	return UpdateUserParams{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt.String(),
	}
}

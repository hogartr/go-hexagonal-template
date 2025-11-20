package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"

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
	dbUser, err := r.q.GetUser(ctx, uuid.UUID(id))
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

func (r *userRepo) SoftDelete(ctx context.Context, id domain.UserID, now time.Time) error {
	return r.q.SoftDeleteUser(ctx, SoftDeleteUserParams{
		ID:        uuid.UUID(id),
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	})
}

// Converts db.User to domain.User
func (u *User) ToDomain() *domain.User {
	domainUser := &domain.User{
		ID:        domain.UserID(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if u.DeletedAt.Valid {
		domainUser.DeletedAt = &u.DeletedAt.Time
	}
	return domainUser
}

func CreateUserParamsFromDomain(user *domain.User) CreateUserParams {
	return CreateUserParams{
		ID:        uuid.UUID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UpdateUserParamsFromDomain(user *domain.User) UpdateUserParams {
	return UpdateUserParams{
		ID:        uuid.UUID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}
}

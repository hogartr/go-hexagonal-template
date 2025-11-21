package domain

import (
	"context"
)

type UserRepository interface {
	Get(ctx context.Context, id UserID) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	SoftDelete(ctx context.Context, id UserID, now Date) error
}

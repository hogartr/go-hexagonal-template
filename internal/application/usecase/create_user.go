package usecase

import (
	"context"
	"fmt"

	"github.com/hogartr/go-hexagonal-template/internal/domain"
	"github.com/hogartr/go-hexagonal-template/internal/infrastructure/clock"
)

type CreateUserCmd struct {
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}

type CreateUserUseCase struct {
	repo  domain.UserRepository
	clock clock.Clock
}

// Constructor (injected by FX)
func NewCreateUserUseCase(repo domain.UserRepository, clock clock.Clock) *CreateUserUseCase {
	return &CreateUserUseCase{
		repo:  repo,
		clock: clock,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, cmd CreateUserCmd) (*domain.User, error) {
	id := domain.NewUserId()
	now := uc.clock.Now()

	user := domain.NewUser(id, cmd.Name, cmd.Email, now)

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

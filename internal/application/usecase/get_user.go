package usecase

import (
	"context"
	"fmt"

	"github.com/hogartr/go-hexagonal-template/internal/domain"
)

type GetUserCmd struct {
	Id string `json:"id" validate:"required,uuid"` // UUID format
}

type GetUserUseCase struct {
	repo domain.UserRepository
}

// Constructor (injected by FX)
func NewGetUserUseCase(repo domain.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		repo: repo,
	}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, cmd GetUserCmd) (*domain.User, error) {
	id, err := domain.ParseUserId(cmd.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	user, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

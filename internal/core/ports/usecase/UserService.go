package usecase

import (
	"context"

	"duon-api/internal/core/domain"
	"duon-api/internal/infra/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.FindByID(ctx, id)
}

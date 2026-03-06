package usecase

import (
	"context"

	"duon-api/internal/core/domain"
	"duon-api/internal/core/ports/repositories"

	"github.com/google/uuid"
)

type FindUserService struct {
	repository repositories.UserRepositoryInterface
}

func NewFindUserService(repository repositories.UserRepositoryInterface) *FindUserService {
	return &FindUserService{repository}
}

func (findUserService *FindUserService) Execute(context context.Context, id uuid.UUID) (*domain.User, error) {
	return findUserService.repository.FindByID(context, id)
}

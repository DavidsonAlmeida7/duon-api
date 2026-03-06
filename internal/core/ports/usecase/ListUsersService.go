package usecase

import (
	"context"

	"duon-api/internal/core/domain"
	"duon-api/internal/core/ports/repositories"
)

type ListUsersService struct {
	repository repositories.UserRepositoryInterface
}

func NewListUsersService(repository repositories.UserRepositoryInterface) *ListUsersService {
	return &ListUsersService{repository}
}

func (listUsersService *ListUsersService) Execute(context context.Context) ([]domain.User, error) {
	return listUsersService.repository.FindAll(context)
}

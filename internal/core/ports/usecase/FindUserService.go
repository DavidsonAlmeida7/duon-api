package usecase

import (
	"context"
	"errors"
	"time"

	"duon-api/internal/core/ports/repositories"
)

type FindUserService struct {
	repository repositories.UserRepositoryInterface
}

type FindUserServiceOutput struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewFindUserService(repository repositories.UserRepositoryInterface) *FindUserService {
	return &FindUserService{repository}
}

func (findUserService *FindUserService) Execute(context context.Context, id string) (*FindUserServiceOutput, error) {
	user, err := findUserService.repository.FindByID(context, id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &FindUserServiceOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

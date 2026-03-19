package userRequestEntity

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserRequest struct {
	Id string `uri:"id" binding:"required"`
}

func NewFindUserRequest(context *gin.Context) (*FindUserRequest, error) {
	findUserRequest := &FindUserRequest{}
	if err := context.ShouldBindUri(findUserRequest); err != nil {
		return nil, err
	}

	parsedId, err := uuid.Parse(findUserRequest.Id)
	if err != nil {
		return nil, err
	}

	findUserRequest.Id = parsedId.String()

	return findUserRequest, nil
}

func (findUserRequest *FindUserRequest) Validate() error {
	parsedId, err := uuid.Parse(findUserRequest.Id)
	if err != nil {
		return errors.New("ID inválido!")
	}

	if parsedId == uuid.Nil {
		return errors.New("ID não pode ser vazio!")
	}

	return nil
}

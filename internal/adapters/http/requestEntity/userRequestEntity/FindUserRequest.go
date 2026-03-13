package userRequestEntity

import (
	"duon-api/internal/core/domain/helper"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserRequest struct {
	Id uuid.UUID `uri:"id" binding:"required"`
}

func NewFindUserRequest(context *gin.Context) (*FindUserRequest, error) {
	findUserRequest := &FindUserRequest{}
	helper.Debug("teste", context.Param("id"))
	if err := context.ShouldBindUri(findUserRequest); err != nil {
		return nil, err
	}

	return findUserRequest, nil
}

func (findUserRequest *FindUserRequest) Validate() error {
	if findUserRequest.Id == uuid.Nil {
		return errors.New("ID inválido!")
	}

	return nil
}

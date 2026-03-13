package handlers

import (
	"duon-api/internal/adapters/http/requestEntity/userRequestEntity"
	"duon-api/internal/core/domain"
	"duon-api/internal/core/ports"
	"duon-api/internal/core/ports/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FindUserHandler struct {
	service *usecase.FindUserService
}

func NewFindUserHandler(service *usecase.FindUserService) ports.HandlerInterface {
	return &FindUserHandler{service}
}

func (findUserHandler *FindUserHandler) Handle(context *gin.Context) {
	//idParam := context.Param("id")

	userRequest, err := findUserHandler.defineUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := findUserHandler.service.Execute(context.Request.Context(), uuid.UUID(userRequest.ID))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado!"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (findUserHandler *FindUserHandler) defineUser(context *gin.Context) (*domain.User, error) {
	//idParam := context.Param("id")
	//
	//id, err := uuid.Parse(idParam)

	userRequest, err := userRequestEntity.NewFindUserRequest(context)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if err := userRequest.Validate(); err != nil {
		return nil, errors.New(err.Error())
	}

	return &domain.User{
		ID: userRequest.Id,
	}, nil
}

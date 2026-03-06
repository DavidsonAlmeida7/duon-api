package handlers

import (
	"duon-api/internal/core/ports"
	"duon-api/internal/core/ports/usecase"
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
	idParam := context.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := findUserHandler.service.Execute(context.Request.Context(), uuid.UUID(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

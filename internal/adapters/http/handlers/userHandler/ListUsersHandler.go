package handlers

import (
	"duon-api/internal/core/ports"
	"duon-api/internal/core/ports/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListUsersHandler struct {
	service *usecase.ListUsersService
}

func NewListUsersHandler(service *usecase.ListUsersService) ports.HandlerInterface {
	return &ListUsersHandler{service}
}

func (userHandler *ListUsersHandler) Handle(context *gin.Context) {
	users, err := userHandler.service.Execute(context.Request.Context())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

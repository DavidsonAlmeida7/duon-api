package routes

import (
	userHandler "duon-api/internal/adapters/http/handlers"
	routesconstants "duon-api/internal/adapters/http/routesConstants"
	userService "duon-api/internal/core/ports/usecase"
	"duon-api/internal/infra/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(group *gin.Engine, db *pgxpool.Pool) {
	userRepo := repository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	group.GET(routesconstants.GetUsersRoutesConst, userHandler.GetUsers)
	group.GET(routesconstants.GetUserByIDRouteConst, userHandler.GetUserByID)
}

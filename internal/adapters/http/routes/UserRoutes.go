package routes

import (
	userHandler "duon-api/internal/adapters/http/handlers/userHandler"
	routesconstants "duon-api/internal/adapters/http/routesConstants"
	userService "duon-api/internal/core/ports/usecase"
	"duon-api/internal/infra/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterUserRoutes(group *gin.Engine, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)

	listUsersService := userService.NewListUsersService(userRepository)
	listUsersHandler := userHandler.NewListUsersHandler(listUsersService)

	findUserService := userService.NewFindUserService(userRepository)
	findUserHandler := userHandler.NewFindUserHandler(findUserService)

	group.GET(routesconstants.GetUsersRoutesConst, listUsersHandler.Handle)
	group.GET(routesconstants.GetUserByIDRouteConst, findUserHandler.Handle)
}

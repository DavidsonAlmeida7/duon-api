package routes

import (
	userHandler "duon-api/internal/adapters/http/handlers"
	"duon-api/internal/adapters/http/routes"
	userService "duon-api/internal/core/ports/usecase"
	"duon-api/internal/infra/repository"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandler.NewUserHandler(userService)

	r.HandleFunc(routes.GetUserRoutesConst, userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	return r
}

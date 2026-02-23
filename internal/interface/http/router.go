package http

import (
	"duon-api/internal/application"
	"duon-api/internal/infra/repository"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	return r
}

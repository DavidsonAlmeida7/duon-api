package http

import (
	"duon-api/internal/adapters/http/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	routes.RegisterUserRoutes(router, db)
	// RegisterClientRoutes(api.Group("/clients"), db)

	return router
}

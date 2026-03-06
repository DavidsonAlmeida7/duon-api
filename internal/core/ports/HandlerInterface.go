package ports

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	Handle(c *gin.Context)
}

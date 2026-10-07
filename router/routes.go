package router

import (
	"chat-go/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	v1 := r.Group("/api/v1")
	{

		v1.POST("/users", userHandler.Create)
	}
}

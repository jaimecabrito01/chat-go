package router

import (
	"chat-go/handler"

	"github.com/gin-gonic/gin"
)

func Init(userHandler *handler.UserHandler) {
	r := gin.Default()
	InitRoutes(r, userHandler)
	r.Run(":8080")

}

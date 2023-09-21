package routes

import (
	contoller "go-jwt-project/contoller"
	"go-jwt-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", contoller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}

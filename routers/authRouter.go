package routers

import (
	"github.com/bostigger/restaurant-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/user/register", controllers.CreateUser)
	incomingRoutes.POST("/api/user/login", controllers.LoginUser)
}

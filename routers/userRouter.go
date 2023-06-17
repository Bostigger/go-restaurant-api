package routers

import (
	"github.com/bostigger/restaurant-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/api/users/get-users", controllers.GetUsers)
	incomingRoutes.GET("/api/users/get-user/:user_id", controllers.GetUserByID)

}

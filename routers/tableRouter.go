package routers

import (
	"github.com/bostigger/restaurant-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/table", controllers.CreateTable)

}

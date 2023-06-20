package routers

import (
	"github.com/bostigger/restaurant-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/food/", controllers.CreateFood)
	incomingRoutes.GET("api/food/:food_id", controllers.GetFood)
	incomingRoutes.GET("api/foods", controllers.GetAllFoods)
	incomingRoutes.PATCH("api/food/:food_id", controllers.UpdateFood)
	incomingRoutes.DELETE("api/food/:food_id", controllers.DeleteFood)
}

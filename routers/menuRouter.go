package routers

import (
	"github.com/bostigger/restaurant-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/menu", controllers.CreateMenu)
	incomingRoutes.PATCH("api/menu/:menu_id", controllers.UpdateMenu)
	incomingRoutes.GET("api/menu/:menu_id", controllers.GetMenu)
	incomingRoutes.GET("api/menus", controllers.GetMenus)
	incomingRoutes.DELETE("api/menu/:menu_id", controllers.DeleteMenu)

}

package main

import (
	"fmt"
	"github.com/bostigger/restaurant-management-api/middlewares"
	"github.com/bostigger/restaurant-management-api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	router.GET("/", func(c *gin.Context) {
		msg := fmt.Sprintf("Api launched successfully")
		c.JSON(http.StatusOK, gin.H{"message": msg})
	})
	err := router.Run(":" + port)
	if err != nil {
		return
	}
	router.Use(middlewares.Authenticate)
	routers.UserRoutes(router)
	routers.FoodRoutes(router)
	routers.InvoiceRoutes(router)
	routers.MenuRoutes(router)
	routers.NoteRoutes(router)
	routers.TableRoutes(router)
	routers.OrderRoutes(router)

}

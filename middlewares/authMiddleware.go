package middlewares

import (
	"github.com/bostigger/restaurant-management-api/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.Set("username", claims.Username)
	c.Set("email", claims.Email)
	c.Set("userId", claims.UserId)
	c.Set("userType", claims.UserType)
	c.Next()
}

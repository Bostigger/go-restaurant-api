package helpers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func checkUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("userType")
	err = nil
	if role != userType {
		err = errors.New("unauthorized")
		return err
	}
	return err
}

func CheckUserAccess(c *gin.Context, userID string) error {
	userType := c.GetString("userType")
	uid := c.GetString("userId")
	fmt.Println(uid)
	fmt.Println(userType)
	if userType == "user" && uid != userID {
		err := errors.New("unauthorized")
		if err != nil {
			return err
		}
	}
	err := checkUserType(c, userType)
	if err != nil {
		return err
	}
	return nil
}

func CheckPrivilege(c *gin.Context) {
	userType := c.GetString("userType")
	if userType != "admin" {
		err := errors.New("unauthorized")
		if err != nil {
			return
		}
	}
	return
}

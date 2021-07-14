package controllers

import (
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func ProvideAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		rollNo, err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid Token: Please try to login again")
			c.Abort()
			return
		}

		c.Set("rollNo", rollNo)
		c.Next()
	}
}

func ProvideAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		rollNo, err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid Token: Please try to login again")
			c.Abort()
			return
		}

		if admin, err := models.IsAdmin(rollNo); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		} else if !admin {
			c.JSON(http.StatusUnauthorized, "You don't have permission to access this route.")
			c.Abort()
			return
		}

		c.Next()
	}
}

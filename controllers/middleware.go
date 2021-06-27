package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProvideAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Invalid Token: Please try to login again")
			c.Abort()
			return
		}
		c.Next()
	}
}

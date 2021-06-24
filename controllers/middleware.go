package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProvideAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Next()
	}
}

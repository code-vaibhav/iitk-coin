package controllers

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	r.POST("/signup", signupHandler)
	r.POST("/login", loginHandler)
}

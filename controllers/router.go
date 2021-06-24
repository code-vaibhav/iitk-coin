package controllers

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	r.POST("/user/signup", signupHandler)
	r.POST("/user/login", loginHandler)
	r.GET("/secretpage", ProvideAuth(), secretHandler)
	r.POST("/coins/reward", rewardCoinsHandler)
	r.POST("/coins/send", transferCoinsHandler)
	r.POST("/coins/balance", balanceCoinsHandler)
}

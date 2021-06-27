package controllers

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	r.POST("/user/signup", signupHandler)
	r.POST("/user/login", loginHandler)
	r.POST("/coins/reward", rewardCoinsHandler)
	r.POST("/coins/send", ProvideAuth(), transferCoinsHandler)
	r.POST("/coins/balance", ProvideAuth(), balanceCoinsHandler)
}

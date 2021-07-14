package controllers

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine) {
	r.POST("/user/signup", signupHandler)
	r.POST("/user/login", loginHandler)

	r.POST("/coins/reward", ProvideAdminAuth(), rewardCoinsHandler)
	r.POST("/coins/send", ProvideAuth(), transferCoinsHandler)
	r.GET("/coins/balance", ProvideAuth(), balanceCoinsHandler)
	r.POST("/coins/redeem", ProvideAuth(), redeemCoinsHandler)

	r.GET("/items", ProvideAuth(), showItemHandler)
	r.POST("/items")

	r.GET("/redeemRequests", ProvideAdminAuth(), showRequestsHandler)
	r.POST("/redeemRequests", ProvideAdminAuth(), changeStatusHandler)
}

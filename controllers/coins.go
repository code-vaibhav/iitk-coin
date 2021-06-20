package controllers

import (
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func rewardCoinsHandler(c *gin.Context) {
	params := models.RewardParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	statusCode, err := rewardCoins(params.RollNo, params.Coins)
	if err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	c.JSON(statusCode, "Transaction completed successfully")
}

func transferCoinsHandler(c *gin.Context) {
	params := models.TransferParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	statusCode, err := tranferCoins(params.Sender, params.Receiver, params.Coins)
	if err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	c.JSON(statusCode, "Transaction completed successfully")
}

func balanceCoinsHandler(c *gin.Context) {
	params := models.BalanceParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := models.FetchUserByRollno(params.RollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
	}
	c.JSON(http.StatusOK, user.Coins)
}

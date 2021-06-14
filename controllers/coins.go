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
	c.JSON(http.StatusOK, params)
}

func transferCoinsHandler(c *gin.Context) {
	params := models.TransferParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, params)
}

func balanceCoinsHandler(c *gin.Context) {
	params := models.BalanceParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, params)
}

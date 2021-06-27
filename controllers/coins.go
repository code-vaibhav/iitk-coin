package controllers

import (
	"errors"
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

	user, err := models.FetchUserByRollno(params.RollNo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if statusCode, err := validate_reciever(*user); err != nil {
		c.JSON(statusCode, err.Error())
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

	sender, errSender := models.FetchUserByRollno(params.Sender)
	reciever, errReciever := models.FetchUserByRollno(params.Receiver)
	if errSender != nil || errReciever != nil {
		c.JSON(http.StatusBadRequest, errors.New("Please provide correct roll numbers").Error())
		return
	}

	if statusCode, err := validate_sender(*sender); err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	if statusCode, err := validate_reciever(*reciever); err != nil {
		c.JSON(statusCode, err.Error())
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

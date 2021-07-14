package controllers

import (
	"errors"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func rewardCoinsHandler(c *gin.Context) {
	params := models.CoinParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := models.FetchUserByRollno(params.RollNo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if statusCode, err := validate_reward_reciever(*user); err != nil {
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
	params := models.CoinParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	senderRollNo := c.MustGet("rollNo").(int)

	sender, errSender := models.FetchUserByRollno(senderRollNo)
	reciever, errReciever := models.FetchUserByRollno(params.RollNo)
	if errSender != nil || errReciever != nil {
		c.JSON(http.StatusBadRequest, errors.New("please provide correct roll numbers").Error())
		return
	}

	if statusCode, err := validate_sender(*sender); err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	if statusCode, err := validate_reciever(*reciever, *sender); err != nil {
		c.JSON(statusCode, err.Error())
		return
	}

	statusCode, err := tranferCoins(senderRollNo, params.RollNo, params.Coins)
	if err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	c.JSON(statusCode, "Transaction completed successfully")
}

func balanceCoinsHandler(c *gin.Context) {
	rollNo := c.MustGet("rollNo").(int)

	user, err := models.FetchUserByRollno(rollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
	}
	c.JSON(http.StatusOK, user.Coins)
}

func redeemCoinsHandler(c *gin.Context) {
	params := models.RedeemParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	rollNo := c.MustGet("rollNo").(int)

	user, err := models.FetchUserByRollno(rollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	item, err := models.FetchItem(params.ItemCode)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	if item.IsAvailable == 0 {
		c.JSON(http.StatusBadRequest, errors.New("item is not availbale at the moment please try woth a different item or try again later"))
	}

	if user.Coins < item.Amount {
		c.JSON(http.StatusBadRequest, "You don't have sufficient coins to redeem items please select another item")
		return
	}

	statusCode, err := redeemCoins(user.RollNo, item)
	if err != nil {
		c.JSON(statusCode, err.Error())
		return
	}
	c.JSON(statusCode, "Your request is pending for admin approval")
}

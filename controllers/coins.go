package controllers

import (
	"database/sql"
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

	//getting database
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	//getting user
	user, err := models.FetchUserByRollno(db, params.RollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	statusCode, err := rewardCoins(db, user, params.Coins)
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

	// getting database
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}

	// getting sender
	sender, err := models.FetchUserByRollno(db, params.Sender)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	// getting receiver
	reciever, err := models.FetchUserByRollno(db, params.Receiver)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	if sender.Coins < params.Coins {
		c.JSON(http.StatusBadRequest, "Insufficient balance")
		return
	}

	statusCode, err := tranferCoins(db, sender, reciever, params.Coins)
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

	//getting database
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}

	user, err := models.FetchUserByRollno(db, params.RollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
	}
	c.JSON(http.StatusOK, user)
}

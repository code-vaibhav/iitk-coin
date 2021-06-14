package controllers

import (
	"database/sql"
	"errors"
	"log"
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

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	user, err := models.FetchUserByRollno(db, params.RollNo)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	// transaction started
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error)
	}
	_, execErr := tx.Exec("UPDATE users SET coins=? WHERE rollNo=?", user.Coins+params.Coins, params.RollNo)
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Cannot commit the transaction"))
		return
	}
	c.JSON(http.StatusOK, "Transaction completed successfully")
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

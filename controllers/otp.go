package controllers

import (
	"net/http"

	"github.com/code-vaibhav/iitk-coin/utils"
	"github.com/gin-gonic/gin"
)

type otpParams struct {
	Otp int `json:"otp"`
}

func makeOtpHandler(c *gin.Context) {
	rollNo := c.MustGet("rollNo").(int)

	err := utils.MakeOtp(rollNo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "otp made successfully")
}

func compareOtpHandler(c *gin.Context) {
	var params otpParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	rollNo := c.MustGet("rollNo").(int)

	err := utils.CompareOtp(rollNo, params.Otp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "otp matched.")
}

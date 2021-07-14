package controllers

import (
	"errors"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func showRequestsHandler(c *gin.Context) {
	requests, err := models.FetchPendingRequests()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(requests) == 0 {
		c.JSON(http.StatusOK, "No requests available with pending status.")
		return
	}
	c.JSON(http.StatusOK, requests)
}

func changeStatusHandler(c *gin.Context) {
	params := models.RedeemRequestParams{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	request, err := models.FetchRequest(params.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("request not found"))
		return
	}

	user, err := models.FetchUserByRollno(request.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("user that made this request not found in database").Error())
		return
	}

	item, err := models.FetchItem(request.ItemCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("item that requested not found").Error())
		return
	}

	statusCode, err := changeStatus(user.RollNo, item.Amount, item.Code, params.Status, params.Id)
	if err != nil {
		c.JSON(statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, "successfully worked on request.")
}

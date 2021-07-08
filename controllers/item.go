package controllers

import (
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func showItemHandler(c *gin.Context) {
	items, err := models.FetchItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusOK, "No items availbale to redeem please comeback later.")
		return
	}

	c.JSON(http.StatusOK, items)
}

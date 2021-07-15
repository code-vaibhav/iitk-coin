package controllers

import (
	"net/http"
	"strconv"

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

func addItemHandler(c *gin.Context) {
	item := models.AddItemParams{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := insertItem(&item); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Item successfully added.")
}

func deleteItemHandler(c *gin.Context) {
	code, err := strconv.Atoi(c.Param("item_code"))
	if err != nil {
		c.JSON(http.StatusBadGateway, "Not a valid route.")
		return
	}

	if _, err := models.FetchItem(code); err != nil {
		c.JSON(http.StatusNotFound, "The item is not found.")
		return
	}

	if err := deleteItem(code); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "item deleted successfully.")
}

func updateItemHandler(c *gin.Context) {
	item := models.Item{}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	code, err := strconv.Atoi(c.Param("item_code"))
	if err != nil {
		c.JSON(http.StatusBadGateway, "Not a valid route.")
		return
	}

	if _, err := models.FetchItem(code); err != nil {
		c.JSON(http.StatusNotFound, "The item is not found.")
		return
	}

	if err := updateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "item updated successfully.")
}

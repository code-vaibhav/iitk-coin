package controllers

import (
	"database/sql"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func signupHandler(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error: Cannot get database")
		return
	}

	if u, _ := models.FetchUserByRollno(db, user.RollNo); u != nil {
		c.JSON(http.StatusAlreadyReported, "User already exist")
		return
	}

	hash, err := models.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	user.Password = hash

	err = insertUser(db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "user successfully added")
}

func loginHandler(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}

	u, _ := models.FetchUserByRollno(db, user.RollNo)
	if u == nil {
		c.JSON(http.StatusNotFound, "Please enter correct rollNo and name")
		return
	}

	if !models.CheckHashPassword(user.Password, u.Password) {
		c.JSON(http.StatusNonAuthoritativeInfo, "Please enter correct password")
		return
	}

	ts, err := createToken(user.RollNo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	tokens := map[string]string{
		"access_token": ts.AccessToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func secretHandler(c *gin.Context) {
	err := TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "User not authenticated")
		return
	}
	c.JSON(http.StatusOK, "User authenticated")
}

package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/gin-gonic/gin"
)

func insertUser(db *sql.DB, user *models.User) error {
	statement, err := db.Prepare("INSERT INTO users(name, rollNo, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Name, user.RollNo, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func displayUsers(db *sql.DB) {
	data := new(models.User)

	rows, err := db.Query("SELECT rollNo , firstname, lastname FROM users")
	CheckErr(err)
	for rows.Next() {
		rows.Scan(&data.RollNo, &data.Name, &data.Password)
		fmt.Println(strconv.Itoa(data.RollNo) + ": " + data.Name)
	}
}

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
		c.JSON(http.StatusUnauthorized, "Please login to acces this page")
		return
	}
	c.JSON(http.StatusOK, "You have login that's why you can see this page")
}

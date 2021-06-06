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

	err := insertUser(db, &user)
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

	err := insertUser(db, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	c.JSON(http.StatusOK, "user successfully added")
}

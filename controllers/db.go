package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/code-vaibhav/iitk-coin/models"
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

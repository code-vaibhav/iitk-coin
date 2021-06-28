package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/code-vaibhav/iitk-coin/sqldb"
	_ "github.com/mattn/go-sqlite3"
)

func insertUser(user *models.User) error {
	statement, err := sqldb.DB.Prepare("INSERT INTO users(name, rollNo, password, isAdmin, isFreezed, coins) VALUES(?, ?, ?, ?, ?, 0)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Name, user.RollNo, user.Password, user.IsAdmin, user.IsFreezed)
	if err != nil {
		return err
	}
	return nil
}

func displayUsers(db *sql.DB) {
	data := new(models.User)

	rows, err := sqldb.DB.Query("SELECT rollNo, name, password, coins, isAdmin, isFreezed FROM users")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&data.RollNo, &data.Name, &data.Password, &data.Coins, &data.IsAdmin, &data.IsFreezed)
		fmt.Println(strconv.Itoa(data.RollNo) + ": " + data.Name)
	}
}

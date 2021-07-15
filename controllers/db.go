package controllers

import (
	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/code-vaibhav/iitk-coin/sqldb"
	_ "github.com/mattn/go-sqlite3"
)

// Item DB operations
func insertItem(item *models.AddItemParams) error {
	stmt, err := sqldb.DB.Prepare("INSERT INTO items(amount, name, isAvailable) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Amount, item.Name, item.IsAvailable)
	if err != nil {
		return err
	}
	return nil
}

func deleteItem(code int) error {
	res, err := sqldb.DB.Exec("DELETE FROM items WHERE code=?", code)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); rows != 1 {
		return err
	}
	return nil
}

func updateItem(item *models.Item) error {
	res, err := sqldb.DB.Exec("UPDATE items SET amount=?, name=?, isAvailable=? WHERE code=?", item.Amount, item.Name, item.IsAvailable, item.Code)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); rows != 1 {
		return err
	}

	return nil
}

// Users DB operations
func insertUser(user *models.User) error {
	statement, err := sqldb.DB.Prepare("INSERT INTO users(name, rollNo, password, isAdmin, isFreezed, coins) VALUES(?, ?, ?, ?, ?, 0)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.RollNo, user.Password, user.IsAdmin, user.IsFreezed)
	if err != nil {
		return err
	}
	return nil
}

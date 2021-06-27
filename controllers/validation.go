package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/code-vaibhav/iitk-coin/sqldb"
)

func validate_reciever(user models.User) (int, error) {
	if user.IsAdmin == 1 {
		return http.StatusBadRequest, errors.New("you cannot transfer money to an admin.")
	}
	return http.StatusOK, nil
}

func validate_sender(user models.User) (int, error) {
	rows, err := sqldb.DB.Query("SELECT COUNT(*) FROM transactions WHERE sender is null AND reciever = ?", user.RollNo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	if count >= 8 {
		return http.StatusBadRequest, errors.New("Sorry you can't send money to others.")
	}
	return http.StatusOK, nil
}

package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
	"github.com/code-vaibhav/iitk-coin/sqldb"
)

func validate_reward_reciever(user models.User) (int, error) {
	if user.IsAdmin == 1 || user.IsFreezed == 1 {
		return http.StatusBadRequest, errors.New("you cannot give reward to an admin or an core team member.")
	}
	return http.StatusOK, nil
}

func validate_reciever(user models.User, sender models.User) (int, error) {
	if user.IsAdmin == 1 {
		return http.StatusBadRequest, errors.New("you cannot transfer money to an admin.")
	}

	if user.IsFreezed == 1 && sender.IsAdmin == 0 {
		return http.StatusBadRequest, errors.New(("you cannot transfer money to an core team member beacause you are not an admin."))
	}

	return http.StatusOK, nil
}

func validate_sender(user models.User) (int, error) {
	rows, err := sqldb.DB.Query("SELECT COUNT(*) FROM transactions WHERE type = ? AND reciever = ?", "Reward", user.RollNo)
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

package controllers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/code-vaibhav/iitk-coin/models"
)

func rewardCoins(db *sql.DB, user *models.User, coins int) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, execErr := tx.Exec("UPDATE users SET coins=? WHERE rollNo=?", user.Coins+coins, user.RollNo)
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, err
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("Cannot commit the transaction")
	}
	return http.StatusOK, nil
}

func tranferCoins(db *sql.DB, sender *models.User, reciever *models.User, coins int) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, execErr := tx.Exec("UPDATE users SET coins=? WHERE rollNo=?", sender.Coins-coins, sender.RollNo)
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, err
	}

	_, execErr = tx.Exec("UPDATE users SET coins=? WHERE rollNo=?", reciever.Coins+coins, reciever.RollNo)
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, err
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("Cannot commit the transaction")
	}
	return http.StatusOK, nil
}

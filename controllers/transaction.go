package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/code-vaibhav/iitk-coin/sqldb"
)

func rewardCoins(rollNo int, coins int) (int, error) {
	tx, err := sqldb.DB.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	res, execErr := tx.Exec("UPDATE users SET coins = coins + ? WHERE rollNo = ? AND coins + ? <= 1000", coins, rollNo, coins)

	if affect, _ := res.RowsAffected(); affect != 1 || execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		if affect == 0 {
			return http.StatusBadRequest, errors.New("Coins limited exceeded for this user")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("Your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	res, execErr = tx.Exec("INSERT INTO transactions (reciever, amount, made_at) VALUES(?, ?, ?)", rollNo, coins, time.Now().Unix())
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}

		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("Cannot commit the transaction")
	}
	return http.StatusOK, nil
}

func tranferCoins(senderRollNo int, recieverRollNo int, coins int) (int, error) {
	tx, err := sqldb.DB.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	res, execErr := tx.Exec("UPDATE users SET coins=coins-? WHERE rollNo=? AND coins-?>=0", coins, senderRollNo, coins)

	if affect, _ := res.RowsAffected(); affect != 1 || execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		if affect == 0 {
			return http.StatusBadRequest, errors.New("You don't have enough coins to transfer")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("Your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	res, execErr = tx.Exec("UPDATE users SET coins=coins+? WHERE rollNo=? AND coins+? <= 1000", coins, recieverRollNo, coins)
	if affect, _ := res.RowsAffected(); affect != 1 || execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		if affect == 0 {
			return http.StatusBadRequest, errors.New("Coins limit exceeded for reciever. You can't do that transaction.")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("Your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	res, execErr = tx.Exec("INSERT INTO transactions (sender, reciever, amount, made_at) VALUES(?, ?, ?, ?)", senderRollNo, recieverRollNo, coins, time.Now().Unix())
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}

		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("Cannot commit the transaction")
	}
	return http.StatusOK, nil
}

package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/code-vaibhav/iitk-coin/models"
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
			return http.StatusBadRequest, errors.New("coins limited exceeded for this user")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	_, execErr = tx.Exec("INSERT INTO transactions (reciever, amount, madeAt, type) VALUES(?, ?, ?, ?)", rollNo, coins, time.Now().Unix(), "Reward")
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("cannot commit the transaction")
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
			return http.StatusBadRequest, errors.New("you don't have enough coins to transfer")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	res, execErr = tx.Exec("UPDATE users SET coins=coins+? WHERE rollNo=? AND coins+? <= 1000", calcAmount(senderRollNo, recieverRollNo, coins), recieverRollNo, calcAmount(senderRollNo, recieverRollNo, coins))
	if affect, _ := res.RowsAffected(); affect != 1 || execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		if affect == 0 {
			return http.StatusBadRequest, errors.New("coins limit exceeded for reciever. You can't do that transaction")
		}
		if affect > 1 {
			return http.StatusBadRequest, errors.New("your request updated more than one entry")
		}
		return http.StatusInternalServerError, execErr
	}

	_, execErr = tx.Exec("INSERT INTO transactions (sender, reciever, amount, madeAt, type) VALUES(?, ?, ?, ?, ?)", senderRollNo, recieverRollNo, coins, time.Now().Unix(), "Transfer")
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("cannot commit the transaction")
	}
	return http.StatusOK, nil
}

func redeemCoins(rollNo int, item *models.Item) (int, error) {
	tx, err := sqldb.DB.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, execErr := tx.Exec("INSERT INTO redeem_requests (user, itemCode, status, madeAt) VALUES(?, ?, ?, ?)", rollNo, item.Code, "Pending", time.Now().Unix())
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, execErr
	}

	_, execErr = tx.Exec("INSERT INTO transactions (reciever, amount, madeAt, type) VALUES(?, ?, ?, ?)", rollNo, item.Amount, time.Now().Unix(), "Redeem")
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("cannot commit the transaction")
	}
	return http.StatusOK, nil
}

func changeStatus(rollNo int, amount int, itemCode int, status string, requestId int) (int, error) {
	tx, err := sqldb.DB.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if status == "Approved" {
		res, execErr := tx.Exec("UPDATE users SET coins = coins - ? WHERE rollNo = ? AND coins-? >= 0", amount, rollNo, amount)
		if execErr != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
			}
			return http.StatusInternalServerError, execErr
		}

		if rows, _ := res.RowsAffected(); rows != 1 {
			return http.StatusInternalServerError, errors.New("user don't have sufficient coins to redeem this item")
		}
	}

	_, execErr := tx.Exec("UPDATE redeem_requests SET status=? WHERE id = ?", status, requestId)
	if execErr != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			log.Fatal("Unable to rollback due to error:", rollBackErr.Error())
		}
		return http.StatusInternalServerError, execErr
	}

	if err = tx.Commit(); err != nil {
		return http.StatusInternalServerError, errors.New("cannot commit the transaction")
	}
	return http.StatusOK, nil
}

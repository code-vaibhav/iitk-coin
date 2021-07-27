package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/code-vaibhav/iitk-coin/sqldb"
)

func MakeOtp(rollNo int) error {
	rows, err := sqldb.DB.Query("SELECT Count(*) FROM otps WHERE user=?", rollNo)
	if err != nil {
		return err
	}

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count != 0 {
		_, err := sqldb.DB.Exec("UPDATE otps SET otp=?, madeAt=? WHERE user=?", rand.Intn(9999-1000)+1000, time.Now().Unix(), rollNo)
		if err != nil {
			return err
		}

		return nil
	}

	_, err = sqldb.DB.Exec("INSERT INTO otps(otp, user, madeAt) values(?, ?, ?)", rand.Intn(9999-1000)+1000, rollNo, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}

func CompareOtp(rollNo int, enteredOtp int) error {
	rows, err := sqldb.DB.Query("SELECT otp, madeAt FROM otps WHERE user=?", rollNo)
	if err != nil {
		return err
	}

	var otp int
	var createdAt int
	var count int = 0
	for rows.Next() {
		rows.Scan(&otp, &createdAt)
		count++
	}

	if count == 0 {
		return errors.New("otp not exist: Request otp and then try with new otp")
	}

	if time.Now().Unix()-int64(createdAt) > 300 {
		return errors.New("otp has expired please request otp again")
	}

	if otp == enteredOtp {
		return nil
	}
	return errors.New("incorrect otp")
}

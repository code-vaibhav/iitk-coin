package models

import (
	"errors"

	"github.com/code-vaibhav/iitk-coin/sqldb"
)

type User struct {
	Name      string `json:"name"`
	RollNo    int    `json:"rollNo"`
	Password  string `json:"password"`
	IsAdmin   int    `json:"isAdmin"`
	IsFreezed int    `json:"isFreezed"`
	Coins     int
}

type LoginParams struct {
	RollNo   int    `json:"rollNo"`
	Password string `json:"password"`
}

func fetch(query string, args ...interface{}) ([]*User, error) {
	rows, err := sqldb.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*User, 0)

	for rows.Next() {
		data := new(User)
		err := rows.Scan(&data.RollNo, &data.Name, &data.Password, &data.Coins, &data.IsAdmin, &data.IsFreezed)
		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func FetchAllUsers() ([]*User, error) {
	query := ("SELECT * FROM users")

	return fetch(query)
}

func FetchUserByRollno(rollNo int) (*User, error) {
	query := "SELECT * FROM users WHERE rollNo=?"

	rows, err := fetch(query, rollNo)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, errors.New("User not found in database")
	}
}

func IsAdmin(rollNo int) (bool, error) {
	user, err := FetchUserByRollno(rollNo)
	if err != nil {
		return false, err
	}

	return user.IsAdmin != 0, nil
}

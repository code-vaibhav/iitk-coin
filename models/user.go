package models

import (
	"database/sql"
	"errors"
)

type User struct {
	Name     string `json:"name"`
	RollNo   int    `json:"rollNo"`
	Password string `json:"password"`
}

func fetch(db *sql.DB, query string, args ...interface{}) ([]*User, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*User, 0)

	for rows.Next() {
		data := new(User)
		err := rows.Scan(&data.RollNo, &data.Name, &data.Password)
		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func FetchAllUsers(db *sql.DB) ([]*User, error) {
	query := ("SELECT rollNo, name, password FROM users")

	return fetch(db, query)
}

func FetchUserByRollno(db *sql.DB, rollNo int) (*User, error) {
	query := "SELECT rollNo, name, password FROM users WHERE rollNo=?"

	rows, err := fetch(db, query, rollNo)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, errors.New("User not found in database")
	}
}

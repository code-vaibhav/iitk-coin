package models

import (
	"errors"

	"github.com/code-vaibhav/iitk-coin/sqldb"
)

type Request struct {
	Id       int
	User     int
	ItemCode int
	Status   string
	MadeAt   string
}

func fetchRequests(query string, args ...interface{}) ([]*Request, error) {
	rows, err := sqldb.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*Request, 0)

	for rows.Next() {
		data := new(Request)
		err := rows.Scan(&data.Id, &data.User, &data.ItemCode, &data.Status, &data.MadeAt)
		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func FetchRequest(id int) (*Request, error) {
	query := `SELECT * FROM redeem_requests WHERE id = ?`

	rows, err := fetchRequests(query, id)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, errors.New("request not found in database")
	}
}

func FetchPendingRequests() ([]*Request, error) {
	query := `SELECT * FROM redeem_requests WHERE status = 'Pending'`

	return fetchRequests(query)
}

func FetchRequests() ([]*Request, error) {
	query := `SELECT * FROM redeem_requests`

	return fetchRequests(query)
}

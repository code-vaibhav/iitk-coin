package models

import (
	"errors"

	"github.com/code-vaibhav/iitk-coin/sqldb"
)

type Item struct {
	Code        int
	Amount      int
	Name        string
	IsAvailable int
}

func fetchItems(query string, args ...interface{}) ([]*Item, error) {
	rows, err := sqldb.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*Item, 0)

	for rows.Next() {
		data := new(Item)
		err := rows.Scan(&data.Code, &data.Amount, &data.Name, &data.IsAvailable)
		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func FetchItem(code int) (*Item, error) {
	query := `SELECT * FROM items WHERE code = ?`

	rows, err := fetchItems(query, code)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, errors.New("item not found in database")
	}
}

func FetchItems() ([]*Item, error) {
	query := `SELECT * FROM items`

	return fetchItems(query)
}

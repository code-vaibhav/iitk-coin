package models

type CoinParams struct {
	RollNo int `json:"rollNo"`
	Coins  int `json:"coins"`
}

type RedeemParams struct {
	ItemCode int `json:"itemCode"`
}

type RedeemRequestParams struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type AddItemParams struct {
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
	IsAvailable int    `json:"isAvailable"`
}

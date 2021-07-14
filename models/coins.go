package models

type CoinParams struct {
	RollNo int `json:"rollNo"`
	Coins  int `json:"coins"`
}

type RedeemParams struct {
	ItemCode int `json:"itemCode"`
}

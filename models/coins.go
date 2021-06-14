package models

type RewardParams struct {
	RollNo int `json:"rollNo"`
	Coins  int `json:"coins"`
}

type TransferParams struct {
	Sender   int `json:"sender_rollNo"`
	Receiver int `json:"receiver_rollNo"`
	Coins    int `json:"coins"`
}

type BalanceParams struct {
	RollNo int `json:"rollNo"`
}

package controllers

import (
	"strconv"
)

func calcAmount(senderRollNo int, recieverRollNo int, coins int) int {
	if strconv.Itoa(senderRollNo)[0:1] == strconv.Itoa(recieverRollNo)[0:1] {
		return coins - ((coins * 2) / 100)
	}

	return coins - ((coins * 33) / 100)
}

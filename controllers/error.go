package controllers

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

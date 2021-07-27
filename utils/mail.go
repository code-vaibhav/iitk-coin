package utils

import (
	"strconv"

	gomail "gopkg.in/mail.v2"
)

func SendMail(rollNo int, otp int) error {
	// password := os.Getenv("EMAIL_PASSWORD")

	m := gomail.NewMessage()
	m.SetHeader("From", "vgoyal20@iitk.ac.in")
	m.SetHeader("To", "vaibhavgoyal2506@gmail.com")
	m.SetHeader("Subject", "OTP")
	m.SetBody("text/plain", strconv.Itoa(otp)+" is your account verification otp.")

	d := gomail.NewDialer("mmtp.iitk.ac.in", 25, "vgoyal20@iitk.ac.in", "Vaibhav@2506")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

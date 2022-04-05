package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "safety.system@hitachi-kenki.com")
	m.SetHeader("To", "h.may.ps@hitachi-kenki.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	d := gomail.Dialer{Host: "10.87.42.15", Port: 25}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

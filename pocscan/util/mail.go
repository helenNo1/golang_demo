package util

import "gopkg.in/gomail.v2"

const (
	mail_pass   string = "THUIBWTUTIICVAUO"
	mail_addr   string = "fckjp1102@163.com"
	mail_server string = "smtp.163.com"
	mail_port   int    = 465
)

func SendMail(dst_mail_addr, title, msg string) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail_addr)
	m.SetHeader("To", mail_addr)
	m.SetAddressHeader("Cc", mail_addr, "Dan")
	m.SetHeader("Subject", title)
	m.SetBody("text/html", msg)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(mail_server, mail_port, mail_addr, mail_pass)

	// Send the email to Bob, Cora and Dan.
	d.DialAndSend(m)
}

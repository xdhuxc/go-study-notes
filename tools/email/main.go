package main

import (
	"gopkg.in/gomail.v2"

	log "github.com/sirupsen/logrus"
)

func main() {
	/*
		username := "noreply@ushareit.com"
		message := "你好，邮件"
		auth := smtp.PlainAuth("", username, "v3naPFgei9fQ6Her", "smtp.exmail.qq.com")
		to := []string{"wanghuan@ushareit.com", "xdhuxc@163.com"}
		// cc := []string{"wanghuan@ushareit.com", "xdhuxc@163.com"}

		err := smtp.SendMail("smtp.exmail.qq.com:465", auth, username, to, []byte(message))

		if err != nil {
			log.Println()
		}

	*/
	sendEmail()

}

func sendEmail() {

	m := gomail.NewMessage()
	message := "你好，邮件"

	m.SetHeader("From", "noreply@ushareit.com")
	m.SetHeader("To", "wanghuan@ushareit.com,xdhuxc@163.com")
	m.SetHeader("Subject", "你好")
	m.SetHeader("Cc", "xdhuxc@163.com, abc@163.com")
	m.SetBody("text/html", message)

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "noreply@ushareit.com", "v3naPFgei9fQ6Her")
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}

}

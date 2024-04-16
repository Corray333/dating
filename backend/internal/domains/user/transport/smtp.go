package transport

import (
	"log"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
)

func SendEmail(to, text string) {
	from := viper.GetString("smtp_email")
	pass := os.Getenv("SMTP_PASSWORD")

	msg := []byte("To: " + to + "\r\n" +
		"Subject: Why aren’t you using Mailtrap yet?\r\n" +
		"\r\n" +
		text + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}

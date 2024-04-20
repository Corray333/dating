package transport

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
)

func formatHeader(key, value string) string {
	return key + ": " + value + "\r\n"
}

// TODO: replace with universal SendEmail and HtmlBody constructor
func SendVerificationCode(to, code string, username string) error {
	from := viper.GetString("smtp_email")
	pass := os.Getenv("SMTP_PASS")

	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = "Verification code?"
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	// Convert the headers to a format suitable for sending
	header := ""
	for k, v := range headers {
		header += formatHeader(k, v)
	}

	tmpl, err := template.ParseFiles("../templates/email_verification.html")
	if err != nil {
		panic(err)
	}

	// Define the data to be inserted into the template
	data := struct {
		Username string
		Code     string
	}{
		Username: username,
		Code:     code,
	}

	var htmlBody bytes.Buffer
	// Execute the template with the data
	err = tmpl.Execute(&htmlBody, data)
	if err != nil {
		panic(err)
	}

	// Combine the headers and the HTML body
	msg := []byte(header + "\r\n" + htmlBody.String())

	if err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, msg); err != nil {
		return err
	}
	return nil
}

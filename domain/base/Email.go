package base

import (
	"SGMS/domain/config"

	"gopkg.in/gomail.v2"
)

const (
	EMAIL_SENDER_SUPPORT = 0
)

type SendEmail struct {
	To, Subject, Html string
}

type IEmailSender interface {
	Send()
}

func NewEmailSender(email *EmailSender) IEmailSender {
	return email
}

type EmailSender struct {
	From              int
	To, Subject, Html string
}

// 发送邮箱验证码
func (this *EmailSender) Send() {
	m := gomail.NewMessage()

	m.SetHeader("To", this.To)
	m.SetHeader("Subject", this.Subject)
	m.SetBody("text/html", this.Html)
	//	m.Attach("/home/Alex/lolcat.jpg")
	var d *gomail.Dialer
	if EMAIL_SENDER_SUPPORT == this.From {
		m.SetHeader("From", config.EmailSupportUsername)
		d = gomail.NewDialer(config.EmailSupportSmtp, config.EmailSupportPort, config.EmailSupportUsername, config.EmailSupportPassword)
	}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

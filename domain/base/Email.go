package base

const (
	EMAIL_SENDER_SUPPORT = 0
)

type SendEmail struct {
	To, Subject, Html string
}

type IEmailSender interface {
	Send()
}

// func NewEmailSender(email *EmailSender) IEmailSender {
// 	return email
// }

type EmailSender struct {
	From              int
	To, Subject, Html string
}

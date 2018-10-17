package contracts

type MailContract interface {
	To(email string) MailContract
	From(email string) MailContract
	Subject(subject string) MailContract
	Send(message string)
}
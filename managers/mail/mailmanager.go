package mail

import (
	"os"
	"github.com/flame/drivers"
	"github.com/flame/contracts"
)



func Compose() contracts.MailContract {
	maildriver := os.Getenv("MAIL_DRIVER")
	if maildriver == "smtp" {
		return drivers.MailSmtpDriver{}
	} else {

	}

	return drivers.MailSmtpDriver{}
}
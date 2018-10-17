package mail

import (
	"os"

	"github.com/mitchdennett/flameframework/contracts"
	"github.com/mitchdennett/flameframework/drivers"
)

func Compose() contracts.MailContract {
	maildriver := os.Getenv("MAIL_DRIVER")
	if maildriver == "smtp" {
		return drivers.MailSmtpDriver{}
	} else {

	}

	return drivers.MailSmtpDriver{}
}

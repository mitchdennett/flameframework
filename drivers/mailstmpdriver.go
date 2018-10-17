package drivers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/mitchdennett/flameframework/contracts"
)

type MailSmtpDriver struct {
	to      string
	from    string
	message string
	subject string
}

func (m MailSmtpDriver) To(email string) contracts.MailContract {
	m.to = email
	return m
}

func (m MailSmtpDriver) Subject(subject string) contracts.MailContract {
	m.subject = subject
	return m
}

func (m MailSmtpDriver) From(email string) contracts.MailContract {
	m.from = email
	return m
}

func (m MailSmtpDriver) Send(message string) {
	m.message = message
	fmt.Println("sending smtp")
	fmt.Println(m.to)
	fmt.Println(m.message)

	if m.from == "" {
		m.from = os.Getenv("MAIL_FROM")
	}

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		host,
	)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", host+":"+port, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(m.from); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(m.to); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	msg := []byte("To: " + m.to + "\r\n" +
		"Subject: " + m.subject + "\r\n" +
		"\r\n" + m.message +
		"\r\n")

	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}

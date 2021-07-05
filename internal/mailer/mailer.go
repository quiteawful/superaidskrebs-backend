package mailer

import (
	"errors"
	"fmt"

	gomail "gopkg.in/mail.v2"

	"github.com/quiteawful/superaidskrebs-backend/pkg/config"
)

var (
	params config.MailConf
)

func SetParams(mailparams config.MailConf) {
	params = mailparams
}

func sendMail(address, subject, message string) error {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", params.From)

	// Set E-Mail receivers
	m.SetHeader("To", address)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", message)

	// Settings for SMTP server
	d := gomail.NewDialer(params.SMTPHost, params.SMTPPort, params.From, params.Password)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return errors.New("Error sending Mail: " + err.Error())
	}
	return nil
}

func SendCode(address string, username string, code string) error {
	message :=
		fmt.Sprintf("Hi %v, \n\n here is your activation code: \n%v\n Use /user activate $Code to activate your user account. \n\n Yours Superaidskrebs Bot",
			username, code)
	err := sendMail(address, "Superaidskrebs User Activation", message)
	if err != nil {
		return errors.New("Error sending Code: " + err.Error())
	}
	return nil
}

func SendPassword(address string, username string, password string) error {
	message :=
		fmt.Sprintf("Hi %v, \n\nhere is your new password: \n%v\n Change it as fast as you can as EMail is not save for transporting passwords. \n\nYours Superaidskrebs Bot",
			username, password)
	err := sendMail(address, "Superaidskrebs Password reset", message)
	if err != nil {
		return errors.New("Error sending password: " + err.Error())
	}
	return nil
}

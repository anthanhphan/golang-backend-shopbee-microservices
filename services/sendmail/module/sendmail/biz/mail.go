package mailbiz

import (
	"errors"
	"log"
	"net/smtp"
	"os"
)

const (
	AcceptUpgrade = "Congratulations!!! Your registration to become retailer" +
		" was reviewed and accepted by our staff." +
		" Now you can start selling your products in our platform.\r\n"

	DeniedUpgrade = "We are sorry to announce that you are not eligible to become our retailer!\r\n"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unkown from server")
		}
	}
	return nil, nil
}

func SendMail(toEmail string, subjectMail string, msgMail string) {
	mail := os.Getenv("MAIL_ADDR")
	pass := os.Getenv("MAIL_PASS")

	auth := LoginAuth(mail, pass)

	to := []string{toEmail}
	msg := []byte(
		"To: " +
			toEmail +
			"\r\n" +
			"Subject: " +
			subjectMail +
			"\r\n" +
			msgMail +
			"\r\n",
	)

	err := smtp.SendMail("smtp.gmail.com:587", auth, mail, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

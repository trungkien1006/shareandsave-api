package emailrepo

import (
	"fmt"
	"net/smtp"
)

type SMTPEmailSender struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewSMTPEmailSender(host string, port int, username, password, from string) *SMTPEmailSender {
	return &SMTPEmailSender{host, port, username, password, from}
}

func (s *SMTPEmailSender) Send(to, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n"+
			"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n%s",
		s.From, to, subject, body,
	))

	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}

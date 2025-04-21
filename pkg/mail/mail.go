package mail

import (
	"medtest/config"
	"net/smtp"
)

var message []byte = []byte("Your IP changed")

type Notify struct {
	From     string
	PWD      string
	SMTPhost string
	SMTPport string
	Type     string
}

func NewNotify(Config *config.Config) *Notify {
	return &Notify{From: Config.Mailer.From, PWD: Config.Mailer.PWD, SMTPhost: Config.Mailer.SMTPhost, SMTPport: Config.Mailer.SMTPport, Type: Config.Mailer.BuildType}
}

func (Notify *Notify) NewMail(To string) error {
	if Notify.Type == "Testing" {
		return nil
	}
	var toArr []string
	toArr = append(toArr, To)
	auth := smtp.PlainAuth("", Notify.From, Notify.PWD, Notify.SMTPhost)
	err := smtp.SendMail(Notify.SMTPhost+":"+Notify.SMTPport, auth, Notify.From, toArr, message)
	return err
}

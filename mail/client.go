package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"stock_reminder/conf"
)

var mailClient *gomail.Dialer

func init() {
	mailClient = gomail.NewDialer(conf.Conf.MailConfig.Host, conf.Conf.MailConfig.Port, conf.Conf.MailConfig.Account, conf.Conf.MailConfig.Password)
	mailClient.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func sendEmail() error {
	err := mailClient.DialAndSend(&gomail.Message{

	})
	return err
}

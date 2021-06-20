package mail

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/sync/errgroup"
	"gopkg.in/gomail.v2"
	"stock_reminder/conf"
	"stock_reminder/model"
	"strconv"
)

var mailChan chan model.StockInfo

var mailClient *gomail.Dialer

func init() {
	mailChan = make(chan model.StockInfo, 16)
	mailClient = gomail.NewDialer(conf.Conf.MailConfig.Host, conf.Conf.MailConfig.Port, conf.Conf.MailConfig.Account, conf.Conf.MailConfig.Password)
	mailClient.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func SendMailSignal(info model.StockInfo) {
	mailChan <- info
}

func SendMail() {
	for {
		select {
		case stockInfo := <-mailChan:
			err := sendEmail(stockInfo)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func sendEmail(info model.StockInfo) error {
	var eg errgroup.Group
	for _, receiver := range conf.Conf.MailConfig.Receivers {
		receiverUserInfo := receiver
		eg.Go(func() error {
			msg := gomail.NewMessage()
			msg.SetHeader("From", conf.Conf.MailConfig.Account)
			msg.SetHeader("To", receiverUserInfo)
			msg.SetHeader("Subject", info.StockCode+info.StockName+"现价:"+strconv.FormatFloat(info.Current, 'f', 2, 64))
			err := mailClient.DialAndSend(msg)
			if err != nil {
				return err
			}
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {
		return err
	}
	return nil
}

func CloseMailChan() {
	close(mailChan)
}

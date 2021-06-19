package mail

import (
	"fmt"
	"stock_reminder/model"
)

var mailChan chan model.StockInfo

func init() {
	mailChan = make(chan model.StockInfo, 16)
}

func SendMailSignal(info model.StockInfo) {
	mailChan <- info
}

func SendMail() {
	for {
		select {
		case <-mailChan:
			err := sendEmail()
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func CloseMailChan() {
	close(mailChan)
}

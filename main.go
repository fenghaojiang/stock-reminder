package main

import (
	"fmt"
	"stock_reminder/action"
	"stock_reminder/mail"

	"github.com/robfig/cron/v3"
)

func main() {
	options := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	cronJob := cron.New(cron.WithParser(options), cron.WithChain())
	_, err := cronJob.AddFunc("0 0/1 9-15 * * *", action.GetStockInfo) //from 9:00AM to 15:00PM
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go mail.HandleSendMail()
	defer mail.CloseMailChan()
	cronJob.Start()
	select {}
}

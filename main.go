package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"stock_reminder/action"
	"stock_reminder/mail"
)

func main() {
	options := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	cronJob := cron.New(cron.WithParser(options), cron.WithChain())
	_, err := cronJob.AddFunc("0 0/1 * * * *", action.GetStockInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go mail.SendMail()
	cronJob.Start()
}

package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Version     string      `toml:"version"`
	StockConfig StockConfig `toml:"stock"`
	MailConfig  MailConfig  `toml:"smtp_mail"`
}

type StockConfig struct {
	XueQiuUrl    string          `toml:"xueqiuDest"`
	EastMoneyUrl string          `toml:"eastmoneyDest"`
	StockList    [][]interface{} `toml:"stockList"`
}

type MailConfig struct {
	Host      string   `toml:"host"`
	Port      int      `toml:"port"`
	Account   string   `toml:"account"`
	Password  string   `toml:"password"`
	Receivers []string `toml:"receivers"`
}

var Conf Config

func init() {
	err := parse()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// /var/data/conf/config.toml

func parse() error {
	if _, err := toml.DecodeFile("./conf/config.toml", &Conf); err != nil {
		return err
	}
	return nil
}

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
	XueQiuUrl    string  `toml:"xueqiuDest"`
	EastMoneyUrl string  `toml:"eastmoneyDest"`
	StockList    []Stock `toml:"stockList"`
}

type Stock struct {
	StockCode string
	Price     float64
}

type MailConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Account  string `toml:"account"`
	Password string `toml:"password"`
}

var Conf Config

func init() {
	err := parse()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func parse() error {
	if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
		return err
	}
	return nil
}

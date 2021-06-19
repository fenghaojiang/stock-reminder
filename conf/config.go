package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Version     string      `toml:"version"`
	StockConfig StockConfig `toml:"stock"`
}

type StockConfig struct {
	EastMoneyUrl string  `toml:"eastmoneyDest"`
	StockList    []int64 `toml:"stockList"`
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

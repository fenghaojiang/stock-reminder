package action

import (
	"fmt"
	"github.com/tidwall/gjson"
	"stock_reminder/conf"
	"testing"
)

func TestSendRequest(t *testing.T) {
	conf.Conf.StockConfig.XueQiuUrl = "https://stock.xueqiu.com/v5/stock/realtime/pankou.json"
	stockInfo, err := sendRequest("SH600183")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
	fmt.Println(gjson.Get(stockInfo, "data.symbol").String())
	fmt.Println(gjson.Get(stockInfo, "data.current").Float())
}

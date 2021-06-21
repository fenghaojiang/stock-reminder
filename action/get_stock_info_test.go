package action

import (
	"fmt"
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
	fmt.Println(stockInfo)
}

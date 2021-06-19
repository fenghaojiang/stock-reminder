package action

import (
	"fmt"
	_ "stock_reminder/conf"
	"testing"
)

func TestOnGetStockInfo(t *testing.T) {
	err := sendRequest("SH600183")
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}
}

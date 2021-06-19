package action

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"stock_reminder/conf"
	"strconv"
	"time"
)

var client = &http.Client{}

func NewGetValues(stockNumStr string) *url.Values {
	getValues := url.Values{}
	getValues.Add("fltt", "2")
	getValues.Add("invt", "2")
	getValues.Add("secid", "1."+stockNumStr)
	getValues.Add("fields", "f43")
	getValues.Add("ut", "b2884a393a59ad64002292a3e90d46a5")
	return &getValues
}

func GetStockInfo() {
	nowUnixNano := time.Now().UnixNano()
	unixStr := strconv.FormatInt(nowUnixNano, 10)
	req, err := http.NewRequest("GET", conf.Conf.StockConfig.EastMoneyUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

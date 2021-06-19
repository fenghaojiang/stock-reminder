package action

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"net/url"
	"stock_reminder/conf"
	"strings"
)

var client = &http.Client{}

func NewGetValues(stockCode string) *url.Values {
	getValues := url.Values{}
	getValues.Add("symbol", stockCode)
	return &getValues
}

func GetStockInfo() {
	ch := make(chan struct{})
	var eg errgroup.Group
	for _, stockCode := range conf.Conf.StockConfig.StockList {
		ch <- struct{}{}
		code := stockCode
		eg.Go(func() error {
			defer func() {
				<-ch
			}()
			err := sendRequest(code)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			return nil
		})
	}
}

func sendRequest(stockCode string) error {
	args := NewGetValues(stockCode)
	req, err := http.NewRequest("GET", conf.Conf.StockConfig.XueQiuUrl, strings.NewReader(args.Encode()))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

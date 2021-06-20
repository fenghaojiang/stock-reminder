package action

import (
	"fmt"
	"github.com/tidwall/gjson"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"net/url"
	"stock_reminder/conf"
	"stock_reminder/mail"
	"stock_reminder/model"
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
	for _, stock := range conf.Conf.StockConfig.StockList {
		ch <- struct{}{}
		code := stock[0].(string)
		name := stock[1].(string)
		price := stock[2].(float64)
		eg.Go(func() error {
			defer func() {
				<-ch
			}()
			stockInfo, err := sendRequest(code)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			sc := gjson.Get(stockInfo, "data.symbol")
			current := gjson.Get(stockInfo, "data.current")
			curFloat := current.Float()
			if curFloat <= price {
				mail.SendMailSignal(model.StockInfo{
					StockCode: sc.String(),
					StockName: name,
					Current:   curFloat,
				})
			}
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {

	}
}

func sendRequest(stockCode string) (string, error) {
	args := NewGetValues(stockCode)
	req, err := http.NewRequest("GET", conf.Conf.StockConfig.XueQiuUrl, strings.NewReader(args.Encode()))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363")
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

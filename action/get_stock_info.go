package action

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"stock_reminder/conf"
	"stock_reminder/mail"
	"stock_reminder/model"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/sync/errgroup"
)

var client = &http.Client{}

func NewGetValues(stockCode string) *url.Values {
	getValues := url.Values{}
	getValues.Add("symbol", stockCode)
	return &getValues
}

func GetStockInfo() {
	ch := make(chan struct{}, 16)
	var eg errgroup.Group
	for _, stock := range conf.Conf.StockConfig.StockList {
		code := stock[0].(string)
		name := stock[1].(string)
		priceStr := stock[2].(string)
		ch <- struct{}{}
		eg.Go(func() error {
			defer func() {
				<-ch
			}()
			stockInfo, err := sendRequest(code)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			nowUnix := time.Now().String()
			sc := gjson.Get(stockInfo, "data.symbol")
			current := gjson.Get(stockInfo, "data.current")
			curFloat := current.Float()
			price, err := strconv.ParseFloat(priceStr, 64)
			fmt.Println("timeNow:", nowUnix, "\t"+"stockCode: ", sc, "\t"+"current", current)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if curFloat <= price && curFloat != float64(0.0) { //查不到
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
		fmt.Println(err.Error())
		return
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
		fmt.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	// fmt.Println(string(body)) //stock info
	return string(body), nil
}

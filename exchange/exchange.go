package exchange

import (
	"errors"
	"fmt"
	"github.com/Toscale-platform/toscale-kit/http"
	"strings"
)

type BinanceExchangeInfo struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Symbols []struct {
		BaseAsset  string `json:"baseAsset"`
		QuoteAsset string `json:"quoteAsset"`
	} `json:"symbols"`
}

type BitfinexConfig [][]string

type KucoinSymbols struct {
	Data []struct {
		BaseCurrency  string `json:"baseCurrency"`
		QuoteCurrency string `json:"quoteCurrency"`
	} `json:"data"`
}

type PoloniexSymbols []struct {
	DisplayName string `json:"displayName"`
}

func GetSymbols(exchange string) (symbols []string, err error) {
	symbols = make([]string, 0)

	switch exchange {
	case "binance":
		rawSymbols := BinanceExchangeInfo{}
		err = http.Get("https://api.binance.com/api/v3/exchangeInfo", &rawSymbols)
		if err != nil {
			return
		}

		if rawSymbols.Code != 0 || rawSymbols.Message != "" {
			err = fmt.Errorf("code: %d, msg: %s", rawSymbols.Code, rawSymbols.Message)
			return
		}

		for _, symbol := range rawSymbols.Symbols {
			symbols = append(symbols, symbol.BaseAsset+"/"+symbol.QuoteAsset)
		}
	case "bitfinex":
		rawSymbols := BitfinexConfig{}
		err = http.Get("https://api-pub.bitfinex.com/v2/conf/pub:list:pair:exchange", &rawSymbols)
		if err != nil {
			return
		}

		for _, symbol := range rawSymbols[0] {
			if strings.Index(symbol, ":") > -1 {
				s := strings.Split(symbol, ":")
				symbols = append(symbols, s[0]+"/"+s[1])
			} else {
				symbols = append(symbols, symbol[:3]+"/"+symbol[3:])
			}
		}
	case "kucoin":
		rawSymbols := KucoinSymbols{}
		err = http.Get("https://api.kucoin.com/api/v2/symbols", &rawSymbols)
		if err != nil {
			return
		}

		for _, symbol := range rawSymbols.Data {
			symbols = append(symbols, symbol.BaseCurrency+"/"+symbol.QuoteCurrency)
		}
	case "poloniex":
		rawSymbols := PoloniexSymbols{}
		err = http.Get("https://api.poloniex.com/markets", &rawSymbols)
		if err != nil {
			return
		}

		for _, symbol := range rawSymbols {
			symbols = append(symbols, symbol.DisplayName)
		}
	default:
		return symbols, errors.New("exchange not found")
	}

	return symbols, nil
}

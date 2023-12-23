package validator

import (
	"slices"
	"strings"
)

var exchanges = []string{"binance", "bitfinex", "kucoin", "poloniex"}

func InitExchangeList(e []string) {
	exchanges = e
}

func Is(s string, slice []string) bool {
	for _, el := range slice {
		if el == s {
			return true
		}
	}

	return false
}

func IsExchange(exchange string) bool {
	return IsExchangeWith(exchange, exchanges)
}

func IsExchangeWith(exchange string, exchanges []string) bool {
	exchange = strings.TrimSpace(exchange)

	if exchange == "" {
		return false
	}

	if !slices.Contains(exchanges, exchange) {
		return false
	}

	return true
}

func IsSymbol(symbol string) bool {
	return IsSymbolWith(symbol, "/")
}

func IsSymbolWith(symbol, separator string) bool {
	symbol = strings.TrimSpace(symbol)

	if symbol == "" {
		return false
	}

	if !strings.Contains(symbol, separator) {
		return false
	}

	return true
}

func IsBuyerMakerStr(side string) bool {
	return side == "buy"
}

func IsBuyerMakerOC(open, close float64) bool {
	return !(open > close)
}

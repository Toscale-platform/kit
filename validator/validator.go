package validator

import (
	"slices"
)

var Exchanges = []string{"binance", "bitfinex", "kucoin", "poloniex"}

func IsExchange(exchange string) bool {
	return IsExchangeWith(exchange, Exchanges)
}

func IsExchangeWith(exchange string, exchanges []string) bool {
	return slices.Contains(exchanges, exchange)
}

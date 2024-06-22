package validator

import (
	v10 "github.com/go-playground/validator/v10"
	"slices"
	"strings"
)

var (
	v         = v10.New(v10.WithRequiredStructEnabled())
	exchanges = []string{"binance", "bitfinex", "kucoin", "poloniex"}
)

func InitExchangeList(e []string) {
	exchanges = e
}

// Is call go-playground/validator
// https://pkg.go.dev/github.com/go-playground/validator/v10
func Is(field interface{}, tag string) bool {
	return v.Var(field, tag) == nil
}

func InSlice(s string, slice []string) bool {
	return slices.Contains(slice, s)
}

func IsExchange(exchange string) bool {
	return IsExchangeWith(exchange, exchanges)
}

func IsExchangeWith(exchange string, exchanges []string) bool {
	return slices.Contains(exchanges, exchange)
}

func IsSymbol(symbol string) bool {
	return IsSymbolWith(symbol, "/")
}

func IsSymbolWith(symbol, sep string) bool {
	return strings.Contains(symbol, sep)
}

func IsBuyerMaker(side string) bool {
	return side == "buy"
}

func IsBuyerMakerOC(open, close float64) bool {
	return !(open > close)
}

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

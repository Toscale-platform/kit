package exchange

import "testing"

func TestGetSymbols(t *testing.T) {
	exchanges := []string{"binance", "bitfinex", "kucoin", "poloniex", "noname"}

	for _, exchange := range exchanges {
		symbols, err := GetSymbols(exchange)
		if err != nil && err.Error() != "exchange not found" {
			t.Error(err)
			return
		}

		if len(symbols) == 0 && exchange != "noname" {
			t.Error("number of symbols must be greater than 0")
			return
		}
	}
}

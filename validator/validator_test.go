package validator

import (
	"github.com/Toscale-platform/kit/tests"
	"testing"
)

func TestIs(t *testing.T) {
	tests.True(t, Is("email@example.com", "required,email"))
}

func TestInSlice(t *testing.T) {
	tests.True(t, InSlice("ABC", []string{"ABC", "DEF"}))
	tests.False(t, InSlice("NOO", []string{"ABC", "DEF"}))
}

func TestIsExchange(t *testing.T) {
	tests.True(t, IsExchange("binance"))
}

func TestIsExchangeWith(t *testing.T) {
	tests.True(t, IsExchangeWith("myexchange", []string{"myexchange"}))
}

func TestIsExchangeWithInit(t *testing.T) {
	InitExchangeList([]string{"ownexchange"})
	tests.True(t, IsExchange("ownexchange"))
}

func TestIsSymbol(t *testing.T) {
	tests.True(t, IsSymbol("BTC/USDT"))
}

func TestIsSymbolWith(t *testing.T) {
	tests.True(t, IsSymbolWith("BTC-USDT", "-"))
}

func TestIsBuyerMakerStr(t *testing.T) {
	tests.True(t, IsBuyerMaker("buy"))
}

func TestIsBuyerMakerOC(t *testing.T) {
	tests.True(t, IsBuyerMakerOC(2.0, 3.0))
}

func TestIsEmpty(t *testing.T) {
	tests.True(t, IsEmpty("       "))
}

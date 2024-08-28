package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIs(t *testing.T) {
	assert.True(t, Is("email@example.com", "required,email"))
}

func TestInSlice(t *testing.T) {
	assert.True(t, InSlice("ABC", []string{"ABC", "DEF"}))
	assert.False(t, InSlice("NOO", []string{"ABC", "DEF"}))
}

func TestIsExchange(t *testing.T) {
	assert.True(t, IsExchange("binance"))
}

func TestIsExchangeWith(t *testing.T) {
	assert.True(t, IsExchangeWith("myexchange", []string{"myexchange"}))
}

func TestIsExchangeWithInit(t *testing.T) {
	InitExchangeList([]string{"ownexchange"})
	assert.True(t, IsExchange("ownexchange"))
}

func TestIsSymbol(t *testing.T) {
	assert.True(t, IsSymbol("BTC/USDT"))
}

func TestIsSymbolWith(t *testing.T) {
	assert.True(t, IsSymbolWith("BTC-USDT", "-"))
}

func TestIsBuyerMakerStr(t *testing.T) {
	assert.True(t, IsBuyerMaker("buy"))
}

func TestIsBuyerMakerOC(t *testing.T) {
	assert.True(t, IsBuyerMakerOC(2.0, 3.0))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty("       "))
}

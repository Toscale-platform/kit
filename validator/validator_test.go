package validator

import "testing"

func TestIs(t *testing.T) {
	if !Is("email@example.com", "required,email") {
		t.Error("Expected true, got false")
	}
}

func TestInSlice(t *testing.T) {
	if !InSlice("ABC", []string{"ABC", "DEF"}) {
		t.Error("Expected true, got false")
	}

	if InSlice("NOO", []string{"ABC", "DEF"}) {
		t.Error("Expected false, got true")
	}
}

func TestIsExchange(t *testing.T) {
	if !IsExchange("binance") {
		t.Error("Expected true, got false")
	}
}

func TestIsExchangeWith(t *testing.T) {
	if !IsExchangeWith("myexchange", []string{"myexchange"}) {
		t.Error("Expected true, got false")
	}
}

func TestIsExchangeWithInit(t *testing.T) {
	InitExchangeList([]string{"ownexchange"})
	if !IsExchange("ownexchange") {
		t.Error("Expected true, got false")
	}
}

func TestIsSymbol(t *testing.T) {
	if !IsSymbol("BTC/USDT") {
		t.Error("Expected true, got false")
	}
}

func TestIsSymbolWith(t *testing.T) {
	if !IsSymbolWith("BTC-USDT", "-") {
		t.Error("Expected true, got false")
	}
}

func TestIsBuyerMakerStr(t *testing.T) {
	if !IsBuyerMaker("buy") {
		t.Error("Expected true, got false")
	}
}

func TestIsBuyerMakerOC(t *testing.T) {
	if !IsBuyerMakerOC(2.0, 3.0) {
		t.Error("Expected true, got false")
	}
}

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("       ") {
		t.Error("Expected true, got false")
	}
}

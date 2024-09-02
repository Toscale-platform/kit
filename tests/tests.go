package tests

import "testing"

func Err(t *testing.T, err error) {
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func HasErr(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error")
	}
}

func Nil(t *testing.T, v interface{}) {
	if v == nil {
		t.Error("expected not nil")
	}
}

func Equal[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func True(t *testing.T, v bool) {
	if !v {
		t.Error("expected true")
	}
}

func False(t *testing.T, v bool) {
	if v {
		t.Error("expected false")
	}
}

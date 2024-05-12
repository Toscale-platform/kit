package env

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func test[X comparable](t *testing.T, wantType string, got X, want X) {
	gotType := reflect.TypeOf(got).String()
	if gotType != wantType {
		t.Errorf("got type %s, want %s", gotType, wantType)
		return
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func testSlice[X ~[]E, E cmp.Ordered](t *testing.T, wantType string, got X, want X) {
	gotType := reflect.TypeOf(got).String()
	if gotType != wantType {
		t.Errorf("got type %s, want %s", gotType, wantType)
		return
	}

	if slices.Compare(got, want) != 0 {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// String

func TestGetString(t *testing.T) {
	test(t, "string", GetString("STR"), "text")
}

// Slices of string, bytes and runes

func TestGetSlice(t *testing.T) {
	testSlice(t, "[]string", GetSlice("SLICE"), []string{"a", "b", "c"})
}

func TestGetBytes(t *testing.T) {
	testSlice(t, "[]uint8", GetBytes("SLICE"), []byte{97, 44, 98, 44, 99})
}

func TestGetRunes(t *testing.T) {
	testSlice(t, "[]int32", GetRunes("SLICE"), []rune{'a', ',', 'b', ',', 'c'})
}

// Bool

func TestGetBool(t *testing.T) {
	test(t, "bool", GetBool("BOOL"), true)
}

// Int

func TestGetInt(t *testing.T) {
	test(t, "int", GetInt("INT"), 100)
}

func TestGetInt8(t *testing.T) {
	test(t, "int8", GetInt8("INT"), 100)
}

func TestGetInt16(t *testing.T) {
	test(t, "int16", GetInt16("INT"), 100)
}

func TestGetInt32(t *testing.T) {
	test(t, "int32", GetInt32("INT"), 100)
}

func TestGetInt64(t *testing.T) {
	test(t, "int64", GetInt64("INT"), 100)
}

// Uint

func TestGetUint(t *testing.T) {
	test(t, "uint", GetUint("INT"), 100)
}

func TestGetUint8(t *testing.T) {
	test(t, "uint8", GetUint8("INT"), 100)
}

func TestGetUint16(t *testing.T) {
	test(t, "uint16", GetUint16("INT"), 100)
}

func TestGetUint32(t *testing.T) {
	test(t, "uint32", GetUint32("INT"), 100)
}

func TestGetUint64(t *testing.T) {
	test(t, "uint64", GetUint64("INT"), 100)
}

// Float

func TestGetFloat32(t *testing.T) {
	test(t, "float32", GetFloat32("INT"), 100)
}

func TestGetFloat64(t *testing.T) {
	test(t, "float64", GetFloat64("INT"), 100)
}

// Complex

func TestGetComplex64(t *testing.T) {
	test(t, "complex64", GetComplex64("INT"), 100)
}

func TestGetComplex128(t *testing.T) {
	test(t, "complex128", GetComplex128("INT"), 100)
}

// Empty

func TestEmptyString(t *testing.T) {
	test(t, "string", GetString("NOT EXIST KEY"), "")
}

func TestEmptySlice(t *testing.T) {
	testSlice(t, "[]string", GetSlice("NOT EXIST KEY"), nil)
}

func TestEmptyBool(t *testing.T) {
	test(t, "bool", GetBool("NOT EXIST KEY"), false)
}

func TestEmptyInt(t *testing.T) {
	test(t, "int", GetInt("NOT EXIST KEY"), 0)
}

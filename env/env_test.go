package env

import (
	"cmp"
	"github.com/Toscale-platform/kit/tests"
	"reflect"
	"slices"
	"testing"
)

func test[X comparable](t *testing.T, wantType string, got X, want X) {
	tests.Equal(t, wantType, reflect.TypeOf(got).String())
	tests.Equal(t, want, got)
}

func testSlice[X ~[]E, E cmp.Ordered](t *testing.T, wantType string, got X, want X) {
	tests.Equal(t, wantType, reflect.TypeOf(got).String())
	tests.Equal(t, slices.Compare(want, got), 0)
}

// String

func TestGetString(t *testing.T) {
	test(t, "string", GetString("STR"), "text")
}

// Slices of string, bytes and runes

func TestGetSlice(t *testing.T) {
	testSlice(t, "[]string", GetSlice("SLICE"), []string{"a", "b", "c"})
}

// Bool

func TestGetBool(t *testing.T) {
	test(t, "bool", GetBool("BOOL"), true)
}

// Int

func TestGetInt(t *testing.T) {
	test(t, "int", GetInt("INT"), 100)
}

func TestGetInt64(t *testing.T) {
	test(t, "int64", GetInt64("INT"), 100)
}

// Uint

func TestGetUint(t *testing.T) {
	test(t, "uint", GetUint("INT"), 100)
}

func TestGetUint64(t *testing.T) {
	test(t, "uint64", GetUint64("INT"), 100)
}

// Float

func TestGetFloat64(t *testing.T) {
	test(t, "float64", GetFloat64("INT"), 100)
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

// Default

func TestDefaultString(t *testing.T) {
	test(t, "string", GetString("DEF", "default"), "default")
}

func TestDefaultSlice(t *testing.T) {
	testSlice(t, "[]string", GetSlice("DEF", []string{"a", "b"}), []string{"a", "b"})
}

func TestDefaultBool(t *testing.T) {
	test(t, "bool", GetBool("DEF", true), true)
}

func TestDefaultInt(t *testing.T) {
	test(t, "int", GetInt("DEF", 100), 100)
}

package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

func init() {
	_ = godotenv.Load()
}

func parseInt(key string, bitSize int) int64 {
	v, err := strconv.ParseInt(os.Getenv(key), 10, bitSize)
	if err != nil {
		return 0
	}

	return v
}

func parseUint(key string, bitSize int) uint64 {
	v, err := strconv.ParseUint(os.Getenv(key), 10, bitSize)
	if err != nil {
		return 0
	}

	return v
}

func parseFloat(key string, bitSize int) float64 {
	v, err := strconv.ParseFloat(os.Getenv(key), bitSize)
	if err != nil {
		return 0
	}

	return v
}

func parseComplex(key string, bitSize int) complex128 {
	v, err := strconv.ParseComplex(os.Getenv(key), bitSize)
	if err != nil {
		return 0
	}

	return v
}

// String

func GetString(key string) string {
	return os.Getenv(key)
}

// Slices of string, bytes and runes

func GetSlice(key string) []string {
	v := os.Getenv(key)
	if v == "" {
		return []string{}
	}

	return strings.Split(v, ",")
}

func GetBytes(key string) []byte {
	return []byte(os.Getenv(key))
}

func GetRunes(key string) []rune {
	return []rune(os.Getenv(key))
}

// Bool

func GetBool(key string) bool {
	v, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return false
	}

	return v
}

// Int

func GetInt(key string) int {
	return int(parseInt(key, 0))
}

func GetInt8(key string) int8 {
	return int8(parseInt(key, 8))
}

func GetInt16(key string) int16 {
	return int16(parseInt(key, 16))
}

func GetInt32(key string) int32 {
	return int32(parseInt(key, 32))
}

func GetInt64(key string) int64 {
	return parseInt(key, 64)
}

// Uint

func GetUint(key string) uint {
	return uint(parseUint(key, 0))
}

func GetUint8(key string) uint8 {
	return uint8(parseUint(key, 8))
}

func GetUint16(key string) uint16 {
	return uint16(parseUint(key, 16))
}

func GetUint32(key string) uint32 {
	return uint32(parseUint(key, 32))
}

func GetUint64(key string) uint64 {
	return parseUint(key, 64)
}

// Float

func GetFloat32(key string) float32 {
	return float32(parseFloat(key, 32))
}

func GetFloat64(key string) float64 {
	return parseFloat(key, 64)
}

// Complex

func GetComplex64(key string) complex64 {
	return complex64(parseComplex(key, 64))
}

func GetComplex128(key string) complex128 {
	return parseComplex(key, 128)
}

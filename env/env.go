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

// String

func GetString(key string, defaultValue ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Slice of string

func GetSlice(key string, defaultValue ...[]string) []string {
	v := os.Getenv(key)
	if len(v) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return []string{}
	}
	return strings.Split(v, ",")
}

// Bool

func GetBool(key string, defaultValue ...bool) bool {
	v, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return false
	}
	return v
}

// Int

func GetInt(key string, defaultValue ...int) int {
	v, err := strconv.ParseInt(os.Getenv(key), 10, 0)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return int(v)
}

func GetInt64(key string, defaultValue ...int64) int64 {
	v, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return v
}

// Uint

func GetUint(key string, defaultValue ...uint) uint {
	v, err := strconv.ParseUint(os.Getenv(key), 10, 0)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return uint(v)
}

func GetUint64(key string, defaultValue ...uint64) uint64 {
	v, err := strconv.ParseUint(os.Getenv(key), 10, 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return v
}

// Float

func GetFloat64(key string, defaultValue ...float64) float64 {
	v, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}
	return v
}

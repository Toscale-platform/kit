package env

import (
	"errors"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			panic(err)
		}
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetSlice(key string) []string {
	value := viper.GetString(key)
	if value == "" {
		return []string{}
	}

	return strings.Split(value, ",")
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

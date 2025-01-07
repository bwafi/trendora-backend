package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./")       // path working directory
	config.AddConfigPath("./../")    // path for test
	config.AddConfigPath("./../../") // path for subfolder test

	config.AutomaticEnv()

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}

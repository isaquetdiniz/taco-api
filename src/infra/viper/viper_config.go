package config

import "github.com/spf13/viper"

viper.SetDefault("port", 3000)
viper.SetConfigName("env")
viper.AddConfigPath("taco-api")

err := viper.ReadInConfig()
if err != nil {
	panic(fmt.Errorf("fatal error config file: %w", err))
}

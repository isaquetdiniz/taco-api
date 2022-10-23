package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configs struct {
	MODE string
	PORT int
	DSN  string
}

func Init() Configs {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}

	viper.SetDefault("MODE", "LOCAL")
	viper.SetDefault("PORT", 3000)

	mode := viper.GetString("MODE")
	port := viper.GetInt("PORT")
	databaseUrl := viper.GetString("DATABASE_URL")

	return Configs{
		MODE: mode,
		PORT: port,
		DSN:  databaseUrl,
	}
}

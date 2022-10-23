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
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		_, errIsFileNotFound := err.(viper.ConfigFileNotFoundError)

		if errIsFileNotFound {
			fmt.Println("Env file not found")
		} else {
			panic(fmt.Errorf("Fatal error config file: %w", err))
		}
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

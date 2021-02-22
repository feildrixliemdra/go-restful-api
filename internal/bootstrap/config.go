package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatal("failed to read config file, ", err)
		}
	}
	viper.SetEnvPrefix(`app`)
	viper.AutomaticEnv()
}

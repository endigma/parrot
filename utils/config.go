package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath(".")                    // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/.parrot")        // call multiple times to add many search paths
	viper.AddConfigPath("$HOME/.config/parrot") // call multiple times to add many search paths

	viper.SetDefault("format.output", "decimal")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
		} else {
			log.Warn("Config file not provided, defaults will be used.")
		}
	} else {
		viper.WatchConfig()
	}

	viper.WriteConfig()
}

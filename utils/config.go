package utils

import (
	"os"

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
			// Write default config

			log.Warn("Config file not provided, writing default values to config.toml in your config directory")
			_, err := os.Stat(UserHomeDir() + "/parrot/")

			if os.IsNotExist(err) {
				errDir := os.MkdirAll(UserHomeDir()+"/parrot/", 0755)
				if errDir != nil {
					log.Fatal(err)
				}

			}

			_, err = os.Create(UserHomeDir() + "/parrot/config.toml")
			CheckErr(err)
			viper.WriteConfig()
		}
	} else {
		viper.WatchConfig()
	}

	viper.WriteConfig()
}

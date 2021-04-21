package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("parrot") // name of config file (without extension)

	viper.AddConfigPath(".")                    // optionally look for config in the working directory
	viper.AddConfigPath("$HOME/.parrot")        // call multiple times to add many search paths
	viper.AddConfigPath("$HOME/.config/parrot") // call multiple times to add many search paths

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WriteConfig()
}

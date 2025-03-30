// config file
package config

import (
	"github.com/spf13/viper"
	"fmt"
)

// cpath: config path
func check_config(cpath string) bool {
	
	viper.AddConfigPath(cpath)
	viper.SetConfigName("config")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			fmt.Errorf("Config file not found.")
			//...
			return false
		} else {
			// Config file was found but another error was produced
			fmt.Errorf("Error reading config file: %v", err)
			//...
			return false
		}
	}
	// Config file found and successfully parsed
	fmt.Println("Config file found and successfully parsed.")
	return true
}

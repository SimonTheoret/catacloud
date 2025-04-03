// config file
package config

import (
	"os"
	"github.com/spf13/viper"
	"fmt"
	"path/filepath"
)

// Wrapper around Viper
type Config struct {
	v *viper.Viper
	path string
}

// Initializes Config instance
func NewConfig(cpath string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(cpath)
	v.SetConfigName("config")

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found in path: %s", cpath)
		} else {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	fmt.Println("Config file found and successfully parsed.")
	return &Config{v: v, path: cpath}, nil
}



// Basics fonctionnalities inspired by Viper

// Get returns the value associated with the given key
func (c *Config) Get(key string) interface{} {
	return c.v.Get(key)
}
// Set sets or overwrite the value for the given key
func (c *Config) Set(key string, value interface{}) {
	c.v.Set(key, value)
}
// GetPath returns the path where the configuration file
func (c *Config) GetPath() string {
	return c.path
}

// Save writes the current configuration to the file it was read from
func (c *Config) Save() error {
	// Ensure the directory exists
	if err := os.MkdirAll(c.path, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write the configuration to the file
	if err := c.v.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	fmt.Println("Config file saved successfully.")
	return nil
}

// SaveAs writes the current configuration to a specified file with a specified format
func (c *Config) SaveAs(filePath, fileFormat string) error {
	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Write the configuration to the specified file
	if err := c.v.WriteConfigAs(filePath); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	// Update the path
	c.path = filePath

	fmt.Printf("Config file saved successfully as %s.\n", filePath)
	return nil
}

package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestCheckConfig(t *testing.T) {
	// Create a temporary directory for the config file
	tempDir := t.TempDir()

	// Test case 1: Config file exists
	configContent := []byte("key: value")
	configPath := tempDir + "/config.yaml"
	err := os.WriteFile(configPath, configContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	viper.Reset() // Reset Viper to clear any previous configuration
	result := check_config(tempDir)
	if !result {
		t.Errorf("Expected true for existing config file, got false")
	}
	if viper.Get("key") != "value" {
		t.Errorf("Expected config value 'value', got %v", viper.Get("key"))
	}

	// Test case 2: Config file does not exist
	viper.Reset()
	result = check_config("/nonexistent/path")
	if result {
		t.Errorf("Expected false for non-existent config file, got true")
	}
}


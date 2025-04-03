package config

import (
	"os"
	"testing"
	"path/filepath"
)

func TestNewConfig(t *testing.T) {
	// Create a temporary directory for the config file
	tempDir := t.TempDir()


	// Test case 1: Config file exists
	configContent := []byte(`
		key = "value"
		path = "/path/to/config"
	`)
	configPath := tempDir + "/config.toml"
	err := os.WriteFile(configPath, configContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	cfg, err := NewConfig(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Check Get 
	if cfg.Get("key") != "value" {
		t.Errorf("Expected config value 'value', got %v", cfg.Get("key"))
	}
	
	// Check Set
	cfg.Set("new_key", "new_value")
	if cfg.Get("new_key") != "new_value" {
		t.Errorf("Expected config value 'new_value', got %v", cfg.Get("new_key"))
	}
	// Check if the path is correctly returned
	if cfg.GetPath() != tempDir {
		t.Errorf("Expected path '%s', got %v", tempDir, cfg.GetPath())
	}

	// Test case 2: Config file does not exist
	_, err = NewConfig("/nonexistent/path")
	if err == nil {
		t.Errorf("Expected error for non-existent config file, got nil")
	}
}



func TestSave(t *testing.T) {
	// Create a temporary directory for the config file
	tempDir := t.TempDir()

	// Create a config file in the temporary directory
	configContent := []byte(`
		key = "value"
	`)
	configPath := filepath.Join(tempDir, "config.toml")
	err := os.WriteFile(configPath, configContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Initialize the config
	cfg, err := NewConfig(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	} 

	// Modify the configuration
	cfg.Set("key", "new_value")

	// Save the configuration
	if err := cfg.Save(); err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Read the saved configuration file
	savedContent, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read saved config file: %v", err)
	}

	// Check if the configuration was saved correctly
	expectedContent := "key = 'new_value'\n"
	if string(savedContent) != expectedContent {
		t.Errorf("Expected config content:\n%s\nGot:\n%s", expectedContent, string(savedContent))
	}
}

func TestSaveAs(t *testing.T) {
	// Create a temporary directory for the config file
	tempDir := t.TempDir()

	// Create a config file in the temporary directory
	configContent := []byte(`
		key = "value"
	`)
	configPath := filepath.Join(tempDir, "config.toml")
	err := os.WriteFile(configPath, configContent, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Initialize the config
	cfg, err := NewConfig(tempDir)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Modify the configuration
	cfg.Set("key", "new_value")

	// Save the configuration to a new file
	newConfigPath := filepath.Join(tempDir, "config_temp_test.json") 
	if err := cfg.SaveAs(newConfigPath, "json"); err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Read the saved configuration file
	savedContent, err := os.ReadFile(newConfigPath)
	if err != nil {
		t.Fatalf("Failed to read saved config file: %v", err)
	}

	// Check if the configuration was saved correctly
	expectedContent := "{\n  \"key\": \"new_value\"\n}"
	if string(savedContent) != expectedContent {
		t.Errorf("Expected config content:\n%s\nGot:\n%s", expectedContent, string(savedContent))
	}

	// Check if the Config as the new filepath
	if cfg.GetPath() != newConfigPath {
		t.Errorf("Expected config path \n%s\nGot:\n%s", newConfigPath, cfg.GetPath())
	}
}


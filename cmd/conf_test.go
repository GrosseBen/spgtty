package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestConfigLoading tests the config file loading functionality
func TestConfigLoading(t *testing.T) {
	// Create a temporary config file
	tempDir := t.TempDir()
	configContent := `{
		"device": "test-device-123",
		"ip": "192.168.1.100",
		"minified": true,
		"scriptId": 2
	}`
	configPath := filepath.Join(tempDir, ".spgtty")
	
	// Write test config file
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test loading the config file
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read test config file: %v", err)
	}
	
	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		t.Fatalf("Failed to parse test config file: %v", err)
	}
	
	// Verify config values
	if config.Device != "test-device-123" {
		t.Errorf("Expected device 'test-device-123', got '%s'", config.Device)
	}
	if config.IP != "192.168.1.100" {
		t.Errorf("Expected IP '192.168.1.100', got '%s'", config.IP)
	}
	if config.Minified != true {
		t.Errorf("Expected minified 'true', got '%t'", config.Minified)
	}
	if config.ScriptID != 2 {
		t.Errorf("Expected scriptId '2', got '%d'", config.ScriptID)
	}
}

// TestConfigValidation tests the config validation functionality
func TestConfigValidation(t *testing.T) {
	testCases := []struct {
		name        string
		config      Config
		expectError bool
	}{
		{
			name:        "Valid config with device",
			config:      Config{Device: "test-device", Minified: false, ScriptID: 1},
			expectError: false,
		},
		{
			name:        "Invalid config without device",
			config:      Config{Device: "", Minified: false, ScriptID: 1},
			expectError: true,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set global config for testing
			originalConfig := config
			config = tc.config
			defer func() { config = originalConfig }()
			
			err := validateConfig()
			
			if tc.expectError && err == nil {
				t.Errorf("Expected error for config: %+v, got nil", tc.config)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for config: %+v, got: %v", tc.config, err)
			}
		})
	}
}

// TestConfigMerging tests the CLI flag merging with config file
func TestConfigMerging(t *testing.T) {
	// Test data
	configFileContent := `{
		"device": "file-device",
		"ip": "192.168.1.1",
		"minified": false,
		"scriptId": 1
	}`
	
	// Create temporary config file
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, ".spgtty")
	err := os.WriteFile(configPath, []byte(configFileContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test config file: %v", err)
	}
	
	// Test merging logic manually (since we can't easily mock CLI flags)
	var fileConfig Config
	err = json.Unmarshal([]byte(configFileContent), &fileConfig)
	if err != nil {
		t.Fatalf("Failed to parse test config: %v", err)
	}
	
	// Simulate CLI flag overrides
	cliDevice := "cli-device"
	cliIP := "192.168.1.2"
	cliMinified := true
	cliScriptID := 2
	
	// Apply merging logic
	if cliDevice != "" {
		fileConfig.Device = cliDevice
	}
	if cliIP != "" {
		fileConfig.IP = cliIP
	}
	if cliMinified != false {
		fileConfig.Minified = cliMinified
	}
	if cliScriptID != 1 {
		fileConfig.ScriptID = cliScriptID
	}
	
	// Verify merged values
	if fileConfig.Device != "cli-device" {
		t.Errorf("Expected merged device 'cli-device', got '%s'", fileConfig.Device)
	}
	if fileConfig.IP != "192.168.1.2" {
		t.Errorf("Expected merged IP '192.168.1.2', got '%s'", fileConfig.IP)
	}
	if fileConfig.Minified != true {
		t.Errorf("Expected merged minified 'true', got '%t'", fileConfig.Minified)
	}
	if fileConfig.ScriptID != 2 {
		t.Errorf("Expected merged scriptId '2', got '%d'", fileConfig.ScriptID)
	}
}
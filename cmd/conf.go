package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Config-Struct mit Defaults + JSON-Tags
type Config struct {
	Device   string `json:"device"`       // Required
	IP       string `json:"ip,omitempty"` // Optional (mDNS-Fallback)
	Minified bool   `json:"minified"`     // Optional (default: false)
	ScriptID int    `json:"scriptId"`     // Optional (default: 1)
}

var config = Config{
	Minified: false, // Defaults für Felder, die nicht in .spgtty stehen
	ScriptID: 1,
}

// ConfigFileLocations defines the search order for .spgtty config files
var configFileLocations = []string{
	".spgtty",                          // Current directory
	filepath.Join("$", "HOME", ".spgtty"), // User home directory
	"/etc/spgtty/.spgtty",              // System-wide config
}

// CLI flag variables (separate from config to allow proper merging)
var (
	cliDevice   string
	cliIP       string
	cliMinified bool
	cliScriptID int
)

func init() {
	// 1. CLI-Flags definieren
	buildCmd.Flags().StringVar(&cliDevice, "device", "", "Shelly device ID (required)")
	buildCmd.Flags().StringVar(&cliIP, "ip", "", "Override device IP (optional)")
	buildCmd.Flags().BoolVar(&cliMinified, "minified", false, "Minify output (optional)")
	buildCmd.Flags().IntVar(&cliScriptID, "script-id", 1, "Script ID (optional)")

	// 2. .spgtty einlesen (falls vorhanden)
	configLoadErr := loadConfig()
	
	// 3. CLI-Flags mit Config mergen (CLI hat höchste Priorität)
	mergeCLIWithConfig()
	
	// 4. Device-Pflicht prüfen (nur wenn keine Config gefunden wurde UND kein Device gesetzt)
	if (configLoadErr != nil || config.Device == "") && cliDevice == "" {
		_ = buildCmd.MarkFlagRequired("device") // Device ist Pflicht!
	}
}

// mergeCLIWithConfig merges CLI flags with config file values
// Priority: CLI flags > config file > defaults
func mergeCLIWithConfig() {
	// CLI flags override config file values
	if cliDevice != "" {
		config.Device = cliDevice
	}
	if cliIP != "" {
		config.IP = cliIP
	}
	if cliMinified != false { // Only override if explicitly set
		config.Minified = cliMinified
	}
	if cliScriptID != 1 { // Only override if explicitly set (not default)
		config.ScriptID = cliScriptID
	}
}

func loadConfig() error {
	var fileContent []byte
	var err error
	
	// Try multiple locations in order
	for _, location := range configFileLocations {
		// Expand home directory placeholder
		if location == filepath.Join("$", "HOME", ".spgtty") {
			homeDir, homeErr := os.UserHomeDir()
			if homeErr != nil {
				continue // Skip if home dir can't be determined
			}
			location = filepath.Join(homeDir, ".spgtty")
		}
		
		fileContent, err = os.ReadFile(location)
		if err == nil {
			fmt.Printf("Using config file: %s\n", location)
			break // Found a config file, stop searching
		}
		// Continue to next location if file not found
	}
	
	if err != nil {
		// No config file found anywhere
		return errors.New("no .spgtty config file found in search locations")
	}
	
	// Parse the config file
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}
	
	return nil
}

// validateConfig checks that required fields are present
func validateConfig() error {
	if config.Device == "" {
		return errors.New("device is required. Please specify --device flag or set device in .spgtty config file")
	}
	return nil
}

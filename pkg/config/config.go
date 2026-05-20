// Package config provides hierarchical configuration management for spgtty.
// It uses Viper to load configuration from multiple sources in this order
// (later sources override earlier ones):
//
//  1. Defaults (hardcoded)
//  2. Global config: ~/.config/spgtty/config.yaml
//  3. Local config: .spgtty.yaml in current directory
//  4. Environment variables: SPGTTY_*
//  5. CLI flags (bound via BindPFlag)
package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Init initializes the Viper configuration.
// Call this in cobra.OnInitialize() in cmd/root.go.
func Init() {
	// Set config file name (without extension)
	viper.SetConfigName(".spgtty")
	viper.SetConfigType("yaml")

	// 1. Global config directory: ~/.config/spgtty/
	if home, err := os.UserHomeDir(); err == nil {
		globalConfigPath := filepath.Join(home, ".config", "spgtty")
		viper.AddConfigPath(globalConfigPath)
	}

	// 2. Local config: current working directory
	viper.AddConfigPath(".")

	// 3. Environment variables with SPGTTY_ prefix
	// Example: SPGTTY_SHELLY_HOST becomes shelly.host
	viper.SetEnvPrefix("SPGTTY")
	viper.AutomaticEnv()

	// Replace dots with underscores for env var names
	// shelly.host -> SPGTTY_SHELLY_HOST
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set defaults
	setDefaults()

	// Read config file (ignore error if not found)
	_ = viper.ReadInConfig()
}

// setDefaults sets default values for all configuration options.
func setDefaults() {
	// Build defaults
	viper.SetDefault("build.entry", "main.js")
	viper.SetDefault("build.output", "dist/main.js")
	viper.SetDefault("build.minify", true)

	// Shelly defaults
	viper.SetDefault("shelly.host", "")
	viper.SetDefault("shelly.script_id", 1)
}

// GetShellyHost returns the configured Shelly device host.
func GetShellyHost() string {
	return viper.GetString("shelly.host")
}

// GetShellyScriptID returns the configured script slot ID.
func GetShellyScriptID() int {
	return viper.GetInt("shelly.script_id")
}

// GetBuildEntry returns the entry file path.
func GetBuildEntry() string {
	return viper.GetString("build.entry")
}

// GetBuildOutput returns the output file path.
func GetBuildOutput() string {
	return viper.GetString("build.output")
}

// GetBuildMinify returns whether minification is enabled.
func GetBuildMinify() bool {
	return viper.GetBool("build.minify")
}

// ConfigFileUsed returns the config file path if one was loaded.
func ConfigFileUsed() string {
	return viper.ConfigFileUsed()
}

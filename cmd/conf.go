package cmd

import (
	"encoding/json"
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

func init() {
	// 1. CLI-Flags definieren (optional, überschreiben .spgtty)
	buildCmd.Flags().StringVar(&config.Device, "device", "", "Shelly device ID (required)")
	buildCmd.Flags().StringVar(&config.IP, "ip", "", "Override device IP (optional)")
	buildCmd.Flags().BoolVar(&config.Minified, "minified", false, "Minify output (optional)")
	buildCmd.Flags().IntVar(&config.ScriptID, "script-id", 1, "Script ID (optional)")

	// 2. .spgtty einlesen (falls vorhanden)
	if err := loadConfig(); err != nil {
		// Ignorieren, wenn Datei fehlt – Flags/Defaults gelten
		_ = buildCmd.MarkFlagRequired("device") // Device ist Pflicht!
	}
}

func loadConfig() error {
	file, err := os.ReadFile(filepath.Join(".", ".spgtty"))
	if err != nil {
		return err // Datei nicht gefunden? Egal, Flags/Defaults gelten.
	}
	return json.Unmarshal(file, &config) // Überschreibt config mit Werten aus .spgtty
}

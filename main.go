package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/evanw/esbuild/pkg/api"
)

// BuildShellyScript transpiliert JavaScript für Shelly-Geräte (Gen 2+)
// und schreibt das Ergebnis in `dist/main.js`.
func BuildShellyScript(entryPoint, outputPath string) error {
	// 1. Esbuild-Konfiguration (mit Dateiausgabe)
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{entryPoint},
		Bundle:            true,
		Target:            api.ES2015, // Niedrigste unterstützte Version
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		Outfile:           outputPath, // Ziel: dist/main.js
		Write:             true,       // Schreibe direkt in Datei
	})

	// 2. Fehlerbehandlung (ohne nested ifs)
	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			log.Printf("❌ Esbuild-Fehler: %v\n", err)
		}
		return fmt.Errorf("esbuild fehlgeschlagen")
	}

	// 3. Post-Processing: Hängende Kommas entfernen (direkt in der Datei)
	jsCode, err := os.ReadFile(outputPath)
	if err != nil {
		return fmt.Errorf("Fehler beim Lesen der Datei: %v", err)
	}

	re := regexp.MustCompile(`,[\s]*([}\]])`)
	cleanedCode := re.ReplaceAll(jsCode, []byte("$1"))

	// 4. Überschreibe die Datei mit dem bereinigten Code
	err = os.WriteFile(outputPath, cleanedCode, 0644)
	if err != nil {
		return fmt.Errorf("Fehler beim Schreiben der Datei: %v", err)
	}

	log.Printf("✅ Erfolg! Code nach %s geschrieben (Größe: %d Bytes)\n", outputPath, len(cleanedCode))
	return nil
}

func main() {
	// 1. Code transpilieren und in dist/main.js schreiben
	err := BuildShellyScript("main.js", "dist/main.js")
	if err != nil {
		log.Fatalf("❌ Fehler: %v", err)
	}

	// 2. Optional: Code direkt an Shelly deployen (z. B. per HTTP)
	// Beispiel: err = DeployToShelly("dist/main.js", "http://shelly-gen2-local/rpc")
}

package cmd

import (
	"log"
	"os"

	"github.com/GrosseBen/spgtty/pkg/builder"
	"github.com/spf13/cobra"
)

func build(cmd *cobra.Command, args []string) {
	noMinify := true
	entryPath := "main.js"
	outPath := "dist/main.js"
	code, err := builder.BuildShellyScript(entryPath, !noMinify)

	// 2. Optional: In Datei schreiben
	err = os.MkdirAll("dist", 0755) // Stelle sicher, dass das dist-Verzeichnis existiert
	if err != nil {
		log.Fatalf("❌ Fehler beim Erstellen des dist-Verzeichnisses: %v ", err)
	}
	err = os.WriteFile(outPath, code, 0644)
	if err != nil {
		log.Fatalf("❌ Fehler beim Schreiben der Datei: %v", err)
	}

	log.Printf("✅ Code nach %s geschrieben (%d Bytes)\n", outPath, len(code))
}

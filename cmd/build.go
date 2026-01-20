package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/GrosseBen/spgtty/pkg/builder"
	"github.com/spf13/cobra"
)

func build(cmd *cobra.Command, args []string) {
	var scriptName string
	if len(args) == 0 {
		args = append(args, "main.js")
	}
	var err error
	scriptName = args[len(args)-1]
	outPath := "./dist/"
	code, err := builder.BuildShellyScript(args[0], !notMinimizeFlagValue)
	if err != nil {
		log.Fatalf("Bulild failed %v ", err)
	}

	// 2. Optional: In Datei schreiben
	err = os.MkdirAll(outPath, 0755) // Stelle sicher, dass das dist-Verzeichnis existiert
	if err != nil {
		log.Fatalf("❌ Fehler beim Erstellen des dist-Verzeichnisses: %v ", err)
	}
	err = os.WriteFile(outPath+scriptName, code, 0644)
	if err != nil {
		log.Fatalf("❌ Fehler beim Schreiben der Datei: %v", err)
	}

	log.Printf("✅ Code nach %s geschrieben (%d Bytes)\n", outPath+scriptName, len(code))
	fmt.Printf("Building script from: %s\n", scriptName)
	fmt.Printf("Minimizing: %t\n", !notMinimizeFlagValue) // Zeigt an, ob minimiert wird
}

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/GrosseBen/spgtty/pkg/builder"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Translates your js code in a Spagetty monster for Shally devices",
	Long: `To enabele development fo more compley Js aplications for Shally devices

spgtty is a CLI tool that empowers developers to bundle complex, multifile applications, minmised and bundeld to one Fiel.
so that you could upload it to Shally Gen2+ Device`,
	Args: cobra.MaximumNArgs(1),
	Run:  build,
}

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

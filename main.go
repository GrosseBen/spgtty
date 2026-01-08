package main

import (
	"flag"
	"log"
	"os"

	"github.com/GrosseBen/spgtty/pkg/builder"
	"github.com/GrosseBen/spgtty/pkg/deployer"
	"github.com/GrosseBen/spgtty/pkg/utils"
)

func main() {
	// ------------------- Flags -------------------
	noMinify := flag.Bool("no-minify", false,
		"Deaktiviert Minifizierung (für Debugging)")
	outputPath := flag.String("out", "dist/main.js",
		"Ausgabepfad (optional, standardmäßig dist/main.js)")
	deployURL := flag.String("deploy", "",
		"Shelly-URL für direktes Deployment (optional)")
	version := flag.Bool("v", false,
		"shows the version)")
	// Parse nur die Optionen (alles, das mit '-' beginnt)
	flag.Parse()

	if *version {
		utils.PrintVersion()
		os.Exit(0)
	}
	// ------------------- Entry‑Datei -------------------
	// Standardwert:
	entryPath := "main.js"

	// Wenn ein Positions‑Argument vorhanden ist, dieses übernehmen
	if args := flag.Args(); len(args) > 0 && args[0] != "" {
		entryPath = args[0] // erstes Argument überschreibt den Default
	}
	// 1. Code transpilieren
	code, err := builder.BuildShellyScript(entryPath, !*noMinify)
	if err != nil {
		log.Fatalf("❌ Build fehlgeschlagen: %v", err)
	}

	// 2. Optional: In Datei schreiben
	if *outputPath != "" {
		err = os.MkdirAll("dist", 0755) // Stelle sicher, dass das dist-Verzeichnis existiert
		if err != nil {
			log.Fatalf("❌ Fehler beim Erstellen des dist-Verzeichnisses: %v ", err)
		}
		err = os.WriteFile(*outputPath, code, 0644)
		if err != nil {
			log.Fatalf("❌ Fehler beim Schreiben der Datei: %v", err)
		}
		log.Printf("✅ Code nach %s geschrieben (%d Bytes)\n", *outputPath, len(code))
	}
	// 3. Optional: Direkt an Shelly deployen
	if *deployURL != "" {
		err = deployer.DeployToShelly(code, *deployURL)
		if err != nil {
			log.Fatalf("❌ Deployment fehlgeschlagen: %v", err)
		}
		log.Printf("🚀 Code erfolgreich an %s gesendet!\n", *deployURL)
	}

}

package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/GrosseBen/spgtty/pkg/builder"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [file]",
	Short: "Translates your JS code into a Spagetty monster for Shelly devices",
	Long: `To enable development of more complex JS applications for Shelly devices,

spgtty is a CLI tool that empowers developers to bundle complex, multifile applications,
minified and bundled to one file, so that you can upload it to Shelly Gen2+ devices.`,
	Args: cobra.MaximumNArgs(1),
	Run:  build,
}

func build(cmd *cobra.Command, args []string) {
	// Determine input file
	entryFile := "main.js"
	if len(args) > 0 {
		entryFile = args[0]
	}

	// Output path: always dist/<filename>
	outDir := "./dist/"
	outFile := filepath.Base(entryFile) // Extract just the filename
	outPath := filepath.Join(outDir, outFile)

	// Build the script
	minify := !notMinimizeFlagValue
	code, err := builder.BuildShellyScript(entryFile, minify)
	if err != nil {
		log.Fatalf("Build failed: %v", err)
	}

	// Ensure dist directory exists
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create dist directory: %v", err)
	}

	// Write output file
	if err := os.WriteFile(outPath, code, 0644); err != nil {
		log.Fatalf("❌ Failed to write output file: %v", err)
	}

	log.Printf("✅ Code written to %s (%d bytes)\n", outPath, len(code))
	fmt.Printf("Entry: %s\n", entryFile)
	fmt.Printf("Minified: %t\n", minify)
}

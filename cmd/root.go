/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	versionFlagValue     bool // Speichert den Wert für --version / -v
	notMinimizeFlagValue bool // Speichert den Wert für --notMinimize / -m
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spgtty",
	Short: "Translates your js code in a Spagetty monster for Shally devices",
	Long: `
  SPGTTY - easy Shally script spagetty code handleing
#######################################################

To enabele development fo more compley Js aplications for Shally devices,
spgtty is a CLI tool that empowers developers to bundle complex, multifile applications,
minmised and bundeld to one File, so that you could upload it to Shally Gen2+ Device

If no subcommand is given, 'spgtty' will attempt to build the current project.
You can also explicitly use 'spgtty build <source_file>'.

You can use flags like '-v' or '-m' directly with 'spgtty' (e.g., 'spgtty -v')
or with any subcommand (e.g., 'spgtty build -m <file>').`,
	PreRunE: version,
	Run:     build,
	Args:    cobra.MaximumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Ausführen des Kommandos: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)

	// Flags an die Paket-Variablen binden
	rootCmd.PersistentFlags().BoolVarP(&versionFlagValue, "version", "v", false, "shows the version and build parameters")
	rootCmd.PersistentFlags().BoolVarP(&notMinimizeFlagValue, "notMinimize", "m", false, "do not minimize e.g. for debugging")
}

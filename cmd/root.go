/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	versionFlagValue     bool // Speichert den Wert für --version / -v
	notMinimizeFlagValue bool // Speichert den Wert für --notMinimize / -m
)

//go:embed "LongDescription.txt"
var long string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "spgtty",
	Short:   "Translates your js code in a Spagetty monster for Shally devices",
	Long:    long,
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

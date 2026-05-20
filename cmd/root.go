/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/GrosseBen/spgtty/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	versionFlagValue     bool // Stores the value for --version / -v
	notMinimizeFlagValue bool // Stores the value for --notMinimize / -m
)

//go:embed "LongDescription.txt"
var long string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spgtty",
	Short: "Translates your JS code into a Spagetty monster for Shelly devices",
	Long: `
  SPGTTY - easy Shelly script Spagetty code handling
#######################################################

To enable development of more complex JS applications for Shelly devices,
spgtty is a CLI tool that empowers developers to bundle complex, multifile applications,
minified and bundled to one file, so that you can upload it to Shelly Gen2+ devices.

If no subcommand is given, 'spgtty' will attempt to build the current project.
You can also explicitly use 'spgtty build <source_file>'.

Configuration can be set via:
  - Config file: .spgtty.yaml (local) or ~/.config/spgtty/config.yaml (global)
  - Environment: SPGTTY_SHELLY_HOST, SPGTTY_BUILD_MINIFY, etc.
  - CLI flags: --host, -m, etc. (highest priority)

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
		fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	// Initialize Viper configuration before commands run
	cobra.OnInitialize(config.Init)

	// Add subcommands
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)

	// Global flags (available on all commands)
	rootCmd.PersistentFlags().BoolVarP(&versionFlagValue, "version", "v", false, "Show version and build parameters")
	rootCmd.PersistentFlags().BoolVarP(&notMinimizeFlagValue, "notMinimize", "m", false, "Do not minimize output (for debugging)")

	// Shelly connection flags (for upload command)
	rootCmd.PersistentFlags().StringP("host", "H", "", "Shelly device IP or hostname")
	rootCmd.PersistentFlags().Int("script-id", 1, "Script slot on Shelly device (1-10)")

	// Bind flags to Viper config keys
	viper.BindPFlag("shelly.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("shelly.script_id", rootCmd.PersistentFlags().Lookup("script-id"))
	viper.BindPFlag("build.minify", rootCmd.PersistentFlags().Lookup("notMinimize"))
}

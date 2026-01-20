/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spgtty",
	Short: "Translates your js code in a Spagetty monster for Shally devices",
	Long: `To enabele development fo more compley Js aplications for Shally devices,
spgtty is a CLI tool that empowers developers to bundle complex, multifile applications, minmised and bundeld to one Fiel.
so that you could upload it to Shally Gen2+ Device`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Korrektur: PersistentFlags() statt Flags() verwenden, um persistente Flags abzufragen
		versionFlag, err := cmd.PersistentFlags().GetBool("version")
		if err != nil {
			// Falls ein Fehler beim Abrufen des Flags auftritt, gib ihn zurück.
			// Dies kann passieren, wenn das Flag nicht existiert oder der Typ falsch ist.
			return fmt.Errorf("Fehler beim Abrufen des 'version'-Flags: %w", err)
		}
		if versionFlag {
			utils.PrintVersion()
			os.Exit(0) // Beendet das Programm, nachdem die Version angezeigt wurde
		}
		// Wichtig: Hier nil zurückgeben, damit die Ausführung fortgesetzt wird,
		// wenn das Flag nicht gesetzt ist und keine Version angezeigt werden soll.
		return nil
	},
	Run: build,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Ausführen des Kommandos: %s\n", err)
		os.Exit(1)
	}
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Translates your js code in a Spagetty monster for Shally devices",
	Long: `To enabele development fo more compley Js aplications for Shally devices

spgtty is a CLI tool that empowers developers to bundle complex, multifile applications, minmised and bundeld to one Fiel.
so that you could upload it to Shally Gen2+ Device`,
	Args: cobra.MaximumNArgs(1),
	Run:  build,
}

// Definiere dein Cobra-Befehlsobjekt für den Upload
var uploadCmd = &cobra.Command{
	Use:   "upload", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "Uploads files to a specified destination",
	Long:  `This command handles the uploading of various file types to a remote server or service.`,
	Run:   upload, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

var initCmd = &cobra.Command{
	Use:   "init", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "create a new JS poject for shally Gen2+",
	Long:  `create a new JS poject for shally Gen2+`,
	Run:   initProj, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "shows the version and build parameters")
	rootCmd.Flags().BoolP("notMinimize", "m", false, "do not minimize e.g. for debugging")
}

/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spgtty",
	Short: "Translates your js code in a Spagetty monster for Shally devices",
	Long: `To enabele development fo more compley Js aplications for Shally devices

spgtty is a CLI tool that empowers developers to bundle complex, multifile applications, minmised and bundeld to one Fiel.
so that you could upload it to Shally Gen2+ Device`,
	Run: build,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
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

var VersionCmd = &cobra.Command{
	Use:     "version", // Der Name des Unterbefehls
	Aliases: []string{"-v"},
	Short:   "shows actual versionx",
	Long:    `shows the verson and build parameters`,
	Run:     version, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

func init() {
	//entryPath := "main.js"
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//

	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(VersionCmd)
	rootCmd.PersistentFlags().BoolP("version", "v", false, "shows the version and build parameters")

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spgtty.yaml)")

	//Cobra also supports local flags, which will only run
	//when this action is called directly.
	rootCmd.Flags().BoolP("-notMinimize", "m", false, "do not minimize e.g. for debugging")
}

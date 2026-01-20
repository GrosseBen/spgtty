package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "create a new JS poject for shally Gen2+",
	Long:  `create a new JS poject for shally Gen2+`,
	Run:   initProj, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

func initProj(cmd *cobra.Command, args []string) {
	panic("init not jet inmplemented")
}

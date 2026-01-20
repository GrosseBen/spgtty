package cmd

import (
	"github.com/spf13/cobra"
)

// Definiere dein Cobra-Befehlsobjekt für den Upload
var uploadCmd = &cobra.Command{
	Use:   "upload", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "Uploads files to a specified destination",
	Long:  `This command handles the uploading of various file types to a remote server or service.`,
	//Args:  cobra.MinARGES(1),
	Run: upload, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

func upload(cmd *cobra.Command, args []string) {
	panic("upload not jet inmplemented")
}

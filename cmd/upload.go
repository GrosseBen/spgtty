package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// Definiere dein Cobra-Befehlsobjekt für den Upload
var uploadCmd = &cobra.Command{
	Use:   "upload [http://<hostname or IP>] [<file.js>]", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "Uploads files to a specified destination",
	Long:  `This command handles the uploading of various file types to a remote server or service.`,
	Args:  cobra.OnlyValidArgs,
	Run:   upload, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
	//toDo: add script ID
}

func upload(cmd *cobra.Command, args []string) {
	j, _ := json.Marshal(args)
	log.Println("upload ", string(j))
	// Use the config for uploading
	fmt.Printf("Uploading to device %s with IP %s, minified: %t, script ID: %d\n", config.Device, config.IP, config.Minified, config.ScriptID)
}

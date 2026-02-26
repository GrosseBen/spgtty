package cmd

import (
	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func version(cmd *cobra.Command, args []string) error {
	versionFlag, err := cmd.PersistentFlags().GetBool("version")
	if err != nil {
		log.Fatalf("Fehler beim Abrufen des 'version'-Flags: %v", err)
	}
	if versionFlag {
		utils.PrintVersion()
		os.Exit(0) // Beendet das Programm, nachdem die Version angezeigt wurde
	}

	return err
}

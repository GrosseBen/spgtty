package cmd

import (
	"fmt"
	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func version(cmd *cobra.Command, args []string) error {
	versionFlag, err := cmd.PersistentFlags().GetBool("version")
	if err != nil {
		return fmt.Errorf("Fehler beim Abrufen des 'version'-Flags: %w", err)
	}
	if versionFlag {
		utils.PrintVersion()
		os.Exit(0) // Beendet das Programm, nachdem die Version angezeigt wurde
	}

	notMinimizeFlagValue, err = cmd.PersistentFlags().GetBool("notMinimize")
	if err != nil {
		log.Fatalf("notMinimize failed %v ", err)
	}

	return nil
}

package cmd

import (
	"log"
	"os"

	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
)

func version(cmd *cobra.Command, args []string) error {
	versionFlag, err := cmd.PersistentFlags().GetBool("version")
	if err != nil {
		log.Fatalf("Error getting 'version' flag: %v", err)
	}
	if versionFlag {
		utils.PrintVersion()
		os.Exit(0)
	}

	return err
}

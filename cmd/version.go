package cmd

import (
	"os"

	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
)

func version(cmd *cobra.Command, args []string) {
	utils.PrintVersion()
	os.Exit(0)
}

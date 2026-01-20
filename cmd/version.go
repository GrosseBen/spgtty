package cmd

import (
	"fmt"
	"github.com/GrosseBen/spgtty/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

func version(cmd *cobra.Command, args []string) {
	err := utils.PrintVersion()
	os.Exit(0)
}

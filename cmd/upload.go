package cmd

import (
	"log"

	"github.com/GrosseBen/spgtty/pkg/config"
	"github.com/GrosseBen/spgtty/pkg/deployer"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload script to a Shelly device",
	Long: `Uploads a JavaScript file to a Shelly Gen2+ device via the RPC API.

The file will be uploaded in chunks (1024 bytes each) to handle
Shelly's request size limitations.

Configuration can be set via:
  - CLI flags: --host, --script-id
  - Environment: SPGTTY_SHELLY_HOST, SPGTTY_SHELLY_SCRIPT_ID
  - Config file: .spgtty.yaml

Examples:
  spgtty upload                           # Upload dist/main.js to configured host
  spgtty upload dist/app.js               # Upload specific file
  spgtty upload --host 192.168.1.100      # Specify host via flag
  spgtty upload -H shelly.local --script-id 2  # Upload to script slot 2`,
	Args: cobra.MaximumNArgs(1),
	Run:  upload,
}

func upload(cmd *cobra.Command, args []string) {
	// Get host from config/flags
	host := config.GetShellyHost()
	if host == "" {
		log.Fatal("No Shelly host configured.\n" +
			"Set it via:\n" +
			"  - Flag: spgtty upload --host 192.168.1.100\n" +
			"  - Environment: export SPGTTY_SHELLY_HOST=192.168.1.100\n" +
			"  - Config file: shelly.host in .spgtty.yaml")
	}

	// Get script ID from config/flags
	scriptID := config.GetShellyScriptID()

	// Determine file to upload
	filePath := config.GetBuildOutput() // Default: dist/main.js
	if len(args) > 0 {
		filePath = args[0]
	}

	// Upload the file
	if err := deployer.Upload(host, scriptID, filePath); err != nil {
		log.Fatalf("Upload failed: %v", err)
	}
}

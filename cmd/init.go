package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI flags for init command
var (
	initDevice   string
	initIP       string
	initMinified bool
	initScriptID int
)

var initCmd = &cobra.Command{
	Use:   "init", // Der Name des Unterbefehls, den der Benutzer eingibt (z.B. "yourcli upload")
	Short: "create a new JS project for Shelly Gen2+ devices",
	Long:  "create a new JS project for Shelly Gen2+ devices with a .spgtty config file",
	Run:   initProj, // Hier wird deine 'upload'-Funktion als Run-Feld zugewiesen
}

func init() {
	// Define CLI flags for init command
	initCmd.Flags().StringVar(&initDevice, "device", "", "Shelly device ID (optional)")
	initCmd.Flags().StringVar(&initIP, "ip", "", "Device IP address (optional)")
	initCmd.Flags().BoolVar(&initMinified, "minified", false, "Minify output (optional)")
	initCmd.Flags().IntVar(&initScriptID, "script-id", 1, "Script ID (optional)")
}

func initProj(cmd *cobra.Command, args []string) {
	// Create default config
	config := Config{
		Device:   initDevice,
		IP:       initIP,
		Minified: initMinified,
		ScriptID: initScriptID,
	}
	
	// If device not specified via flag, prompt for it
	if config.Device == "" {
		fmt.Print("Enter Shelly device ID (e.g., shellyplus1pm-123456): ")
		var deviceInput string
		fmt.Scanln(&deviceInput)
		if deviceInput != "" {
			config.Device = deviceInput
		}
	}
	
	// Create .spgtty config file
	configContent, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error creating config file: %v\n", err)
		return
	}
	
	// Write config file
	err = os.WriteFile(".spgtty", configContent, 0644)
	if err != nil {
		fmt.Printf("❌ Error writing config file: %v\n", err)
		return
	}
	
	// Create dist directory
	err = os.MkdirAll("./dist", 0755)
	if err != nil {
		fmt.Printf("⚠️  Warning: Could not create dist directory: %v\n", err)
	}
	
	// Create example main.js file
	exampleJS := `// Shelly Gen2+ JavaScript example
// This is a basic example to get you started

print("Hello from Shelly device: " + Shelly.getDeviceInfo().id);

// Example: Toggle relay 0 every 5 seconds
// Timer.set(5000, true, function() {
//   let state = Shelly.getComponentStatus("switch:0").output;
//   Shelly.call("switch.set", {id: 0, on: !state});
// });
`
	
	err = os.WriteFile("main.js", []byte(exampleJS), 0644)
	if err != nil {
		fmt.Printf("⚠️  Warning: Could not create main.js file: %v\n", err)
	}
	
	fmt.Printf("✅ Project initialized successfully!\n")
	fmt.Printf("📄 Created .spgtty config file\n")
	fmt.Printf("📄 Created main.js example file\n")
	fmt.Printf("📁 Created dist/ directory\n")
	fmt.Printf("\n🚀 Ready to build! Run: spgtty build\n")
	
	// Show the created config
	fmt.Printf("\n📋 Config file contents:\n")
	fmt.Printf("```json\n%s\n```\n", string(configContent))
}

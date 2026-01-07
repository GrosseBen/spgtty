package builder

import (
	"fmt"
	"regexp"

	"github.com/evanw/esbuild/pkg/api"
)

// BuildShellyScript transpiliert JavaScript für Shelly-Geräte (Gen 2+)
func BuildShellyScript(entryPath string, minify bool) ([]byte, error) { // minify-Parameter hinzugefügt
	// Esbuild-Konfiguration (mit dynamischer Minifizierung)
	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{entryPath},
		Bundle:            true,
		Target:            api.ES2015,
		Format:            api.FormatCommonJS, // WICHTIG für Shelly!
		MinifyIdentifiers: minify,             // Dynamisch!
		MinifySyntax:      minify,             // Dynamisch!
		MinifyWhitespace:  minify,             // Dynamisch!
		Write:             false,
	})

	// 2. Fehlerbehandlung
	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("esbuild Fehler: %v", result.Errors)
	}
	if len(result.OutputFiles) == 0 {
		return nil, fmt.Errorf("keine Ausgabe von esbuild")
	}

	// 3. Post-Processing: Hängende Kommas entfernen
	jsCode := result.OutputFiles[0].Contents
	re := regexp.MustCompile(`,[\s]*([}\]])`)
	cleanedCode := re.ReplaceAll(jsCode, []byte("$1"))

	return cleanedCode, nil
}

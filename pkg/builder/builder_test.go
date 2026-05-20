package builder

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// getFixturePath returns the absolute path to a test fixture file.
func getFixturePath(t *testing.T, name string) string {
	// Get the directory of this test file
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	
	// Navigate to test/fixtures from pkg/builder
	fixturePath := filepath.Join(wd, "..", "..", "test", "fixtures", name)
	
	// Verify file exists
	if _, err := os.Stat(fixturePath); err != nil {
		t.Fatalf("Fixture file not found: %s", fixturePath)
	}
	
	return fixturePath
}

func TestBuildSimpleScript(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "simple.js"), true)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	if len(result) == 0 {
		t.Fatal("Empty output")
	}
	
	output := string(result)
	if !strings.Contains(output, "print") {
		t.Error("Output should contain print statement")
	}
	if !strings.Contains(output, "Shelly") {
		t.Error("Output should contain 'Shelly' string")
	}
}

func TestBuildWithMinify(t *testing.T) {
	fixture := getFixturePath(t, "functions.js")
	
	minified, err := BuildShellyScript(fixture, true)
	if err != nil {
		t.Fatalf("Minified build failed: %v", err)
	}
	
	unminified, err := BuildShellyScript(fixture, false)
	if err != nil {
		t.Fatalf("Unminified build failed: %v", err)
	}
	
	if len(minified) >= len(unminified) {
		t.Errorf("Minified (%d bytes) should be smaller than unminified (%d bytes)",
			len(minified), len(unminified))
	}
}

func TestTrailingCommaRemovalInObjects(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "objects.js"), false)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	output := string(result)
	
	// Check for trailing commas before closing braces
	if strings.Contains(output, ",}") {
		t.Error("Trailing commas before } should be removed")
	}
	if strings.Contains(output, ", }") {
		t.Error("Trailing commas before } (with space) should be removed")
	}
}

func TestTrailingCommaRemovalInArrays(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "arrays.js"), false)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	output := string(result)
	
	// Check for trailing commas before closing brackets
	if strings.Contains(output, ",]") {
		t.Error("Trailing commas before ] should be removed")
	}
	if strings.Contains(output, ", ]") {
		t.Error("Trailing commas before ] (with space) should be removed")
	}
}

func TestBuildWithImports(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "imports.js"), true)
	if err != nil {
		t.Fatalf("Build with imports failed: %v", err)
	}
	
	output := string(result)
	
	// Import statements should be resolved and bundled
	if strings.Contains(output, "import ") {
		t.Error("Imports should be bundled, not left as import statements")
	}
	if strings.Contains(output, "export ") {
		t.Error("Exports should be removed after bundling")
	}
	
	// The helper function content should be included
	if !strings.Contains(output, "helper") {
		t.Error("Helper function should be included in bundled output")
	}
}

func TestBuildNonExistentFile(t *testing.T) {
	_, err := BuildShellyScript("nonexistent_file_that_does_not_exist.js", true)
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestBuildEmptyPath(t *testing.T) {
	// BuildShellyScript with empty path should default to main.js
	// which likely doesn't exist, so we expect an error
	_, err := BuildShellyScript("", true)
	if err == nil {
		t.Error("Expected error for empty/default path that doesn't exist")
	}
}

func TestBuildVariables(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "variables.js"), true)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	output := string(result)
	
	// Check that variable declarations are present
	if !strings.Contains(output, "Shelly") {
		t.Error("Output should contain 'Shelly' string from variable")
	}
}

func TestBuildArrowFunctions(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "arrow.js"), true)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	if len(result) == 0 {
		t.Fatal("Empty output")
	}
	
	// Arrow functions should be transpiled (ES2015 target)
	// The exact output depends on esbuild's transpilation
}

func TestBuildShellyAPI(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "shelly_api.js"), true)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	output := string(result)
	
	// Shelly API calls should be preserved
	if !strings.Contains(output, "Shelly") {
		t.Error("Output should contain Shelly API calls")
	}
	if !strings.Contains(output, "Timer") {
		t.Error("Output should contain Timer API calls")
	}
}

// TestBuildAllFixtures tests that all fixture files can be built successfully
func TestBuildAllFixtures(t *testing.T) {
	fixtures := []string{
		"simple.js",
		"variables.js",
		"functions.js",
		"arrow.js",
		"objects.js",
		"arrays.js",
		"imports.js",
		"shelly_api.js",
	}
	
	for _, fixture := range fixtures {
		t.Run(fixture, func(t *testing.T) {
			result, err := BuildShellyScript(getFixturePath(t, fixture), true)
			if err != nil {
				t.Errorf("Failed to build %s: %v", fixture, err)
				return
			}
			if len(result) == 0 {
				t.Errorf("Empty output for %s", fixture)
			}
		})
	}
}

// TestBuildOutputIsValidJS tests that the output doesn't contain obvious JS errors
func TestBuildOutputIsValidJS(t *testing.T) {
	result, err := BuildShellyScript(getFixturePath(t, "functions.js"), false)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}
	
	output := string(result)
	
	// Basic sanity checks
	if strings.Count(output, "{") != strings.Count(output, "}") {
		t.Error("Mismatched curly braces")
	}
	if strings.Count(output, "(") != strings.Count(output, ")") {
		t.Error("Mismatched parentheses")
	}
	if strings.Count(output, "[") != strings.Count(output, "]") {
		t.Error("Mismatched square brackets")
	}
}

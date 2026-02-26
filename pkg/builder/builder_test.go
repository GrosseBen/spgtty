package builder

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestBuildShellyScript tests the basic JavaScript building functionality
func TestBuildShellyScript(t *testing.T) {
	testCases := []struct {
		name        string
		content    string
		minify     bool
		expectError bool
	}{
		{
			name:        "Simple JavaScript",
			content:    "print('Hello World');",
			minify:     false,
			expectError: false,
		},
		{
			name:        "JavaScript with Shelly API",
			content:    `print("Device: " + Shelly.getDeviceInfo().id);`,
			minify:     false,
			expectError: false,
		},
		{
			name:        "Empty content",
			content:    "",
			minify:     false,
			expectError: false, // Empty file is valid JavaScript
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create temporary file
			tempFile, err := createTempJSFile(tc.content)
			if err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			defer os.Remove(tempFile)

			code, err := BuildShellyScript(tempFile, tc.minify)

			if tc.expectError && err == nil {
				t.Errorf("Expected error for content: '%s', got nil", tc.content)
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for content: '%s', got: %v", tc.content, err)
			}
			if !tc.expectError && len(code) == 0 && tc.content != "" {
				t.Errorf("Expected non-empty output for content: '%s', got empty", tc.content)
			}
		})
	}
}

// TestMinification tests the minification functionality
func TestMinification(t *testing.T) {
	content := `
	// This is a comment
	print("Hello"); // Another comment
	print("World");
	`

	// Create temporary file
	tempFile, err := createTempJSFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile)

	// Test without minification
	nonMinified, err := BuildShellyScript(tempFile, false)
	if err != nil {
		t.Fatalf("Failed to build non-minified code: %v", err)
	}

	// Test with minification
	minified, err := BuildShellyScript(tempFile, true)
	if err != nil {
		t.Fatalf("Failed to build minified code: %v", err)
	}

	// Minified version should be shorter or equal
	if len(minified) > len(nonMinified) {
		t.Errorf("Minified code (%d bytes) should be shorter than non-minified (%d bytes)", len(minified), len(nonMinified))
	}

	// Both should contain the essential code
	if !contains(string(nonMinified), "print") || !contains(string(minified), "print") {
		t.Errorf("Both minified and non-minified code should contain 'print' statements")
	}
}

// Helper function to create temporary JavaScript file
func createTempJSFile(content string) (string, error) {
	tempFile, err := ioutil.TempFile("", "test-*.js")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	_, err = tempFile.Write([]byte(content))
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsSubstring(s, substr)))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
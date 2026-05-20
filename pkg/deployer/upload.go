// Package deployer provides functions to upload scripts to Shelly Gen2+ devices.
package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// ChunkSize is the maximum number of bytes per upload request.
// Shelly devices have a limit of approximately 1024 bytes per RPC request.
const ChunkSize = 1024

// Upload uploads a JavaScript file to a Shelly device.
// It reads the file, splits it into chunks, and sends them via the Script.PutCode RPC.
//
// Parameters:
//   - host: IP address or hostname of the Shelly device (e.g., "192.168.1.100")
//   - scriptID: Script slot on the device (1-10)
//   - filePath: Path to the JavaScript file to upload
//
// Returns an error if the upload fails.
func Upload(host string, scriptID int, filePath string) error {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	code := string(content)
	fmt.Printf("Uploading %d bytes to %s (script %d)...\n", len(code), host, scriptID)

	// Upload in chunks
	return uploadChunked(host, scriptID, code)
}

// UploadCode uploads JavaScript code directly to a Shelly device.
// Use this when you already have the code in memory (e.g., after building).
//
// Parameters:
//   - host: IP address or hostname of the Shelly device
//   - scriptID: Script slot on the device (1-10)
//   - code: JavaScript code to upload
//
// Returns an error if the upload fails.
func UploadCode(host string, scriptID int, code string) error {
	fmt.Printf("Uploading %d bytes to %s (script %d)...\n", len(code), host, scriptID)
	return uploadChunked(host, scriptID, code)
}

// uploadChunked uploads code in chunks of ChunkSize bytes.
func uploadChunked(host string, scriptID int, code string) error {
	id := strconv.Itoa(scriptID)
	totalChunks := (len(code) + ChunkSize - 1) / ChunkSize

	pos := 0
	chunkNum := 0
	for pos < len(code) {
		end := pos + ChunkSize
		if end > len(code) {
			end = len(code)
		}
		chunk := code[pos:end]
		chunkNum++

		// First chunk overwrites, subsequent chunks append
		appendMode := pos > 0

		fmt.Printf("  Chunk %d/%d (%d bytes)...\n", chunkNum, totalChunks, len(chunk))
		if err := putChunk(host, id, chunk, appendMode); err != nil {
			return fmt.Errorf("failed to upload chunk %d: %w", chunkNum, err)
		}

		pos = end
	}

	fmt.Printf("✅ Upload complete (%d chunks)\n", chunkNum)
	return nil
}

// putChunk sends a single chunk of code to the Shelly device.
func putChunk(host, id string, data string, appendMode bool) error {
	url := fmt.Sprintf("http://%s/rpc/Script.PutCode", host)

	payload := map[string]interface{}{
		"id":     id,
		"code":   data,
		"append": appendMode,
	}

	reqData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Shelly returned status %d: %s", resp.StatusCode, string(body))
	}

	// Check for RPC error in response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if errObj, ok := result["error"]; ok {
		return fmt.Errorf("Shelly RPC error: %v", errObj)
	}

	return nil
}

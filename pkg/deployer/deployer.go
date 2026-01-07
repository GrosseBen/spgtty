package deployer

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// DeployToShelly sendet den Code an ein Shelly-Gerät
func DeployToShelly(code []byte, shellyURL string) error { // Großbuchstaben!
	resp, err := http.Post(shellyURL, "application/json", bytes.NewReader(code))
	if err != nil {
		return fmt.Errorf("Fehler beim Senden an Shelly: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Shelly antwortete mit %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

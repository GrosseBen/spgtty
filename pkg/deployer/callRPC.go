package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// -------------------------------------------------------------------
// callRPC – führt einen einzelnen RPC‑Aufruf aus
// -------------------------------------------------------------------
func callRPC(url string, payload rpcRequest) error {
	// 1️⃣ Payload → JSON konvertieren
	body, err := json.Marshal(payload)
	if err != nil {
		// Fehler beim Marshallen → sofort zurückgeben
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	// 2️⃣ HTTP‑POST‑Request zusammenbauen
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	// Wichtig: dem Server sagen, dass wir JSON senden
	req.Header.Set("Content-Type", "application/json")

	// 3️⃣ Request absenden
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close() // Body schließen, sobald wir fertig sind

	// 4️⃣ Antwort auslesen (nur zu Debug‑Zwecken)
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nResponse: %s\n", resp.Status, string(respBody))

	// Keine weiteren Fehler → nil zurückgeben
	return nil
}

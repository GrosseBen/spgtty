package deployer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DeployToShelly(code []byte, deployURL string) error {
	var ErrNotImplemented = errors.New("not implemented")

	return ErrNotImplemented
}

// -------------------------------------------------
func deployToShelly(scriptName string, scriptNum int, code string, shellyURL string) error {
	// 1️⃣ RPC‑Payload zusammenbauen
	payload := rpcRequest{
		ID:     1,                // beliebige Request‑ID, hier fest 1
		Method: "Script.PutCode", // Methode, die das Shelly‑Device kennt
		Params: putCodeParams{
			ID:     scriptNum,
			Code:   code,
			Append: false, // beim ersten Aufruf: nicht anhängen
		},
	}

	// 2️⃣ In JSON serialisieren
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON‑Marshal fehlgeschlagen: %w", err)
	}

	// 3️⃣ POST an das Gerät senden
	resp, err := http.Post(shellyURL+"/rpc/", "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return fmt.Errorf("Fehler beim Senden an Shelly (%s): %w", scriptName, err)
	}
	defer resp.Body.Close()

	// 4️⃣ Antwort prüfen
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Shelly (%s) antwortete mit %d: %s", scriptName, resp.StatusCode, string(body))
	}

	// Optional: Ausgabe für das Log
	fmt.Printf("✅ %s (ID %d) erfolgreich deployed\n", scriptName, scriptNum)
	return nil
}

func DeployAppend(scriptName string, scriptNum int, code string, shellyURL string) error {
	// Fast identisch zu DeployToShelly, nur Append = true
	payload := rpcRequest{
		ID:     1,
		Method: "Script.PutCode",
		Params: putCodeParams{
			ID:     scriptNum,
			Code:   code,
			Append: true,
		},
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON‑Marshal fehlgeschlagen: %w", err)
	}
	resp, err := http.Post(shellyURL+"/rpc/", "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return fmt.Errorf("Fehler beim Anhängen an Shelly (%s): %w", scriptName, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Shelly (%s) antwortete mit %d: %s", scriptName, resp.StatusCode, string(b))
	}
	fmt.Printf("✅ %s (ID %d) – Anhang erfolgreich\n", scriptName, scriptNum)
	return nil
}

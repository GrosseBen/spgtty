package deployer

// Rückgabe‑Struktur für Existenz‑Abfrage
type existsResult struct {
	Exists bool `json:"exists"`
}

// Rückgabe‑Struktur für Lauf‑Status
type runningResult struct {
	Running bool `json:"running"`
}

// Struktur für das allgemeine RPC‑Objekt
type rpcRequest struct {
	ID     int         `json:"id"`     // eindeutige Anfrage‑ID
	Method string      `json:"method"` // Name der aufzurufenden RPC‑Methode
	Params interface{} `json:"params"` // beliebige Parameter‑Payload
}

type putCodeParams struct {
	ID     int    `json:"id"`     // Script‑Nummer, die das Gerät erwartet
	Code   string `json:"code"`   // Quellcode, den du ausführen willst
	Append bool   `json:"append"` // true → an vorhandenen Code anhängen
}

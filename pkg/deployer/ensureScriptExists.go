package deployer

import "fmt"

func ensureScriptExists(url string, scriptID int) error {
	// 1️⃣ Existenz prüfen
	existsReq := rpcRequest{
		ID:     1,
		Method: "Script.Exists",
		Params: map[string]int{"id": scriptID},
	}
	var existsRes struct {
		Result existsResult `json:"result"`
	}
	if err := callRPCWithResult(url, existsReq, &existsRes); err != nil {
		return fmt.Errorf("existence check failed: %w", err)
	}
	if existsRes.Result.Exists {
		fmt.Println("Skript existiert bereits.")
	} else {
		// 2️⃣ Skript anlegen
		createReq := rpcRequest{
			ID:     1,
			Method: "Script.Create",
			Params: map[string]int{"id": scriptID},
		}
		if err := callRPC(url, createReq); err != nil {
			return fmt.Errorf("script creation failed: %w", err)
		}
		fmt.Println("Skript wurde neu angelegt.")
	}
	return nil
}

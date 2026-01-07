package deployer

import "fmt"

func abortIfRunning(url string, scriptID int) error {
	runningReq := rpcRequest{
		ID:     1,
		Method: "Script.IsRunning",
		Params: map[string]int{"id": scriptID},
	}
	var runningRes struct {
		Result runningResult `json:"result"`
	}
	if err := callRPCWithResult(url, runningReq, &runningRes); err != nil {
		return fmt.Errorf("running check failed: %w", err)
	}
	if runningRes.Result.Running {
		return fmt.Errorf("Skript %d läuft bereits – Vorgang wird abgebrochen", scriptID)
	}
	return nil
}

package deployer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func callRPCWithResult(url string, payload rpcRequest, result interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nResponse: %s\n", resp.Status, string(respBody))

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("response unmarshal failed: %w", err)
		}
	}
	return nil
}
